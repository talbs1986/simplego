package metrics

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/talbs1986/simplego/app/pkg/logger"
	"github.com/talbs1986/simplego/metrics/pkg/metrics"
)

type promPusher struct {
	ctx        context.Context
	cancelCtx  context.CancelFunc
	logger     logger.ILogger
	underlying *push.Pusher
}

func newPusher(logger logger.ILogger, pushGatewayHost string, serviceName string) metrics.MetricsPusher {
	ctx, cancel := context.WithCancel(context.Background())
	s := &promPusher{
		ctx:        ctx,
		cancelCtx:  cancel,
		logger:     logger,
		underlying: push.New(pushGatewayHost, serviceName),
	}

	return s
}

// Push pushes the metrics to the gateway
func (s *promPusher) Push(ctx context.Context) error {
	return s.underlying.PushContext(ctx)
}

// Start starts an automatic pusher by the interval
func (s *promPusher) Start(pushInterval time.Duration) error {
	go func() {
		for {
			select {
			case <-s.ctx.Done():
				if err := s.flush(); err != nil {
					s.logger.Log().Fatal(err, "simplego metrics: pusher failed to flush metrics")
				}
			case <-time.After(pushInterval):
			}
		}
	}()
	return nil
}

// Close closes automatic pusher
func (s *promPusher) Close(ctx context.Context) error {
	s.cancelCtx()
	return nil
}

func (s *promPusher) flush() error {
	return s.underlying.Push()
}
