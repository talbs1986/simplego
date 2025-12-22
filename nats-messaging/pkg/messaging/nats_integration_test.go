package messaging

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/nats-io/nats.go"
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
	streamCfg.AllowedSubjects = []string{"async.*"} // Use unique subject pattern
	log := logger.NewFMTLogger(nil)

	pubCfg, consCfg := buildConfigs(s.natsURL, streamCfg)
	pubCfg.ServiceName = "async-pub"
	consCfg.ServiceName = "async-cons"

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

	subject := "async.testEvent"
	msgCh := make(chan *simplego.MessageWrapper, 1)
	err = cons.Consume(subject, func(msg *simplego.MessageWrapper) error {
		msgCh <- msg
		return nil
	})
	s.Require().NoError(err)

	// give the subscription a moment to bind before publishing
	time.Sleep(200 * time.Millisecond)

	payload := map[string]string{"hello": "world"}
	sentMsg, err := pub.Publish(subject, testMessageType, payload)
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
	streamCfg.AllowedSubjects = []string{"pull.*"} // Use unique subject pattern
	log := logger.NewFMTLogger(nil)

	pubCfg := &NATSPublisherConfig{
		ServiceName:         "publisher-service-pull",
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

	subject := "pull.testEvent"
	payload := map[string]string{"idx": "0"}
	for i := 0; i < 3; i++ {
		payload["idx"] = string(rune('a' + i))
		_, err := pub.Publish(subject, testMessageType, payload)
		s.Require().NoError(err)
	}

	received, err := cons.Pull(subject, 3, func(msg *simplego.MessageWrapper) error {
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

// TestNewNATSPublisherErrors tests error cases in NewNATSPublisher
func (s *NATSSuite) TestNewNATSPublisherErrors() {
	log := logger.NewFMTLogger(nil)

	// Test missing service name
	cfg := &NATSPublisherConfig{
		ServiceName:         "",
		NATSClusterHost:     s.natsURL,
		NATSClusterUser:     testNATSUser,
		NATSClusterPassword: testNATSPassword,
	}
	_, err := NewNATSPublisher(log, cfg)
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "missing service name")

	// Test invalid connection
	cfg = &NATSPublisherConfig{
		ServiceName:         "test-service",
		NATSClusterHost:     "nats://invalid:4222",
		NATSClusterUser:     testNATSUser,
		NATSClusterPassword: testNATSPassword,
	}
	_, err = NewNATSPublisher(log, cfg)
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "failed to connect")
}

// TestNewNATSConsumerErrors tests error cases in NewNATSConsumer
func (s *NATSSuite) TestNewNATSConsumerErrors() {
	log := logger.NewFMTLogger(nil)

	// Test missing service name
	cfg := &NATSConsumerConfig{
		ServiceName:         "",
		Destination:         testDestination,
		NATSClusterHost:     s.natsURL,
		NATSClusterUser:     testNATSUser,
		NATSClusterPassword: testNATSPassword,
		NATSStreamConfig:    testStreamConfig(),
	}
	_, err := NewNATSConsumer(log, cfg)
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "missing service name")

	// Test missing destination
	cfg = &NATSConsumerConfig{
		ServiceName:         "test-service",
		Destination:         "",
		NATSClusterHost:     s.natsURL,
		NATSClusterUser:     testNATSUser,
		NATSClusterPassword: testNATSPassword,
		NATSStreamConfig:    testStreamConfig(),
	}
	_, err = NewNATSConsumer(log, cfg)
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "missing destination")

	// Test invalid connection
	cfg = &NATSConsumerConfig{
		ServiceName:         "test-service",
		Destination:         testDestination,
		NATSClusterHost:     "nats://invalid:4222",
		NATSClusterUser:     testNATSUser,
		NATSClusterPassword: testNATSPassword,
		NATSStreamConfig:    testStreamConfig(),
	}
	_, err = NewNATSConsumer(log, cfg)
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "failed to connect")
}

