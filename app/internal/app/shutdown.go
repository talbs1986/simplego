package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/talbs1986/simplego/logger/pkg/logger"
)

type CloseableService interface {
	Close(ctx context.Context) error
}

func (s *App) WaitForShutodwn() {
	c := make(chan os.Signal, 1)
	sigs := []os.Signal{syscall.SIGTERM, syscall.SIGKILL}
	signal.Notify(c, sigs...)
	s.slog.With(&logger.LogFields{"signals": sigs}).Info("simplego app: waiting for app shutdown")
	<-c
}

func (s *App) WithCloseableServices(services ...CloseableService) {
	for _, c := range services {
		service := c
		s.closeableServices = append(s.closeableServices, service)
	}
}

func (s *App) Stop() {
	s.slog.With(&logger.LogFields{"total_closeable_services": len(s.closeableServices)}).Info("simplego app: stopping services")
	for _, closeable := range s.closeableServices {
		ctx, cancel := context.WithTimeout(s.ctx, s.stopTimeout)
		err := closeable.Close(ctx)
		cancel()
		if err != nil {
			s.slog.Error(err, "simplego app: failed to stop service")
		}
	}
}
