package messaging

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/talbs1986/simplego/app/pkg/logger"
	simplego "github.com/talbs1986/simplego/messaging/pkg/messaging"
)

const (
	testSubject      = "testEvent"
	testDestination  = "simplego-tests"
	testNATSUser     = "simplego-user"
	testNATSPassword = "simplego-password"
)

var testMessageType = simplego.MessageType("integration-test")

type NATSSuite struct {
	suite.Suite

	natsURL   string
	container testcontainers.Container
}

func TestNATSSuite(t *testing.T) {
	suite.Run(t, new(NATSSuite))
}

// SetupSuite runs once before all tests in the suite.
func (s *NATSSuite) SetupSuite() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Disable Ryuk sidecar to avoid pulling extra images
	err := os.Setenv("TESTCONTAINERS_RYUK_DISABLED", "true")
	s.Require().NoError(err, "failed to set env for testcontainers")

	req := testcontainers.ContainerRequest{
		Image:        "nats:2.10.14",
		ExposedPorts: []string{"4222/tcp"},
		WaitingFor:   wait.ForLog("Server is ready").WithStartupTimeout(30 * time.Second),
		Cmd:          []string{"-js", "--user", testNATSUser, "--pass", testNATSPassword},
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		// If Docker is not available, skip the whole suite
		if strings.Contains(err.Error(), "Cannot connect to the Docker daemon") ||
			strings.Contains(err.Error(), "permission denied") {
			s.T().Skipf("docker not available for NATS integration tests: %v", err)
		}
		s.Require().NoError(err, "failed to start nats container")
	}

	host, err := container.Host(ctx)
	s.Require().NoError(err, "failed to resolve nats container host")

	mappedPort, err := container.MappedPort(ctx, "4222")
	s.Require().NoError(err, "failed to resolve nats container port")

	s.natsURL = "nats://" + host + ":" + mappedPort.Port()
	s.container = container
}

// TearDownSuite runs once after all tests in the suite.
func (s *NATSSuite) TearDownSuite() {
	if s.container == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_ = s.container.Terminate(ctx)
}

func (s *NATSSuite) TestNATSMessagingAsyncConsume() {
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()

	streamCfg := testStreamConfig()
	log := logger.NewFMTLogger(nil)

	pubCfg, consCfg := buildConfigs(s.natsURL, streamCfg)

	pub, err := NewNATSPublisher(log, pubCfg, WithPublisherUpsertStream())
	s.Require().NoError(err)
	defer pub.Close(ctx)

	cons, err := NewNATSConsumer(log, consCfg, WitConsumerUpsertStream())
	s.Require().NoError(err)
	defer cons.Close(ctx)

	// Cleanup any stream data created by this test
	defer func() {
		natsCons := cons.(*natsConsumerImpl)
		err := natsCons.stream.DeleteConsumer(natsCons.streamName, natsCons.consumerGroup)
		s.Require().NoError(err)

		natsPub := pub.(*natsPublisherImpl)
		err = natsPub.stream.PurgeStream(streamCfg.Name)
		s.Require().NoError(err)
	}()

	msgCh := make(chan *simplego.MessageWrapper, 1)
	err = cons.Consume(testSubject, func(msg *simplego.MessageWrapper) error {
		msgCh <- msg
		return nil
	})
	s.Require().NoError(err)

	// give the subscription a moment to bind before publishing
	time.Sleep(200 * time.Millisecond)

	payload := map[string]string{"hello": "world"}
	sentMsg, err := pub.Publish(testSubject, testMessageType, payload)
	s.Require().NoError(err)

	select {
	case received := <-msgCh:
		s.Require().Equal(sentMsg.ID, received.ID)
		s.Require().Equal(testMessageType, received.Type)
		s.Require().Equal(sentMsg.Sender, received.Sender)

		payloadMap, ok := received.Payload.(map[string]interface{})
		s.Require().True(ok, "payload should be a map after json unmarshal")
		s.Require().Equal("world", payloadMap["hello"])
	case <-time.After(5 * time.Second):
		s.T().Fatalf("did not receive message in time")
	}

	// Ensure all messages acked
	pending, err := cons.Status(ctx)
	s.Require().NoError(err)
	s.Require().Zero(pending)
}

