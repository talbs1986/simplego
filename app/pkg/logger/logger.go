package logger

import (
	"errors"
)

// LogFormat defines the log format type
type LogFormat string

// LogLevel defines the log level
type LogLevel string

const (
	// LogFormatJSON log format json
	LogFormatJSON LogFormat = "json"
	// LogFormatLogPrint log format standard print
	LogFormatLogPrint LogFormat = "log"

	// LogLevelTrace log lvl trace
	LogLevelTrace LogLevel = "trace"
	// LogLevelDebug log lvl debug
	LogLevelDebug LogLevel = "debug"
	// LogLevelInfo log lvl info
	LogLevelInfo LogLevel = "info"
	// LogLevelWarn log lvl warn
	LogLevelWarn LogLevel = "warn"
	// LogLevelError log lvl error
	LogLevelError LogLevel = "error"
	// LogLevelFatal log lvl fatal
	LogLevelFatal LogLevel = "fatal"
)

var (
	// ErrUnknownLogLevel unknown log level error
	ErrUnknownLogLevel = errors.New("unknown log level")
)

// ParseLogLevel parses log level type from string
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
