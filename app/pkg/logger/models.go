package logger

// ILogger defines the api for the logger
type ILogger interface {
	Log() LogLine
	With(fields *LogFields) LogLine
}

// LogFields defines the log fields type
type LogFields map[string]interface{}

// LogLine defines the api for a log line
type LogLine interface {
	// With creates a new log line with the fields
	With(fields *LogFields) LogLine
	// WithErr appends an error to the log line
	WithErr(err error) LogLine
	// Trace writes a trace log line
	Trace(msg string)
	// Debug writes a debug log line
	Debug(msg string)
	// Info writes an info log line
	Info(msg string)
	// Warn writes a warn log line
	Warn(err error, msg string)
	// Error writes an error log line
	Error(err error, msg string)
	// Fatal writes a fatal log line
	Fatal(err error, msg string)
}
