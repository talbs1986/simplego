package logger

type ILogger interface {
	Log() LogLine
	With(fields *LogFields) LogLine
}
type LogFields map[string]interface{}

type LogLine interface {
	With(fields *LogFields) LogLine
	WithErr(err error) LogLine
	Trace(msg string)
	Debug(msg string)
	Info(msg string)
	Warn(err error, msg string)
	Error(err error, msg string)
	Fatal(err error, msg string)
}
