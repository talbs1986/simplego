package logger

import (
	"errors"
)

type LogFormat string
type LogLevel string

const (
	LogFormatJSON     LogFormat = "json"
	LogFormatLogPrint LogFormat = "log"

	LogLevelTrace LogLevel = "trace"
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"
	LogLevelFatal LogLevel = "fatal"
)

var (
	ErrUnknownLogLevel = errors.New("unknown log level")
)

func ParseLogLevel(lvl string) (LogLevel, error) {
	switch lvl {
	case string(LogLevelTrace):
		return LogLevelTrace, nil
	case string(LogLevelDebug):
		return LogLevelDebug, nil
	case string(LogLevelInfo):
		return LogLevelInfo, nil
	case string(LogLevelWarn):
		return LogLevelWarn, nil
	case string(LogLevelError):
		return LogLevelError, nil
	case string(LogLevelFatal):
		return LogLevelFatal, nil
	default:
		return "", ErrUnknownLogLevel
	}
}

var (
	DefaultLevel  LogLevel  = LogLevelDebug
	DefaultFormat LogFormat = LogFormatJSON
	DefaultConfig *Config   = &Config{
		Level:  &DefaultLevel,
		Format: &DefaultFormat,
	}
)

type Config struct {
	Level  *LogLevel
	Format *LogFormat
}
