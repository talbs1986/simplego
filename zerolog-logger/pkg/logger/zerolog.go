package logger

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	simplego "github.com/talbs1986/simplego/logger/pkg/logger"
)

var (
	instance simplego.ILogger
)

type zerologImpl struct {
	underyling zerolog.Logger
}

func NewSimpleZerolog(cfg *simplego.Config) (simplego.ILogger, error) {
	if instance != nil {
		return instance, nil
	}
	if cfg == nil {
		cfg = simplego.DefaultConfig
	}
	impl := &zerologImpl{}
	if cfg.Level == nil {
		cfg.Level = &simplego.DefaultLevel
	}
	lvl, err := zerolog.ParseLevel(string(*cfg.Level))
	if err != nil {
		return nil, fmt.Errorf("simplego logger: failed to parse level '%s' to zerolog level, due to: %w", *cfg.Level, err)
	}
	if cfg.Format == nil {
		cfg.Format = &simplego.DefaultFormat
	}
	impl.underyling = zerolog.New(os.Stdout).With().Timestamp().Logger().Level(lvl)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if *cfg.Format == simplego.LogFormatLogPrint {
		impl.underyling = zerolog.Logger(impl.underyling).Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}
	instance = impl
	return instance, nil
}

func (s *zerologImpl) Get() simplego.ILogger {
	return instance
}

func (s *zerologImpl) Log() simplego.LogLine {
	l := &zerologLog{
		parent: s,
		fields: simplego.LogFields{},
	}
	return l
}
func (s *zerologImpl) With(fields *simplego.LogFields) simplego.LogLine {
	l := &zerologLog{
		parent: s,
	}
	if fields == nil {
		l.fields = simplego.LogFields{}
		return l
	}
	l.fields = *fields
	return l
}
