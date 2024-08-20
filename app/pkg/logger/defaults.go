package logger

var (
	// DefaultLevel defines the default log level
	DefaultLevel LogLevel = LogLevelDebug
	// DefaultFormat defines the default log format
	DefaultFormat LogFormat = LogFormatJSON
	// DefaultConfig defines the default log config
	DefaultConfig *Config = &Config{
		Level:  &DefaultLevel,
		Format: &DefaultFormat,
	}
)