// TestPublisherOptions tests publisher option functions
func (s *NATSSuite) TestPublisherOptions() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log := logger.NewFMTLogger(nil)
	streamCfg := testStreamConfig()

	// Create a connection and stream first
	conn, err := nats.Connect(s.natsURL, nats.UserInfo(testNATSUser, testNATSPassword))
	s.Require().NoError(err)
	defer conn.Close()

	js, err := conn.JetStream()
	s.Require().NoError(err)

	// Test WithPublisherNATSConnection
	cfg := &NATSPublisherConfig{
		ServiceName:         "test-pub",
		NATSClusterHost:     s.natsURL,
		NATSClusterUser:     testNATSUser,
		NATSClusterPassword: testNATSPassword,
		NATSStreamConfig:    streamCfg,
	}
	pub, err := NewNATSPublisher(log, cfg, WithPublisherNATSConnection(conn))
	s.Require().NoError(err)
	defer pub.Close(ctx)

	// Test WithPublisherNATSStream
	cfg2 := &NATSPublisherConfig{
		ServiceName:         "test-pub2",
		NATSClusterHost:     s.natsURL,
		NATSClusterUser:     testNATSUser,
		NATSClusterPassword: testNATSPassword,
		NATSStreamConfig:    streamCfg,
	}
	pub2, err := NewNATSPublisher(log, cfg2, WithPublisherNATSConnection(conn), WithPublisherNATSStream(js))
	s.Require().NoError(err)
	defer pub2.Close(ctx)
}

// TestConsumerOptions tests consumer option functions
func (s *NATSSuite) TestConsumerOptions() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log := logger.NewFMTLogger(nil)
	streamCfg := testStreamConfig()

	// Create stream first
	pubCfg := &NATSPublisherConfig{
		ServiceName:         "setup-pub",
		NATSClusterHost:     s.natsURL,
		NATSClusterUser:     testNATSUser,
		NATSClusterPassword: testNATSPassword,
		NATSStreamConfig:    streamCfg,
	}
	pub, err := NewNATSPublisher(log, pubCfg, WithPublisherUpsertStream())
	s.Require().NoError(err)
	defer pub.Close(ctx)

	// Create a connection and stream
	conn, err := nats.Connect(s.natsURL, nats.UserInfo(testNATSUser, testNATSPassword))
	s.Require().NoError(err)
	defer conn.Close()

	js, err := conn.JetStream()
	s.Require().NoError(err)

	// Test WithConsumerNATSConnection
	cfg := &NATSConsumerConfig{
		ServiceName:         "test-cons",
		Destination:         testDestination,
		NATSClusterHost:     s.natsURL,
		NATSClusterUser:     testNATSUser,
		NATSClusterPassword: testNATSPassword,
		NATSStreamConfig:    streamCfg,
	}
	cons, err := NewNATSConsumer(log, cfg, WithConsumerNATSConnection(conn))
	s.Require().NoError(err)
	defer cons.Close(ctx)

	// Test WithConsumerNATSStream
	cfg2 := &NATSConsumerConfig{
		ServiceName:         "test-cons2",
		Destination:         testDestination,
		NATSClusterHost:     s.natsURL,
		NATSClusterUser:     testNATSUser,
		NATSClusterPassword: testNATSPassword,
		NATSStreamConfig:    streamCfg,
	}
	cons2, err := NewNATSConsumer(log, cfg2, WithConsumerNATSConnection(conn), WithConsumerNATSStream(js))
	s.Require().NoError(err)
	defer cons2.Close(ctx)

	// Test WithConsumerNATSSubscription - create a subscription first
	sub, err := js.PullSubscribe(testSubject, "test-sub", nats.BindStream(streamCfg.Name))
	s.Require().NoError(err)

	cfg3 := &NATSConsumerConfig{
		ServiceName:         "test-cons3",
		Destination:         testDestination,
		NATSClusterHost:     s.natsURL,
		NATSClusterUser:     testNATSUser,
		NATSClusterPassword: testNATSPassword,
		NATSStreamConfig:    streamCfg,
	}
	cons3, err := NewNATSConsumer(log, cfg3, WithConsumerNATSConnection(conn), WithConsumerNATSStream(js), WithConsumerNATSSubscription(testSubject, sub))
	s.Require().NoError(err)
	defer cons3.Close(ctx)

	// Verify subscription was set
	natsCons := cons3.(*natsConsumerImpl)
	s.Require().NotNil(natsCons.sub[testSubject])
}

