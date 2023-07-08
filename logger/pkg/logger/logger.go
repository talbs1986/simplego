package logger

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
