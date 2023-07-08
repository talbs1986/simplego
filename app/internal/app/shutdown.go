package app

import (
	"os"
	"os/signal"
	"syscall"
)

func (s *App) WaitForShutodwn() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGKILL)
	<-c
}