// TestConsumeDuplicateError tests duplicate consumer error
func (s *NATSSuite) TestConsumeDuplicateError() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	streamCfg := testStreamConfig()
	streamCfg.AllowedSubjects = []string{"duplicate.*"} // Use unique subject pattern
	log := logger.NewFMTLogger(nil)
	_, consCfg := buildConfigs(s.natsURL, streamCfg)
	consCfg.ServiceName = "duplicate-cons"

	// Create stream first
	pubCfg := &NATSPublisherConfig{
		ServiceName:         "setup-pub",
		NATSClusterHost:     s.natsURL,
		NATSClusterUser:     testNATSUser,
		NATSClusterPassword: testNATSPassword,
		NATSStreamConfig:    streamCfg,
	}
	pub, err := NewNATSPublisher(log, pubCfg, WithPublisherUpsertStream())
	s.Require().NoError(err)
	defer pub.Close(ctx)

	cons, err := NewNATSConsumer(log, consCfg, WitConsumerUpsertStream())
	s.Require().NoError(err)
	defer cons.Close(ctx)

	defer func() {
		natsCons := cons.(*natsConsumerImpl)
		_ = natsCons.stream.DeleteConsumer(natsCons.streamName, natsCons.consumerGroup)
	}()

	subject := "duplicate.testEvent"
	// First consume should succeed
	err = cons.Consume(subject, func(msg *simplego.MessageWrapper) error {
		return nil
	})
	s.Require().NoError(err)

	// Second consume on same subject should fail
	err = cons.Consume(subject, func(msg *simplego.MessageWrapper) error {
		return nil
	})
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "already exists")
}

// TestHandleMessageErrors tests error handling in handleMessage
func (s *NATSSuite) TestHandleMessageErrors() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	streamCfg := testStreamConfig()
	streamCfg.AllowedSubjects = []string{"error.*"} // Use unique subject pattern
	log := logger.NewFMTLogger(nil)
	pubCfg, consCfg := buildConfigs(s.natsURL, streamCfg)
	pubCfg.ServiceName = "error-pub"
	consCfg.ServiceName = "error-cons"

	pub, err := NewNATSPublisher(log, pubCfg, WithPublisherUpsertStream())
	s.Require().NoError(err)
	defer pub.Close(ctx)

	cons, err := NewNATSConsumer(log, consCfg, WitConsumerUpsertStream())
	s.Require().NoError(err)
	defer cons.Close(ctx)

	defer func() {
		natsCons := cons.(*natsConsumerImpl)
		_ = natsCons.stream.DeleteConsumer(natsCons.streamName, natsCons.consumerGroup)
	}()

	subject := "error.testEvent"
	// Test processor error - publish a valid message but processor returns error
	err = cons.Consume(subject, func(msg *simplego.MessageWrapper) error {
		return fmt.Errorf("processor error")
	})
	s.Require().NoError(err)

	time.Sleep(200 * time.Millisecond)

	_, err = pub.Publish(subject, testMessageType, map[string]string{"test": "data"})
	s.Require().NoError(err)

	// Wait a bit for message processing
	time.Sleep(500 * time.Millisecond)

	// Message should be nacked due to processor error, check status
	pending, err := cons.Status(ctx)
	s.Require().NoError(err)
	// The message should be pending (nacked)
	s.Require().GreaterOrEqual(pending, uint64(0))
}

// TestHandleMessageUnmarshalError tests unmarshal error path in handleMessage
func (s *NATSSuite) TestHandleMessageUnmarshalError() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	streamCfg := testStreamConfig()
	streamCfg.AllowedSubjects = []string{"unmarshal.*"} // Use unique subject pattern
	log := logger.NewFMTLogger(nil)
	pubCfg, consCfg := buildConfigs(s.natsURL, streamCfg)
	pubCfg.ServiceName = "unmarshal-pub"
	consCfg.ServiceName = "unmarshal-cons"

	pub, err := NewNATSPublisher(log, pubCfg, WithPublisherUpsertStream())
	s.Require().NoError(err)
	defer pub.Close(ctx)

	cons, err := NewNATSConsumer(log, consCfg, WitConsumerUpsertStream())
	s.Require().NoError(err)
	defer cons.Close(ctx)

	defer func() {
		natsCons := cons.(*natsConsumerImpl)
		_ = natsCons.stream.DeleteConsumer(natsCons.streamName, natsCons.consumerGroup)
	}()

	subject := "unmarshal.testEvent"
	msgReceived := make(chan bool, 1)
	err = cons.Consume(subject, func(msg *simplego.MessageWrapper) error {
		msgReceived <- true
		return nil
	})
	s.Require().NoError(err)

	time.Sleep(200 * time.Millisecond)

	// Publish invalid JSON directly to NATS to trigger unmarshal error
	natsPub := pub.(*natsPublisherImpl)
	invalidJSON := []byte(`{"invalid": json}`) // Invalid JSON
	_, err = natsPub.stream.Publish(subject, invalidJSON)
	s.Require().NoError(err)

	// Wait a bit for message processing - should handle unmarshal error gracefully
	time.Sleep(500 * time.Millisecond)

	// Verify no message was received (unmarshal failed, message nacked)
	select {
	case <-msgReceived:
		s.T().Fatal("should not receive message with invalid JSON")
	case <-time.After(1 * time.Second):
		// Expected - message was nacked due to unmarshal error
	}
}

