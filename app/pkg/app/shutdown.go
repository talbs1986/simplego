package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/talbs1986/simplego/app/pkg/logger"
)

// WaitForShutodwn wait for the app to shutdown by  syscall.SIGTERM, syscall.SIGKILL
func (s *App) WaitForShutodwn() {
	c := make(chan os.Signal, 1)
	sigs := []os.Signal{syscall.SIGTERM, syscall.SIGKILL}
	signal.Notify(c, sigs...)
	s.slog.With(&logger.LogFields{"signals": sigs}).Info("simplego app: waiting for app shutdown")
	<-c
}

// Stop stops the application and the registered closeable services
func (s *App) Stop() {
	s.slog.With(&logger.LogFields{"total_closeable_services": len(s.closeableServices)}).Info("simplego app: stopping services")
	s.cancel()
	for _, closeable := range s.closeableServices {
		ctx, cancel := context.WithTimeout(s.CTX, s.stopTimeout)
		err := closeable.Close(ctx)
		cancel()
		if err != nil {
			s.slog.Error(err, "simplego app: failed to stop service")
		}
	}
}