func (s *NATSSuite) TestNATSMessagingPullConsume() {
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()

	streamCfg := testStreamConfig()
	log := logger.NewFMTLogger(nil)

	pubCfg := &NATSPublisherConfig{
		ServiceName:         "publisher-service",
		NATSClusterHost:     s.natsURL,
		NATSClusterUser:     testNATSUser,
		NATSClusterPassword: testNATSPassword,
		NATSStreamConfig:    streamCfg,
	}
	consCfg := &NATSConsumerConfig{
		ServiceName:         "consumer-service-pull",
		Destination:         testDestination,
		NATSClusterHost:     s.natsURL,
		NATSClusterUser:     testNATSUser,
		NATSClusterPassword: testNATSPassword,
		MaxPendingMsgs:      10,
		NATSStreamConfig:    streamCfg,
	}

	pub, err := NewNATSPublisher(log, pubCfg, WithPublisherUpsertStream())
	s.Require().NoError(err)
	defer pub.Close(ctx)

	cons, err := NewNATSConsumer(log, consCfg, WitConsumerUpsertStream())
	s.Require().NoError(err)
	defer cons.Close(ctx)

	// Cleanup any stream data created by this test
	defer func() {
		natsCons := cons.(*natsConsumerImpl)
		err := natsCons.stream.DeleteConsumer(natsCons.streamName, natsCons.consumerGroup)
		s.Require().NoError(err)

		natsPub := pub.(*natsPublisherImpl)
		err = natsPub.stream.PurgeStream(streamCfg.Name)
		s.Require().NoError(err)
	}()

	payload := map[string]string{"idx": "0"}
	for i := 0; i < 3; i++ {
		payload["idx"] = string(rune('a' + i))
		_, err := pub.Publish(testSubject, testMessageType, payload)
		s.Require().NoError(err)
	}

	received, err := cons.Pull(testSubject, 3, func(msg *simplego.MessageWrapper) error {
		return nil
	})
	s.Require().NoError(err)

	filtered := filterNonNil(received)
	s.Require().Len(filtered, 3, "expected three acknowledged messages")

	for _, msg := range filtered {
		s.Require().Equal(testMessageType, msg.Type)
	}

	pending, err := cons.Status(ctx)
	s.Require().NoError(err)
	s.Require().Zero(pending)
}

func buildConfigs(natsURL string, streamCfg *NATSStreamConfig) (*NATSPublisherConfig, *NATSConsumerConfig) {
	pubCfg := &NATSPublisherConfig{
		ServiceName:         "publisher-service",
		NATSClusterHost:     natsURL,
		NATSClusterUser:     testNATSUser,
		NATSClusterPassword: testNATSPassword,
		NATSStreamConfig:    streamCfg,
	}
	consCfg := &NATSConsumerConfig{
		ServiceName:         "consumer-service",
		Destination:         testDestination,
		NATSClusterHost:     natsURL,
		NATSClusterUser:     testNATSUser,
		NATSClusterPassword: testNATSPassword,
		MaxPendingMsgs:      10,
		NATSStreamConfig:    streamCfg,
	}
	return pubCfg, consCfg
}

func testStreamConfig() *NATSStreamConfig {
	return &NATSStreamConfig{
		RetentionMaxMsgAge: DefaultNATSStreamMsgRetention,
		DeduplicateWindow:  DefaultNATSStreamDeduplicateWindow,
		Replicas:           DefaultNATSStreamReplicas,
		AllowedSubjects:    []string{testSubject},
		Name:               DefaultNATSStreamName,
		CompressionEnabled: false,
	}
}

func filterNonNil(msgs []*simplego.MessageWrapper) []*simplego.MessageWrapper {
	res := make([]*simplego.MessageWrapper, 0, len(msgs))
	for _, m := range msgs {
		if m != nil {
			res = append(res, m)
		}
	}
	return res
}