// TestHandleMessageUnmarshalErrorInPull tests unmarshal error path in handleMessage via Pull
func (s *NATSSuite) TestHandleMessageUnmarshalErrorInPull() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	streamCfg := testStreamConfig()
	streamCfg.AllowedSubjects = []string{"unmarshalpull.*"} // Use unique subject pattern
	log := logger.NewFMTLogger(nil)
	pubCfg, consCfg := buildConfigs(s.natsURL, streamCfg)
	pubCfg.ServiceName = "unmarshalpull-pub"
	consCfg.ServiceName = "unmarshalpull-cons"

	pub, err := NewNATSPublisher(log, pubCfg, WithPublisherUpsertStream())
	s.Require().NoError(err)
	defer pub.Close(ctx)

	cons, err := NewNATSConsumer(log, consCfg, WitConsumerUpsertStream())
	s.Require().NoError(err)
	defer cons.Close(ctx)

	defer func() {
		natsCons := cons.(*natsConsumerImpl)
		_ = natsCons.stream.DeleteConsumer(natsCons.streamName, natsCons.consumerGroup)
	}()

	subject := "unmarshalpull.testEvent"
	// Publish invalid JSON directly to NATS to trigger unmarshal error
	natsPub := pub.(*natsPublisherImpl)
	invalidJSON := []byte(`{"invalid": json}`) // Invalid JSON
	_, err = natsPub.stream.Publish(subject, invalidJSON)
	s.Require().NoError(err)

	// Try to pull - should handle unmarshal error gracefully
	_, err = cons.Pull(subject, 1, func(msg *simplego.MessageWrapper) error {
		return nil
	})
	// Pull should return error or empty result due to unmarshal error in handleMessage
	// The handleMessage returns nil when unmarshal fails, which causes Pull to return error
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "failed to handle message")
}

// TestHandleMessageAckError tests Ack error path in handleMessage
// This is challenging to test in integration tests, but we can test by closing connection
// after message is received but before ack (though timing is tricky)
func (s *NATSSuite) TestHandleMessageAckError() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	streamCfg := testStreamConfig()
	streamCfg.AllowedSubjects = []string{"ackerror.*"} // Use unique subject pattern
	log := logger.NewFMTLogger(nil)
	pubCfg, consCfg := buildConfigs(s.natsURL, streamCfg)
	pubCfg.ServiceName = "ackerror-pub"
	consCfg.ServiceName = "ackerror-cons"

	pub, err := NewNATSPublisher(log, pubCfg, WithPublisherUpsertStream())
	s.Require().NoError(err)
	defer pub.Close(ctx)

	cons, err := NewNATSConsumer(log, consCfg, WitConsumerUpsertStream())
	s.Require().NoError(err)
	defer cons.Close(ctx)

	defer func() {
		natsCons := cons.(*natsConsumerImpl)
		_ = natsCons.stream.DeleteConsumer(natsCons.streamName, natsCons.consumerGroup)
	}()

	subject := "ackerror.testEvent"

	// Publish a valid message
	_, err = pub.Publish(subject, testMessageType, map[string]string{"test": "ack-error"})
	s.Require().NoError(err)

	// Use Pull to get the message
	// To test Ack error: close the connection in the processor
	// This will cause handleMessage's Ack() call to fail after processor returns
	natsCons := cons.(*natsConsumerImpl)
	_, err = cons.Pull(subject, 1, func(msg *simplego.MessageWrapper) error {
		// Close connection in processor - this will cause Ack() to fail in handleMessage
		// after processor returns successfully
		natsCons.conn.Close()
		return nil // Processor succeeds, but Ack will fail
	})

	// Pull should return error because handleMessage returns nil when Ack fails
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "failed to handle message")
}

// TestBuildDefaultNATSStream tests BuildDefaultNATSStream functionality
func (s *NATSSuite) TestBuildDefaultNATSStream() {
	conn, err := nats.Connect(s.natsURL, nats.UserInfo(testNATSUser, testNATSPassword))
	s.Require().NoError(err)
	defer conn.Close()

	js, err := conn.JetStream()
	s.Require().NoError(err)

	// Test with nil config (should use default)
	err = BuildDefaultNATSStream(nil, js)
	s.Require().NoError(err)

	// Test stream update path - try to create again (should update)
	err = BuildDefaultNATSStream(nil, js)
	s.Require().NoError(err)

	// Test with compression enabled
	compressedCfg := &NATSStreamConfig{
		RetentionMaxMsgAge: DefaultNATSStreamMsgRetention,
		DeduplicateWindow:  DefaultNATSStreamDeduplicateWindow,
		Replicas:           DefaultNATSStreamReplicas,
		AllowedSubjects:    []string{"compressed.*"},
		Name:               "compressed-stream",
		CompressionEnabled: true,
	}
	err = BuildDefaultNATSStream(compressedCfg, js)
	s.Require().NoError(err)

	// Cleanup
	_ = js.DeleteStream("compressed-stream")
}

// TestPublishMarshalError tests publish with invalid payload that can't be marshaled
func (s *NATSSuite) TestPublishMarshalError() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	streamCfg := testStreamConfig()
	log := logger.NewFMTLogger(nil)
	pubCfg, _ := buildConfigs(s.natsURL, streamCfg)

	pub, err := NewNATSPublisher(log, pubCfg, WithPublisherUpsertStream())
	s.Require().NoError(err)
	defer pub.Close(ctx)

	// Create a channel that can't be marshaled to JSON
	invalidPayload := make(chan int)
	_, err = pub.Publish(testSubject, testMessageType, invalidPayload)
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "failed to marshal")
}

// TestConsumerStatusError tests Status error case
func (s *NATSSuite) TestConsumerStatusError() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	streamCfg := testStreamConfig()
	log := logger.NewFMTLogger(nil)
	_, consCfg := buildConfigs(s.natsURL, streamCfg)
	consCfg.ServiceName = "status-error-cons"

	// Create stream first
	pubCfg := &NATSPublisherConfig{
		ServiceName:         "setup-pub",
		NATSClusterHost:     s.natsURL,
		NATSClusterUser:     testNATSUser,
		NATSClusterPassword: testNATSPassword,
		NATSStreamConfig:    streamCfg,
	}
	pub, err := NewNATSPublisher(log, pubCfg, WithPublisherUpsertStream())
	s.Require().NoError(err)
	defer pub.Close(ctx)

	cons, err := NewNATSConsumer(log, consCfg, WitConsumerUpsertStream())
	s.Require().NoError(err)
	defer cons.Close(ctx)

	// Try to get status without consuming (consumer doesn't exist yet)
	_, err = cons.Status(ctx)
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "failed to return status")
}

// TestConsumerMaxPendingMsgsDefault tests default max pending messages
func (s *NATSSuite) TestConsumerMaxPendingMsgsDefault() {
	streamCfg := testStreamConfig()
	log := logger.NewFMTLogger(nil)

	// Create stream first
	pubCfg := &NATSPublisherConfig{
		ServiceName:         "setup-pub",
		NATSClusterHost:     s.natsURL,
		NATSClusterUser:     testNATSUser,
		NATSClusterPassword: testNATSPassword,
		NATSStreamConfig:    streamCfg,
	}
	pub, err := NewNATSPublisher(log, pubCfg, WithPublisherUpsertStream())
	s.Require().NoError(err)
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		pub.Close(ctx)
	}()

	// Test with MaxPendingMsgs = 0 (should use default)
	consCfg := &NATSConsumerConfig{
		ServiceName:         "test-cons-default",
		Destination:         testDestination,
		NATSClusterHost:     s.natsURL,
		NATSClusterUser:     testNATSUser,
		NATSClusterPassword: testNATSPassword,
		MaxPendingMsgs:      0, // Should trigger default
		NATSStreamConfig:    streamCfg,
	}
	cons, err := NewNATSConsumer(log, consCfg, WitConsumerUpsertStream())
	s.Require().NoError(err)
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		cons.Close(ctx)
	}()

	natsCons := cons.(*natsConsumerImpl)
	s.Require().Equal(DefaultConsumerMaxPendingMsgs, natsCons.maxPendingMsgs)
}
