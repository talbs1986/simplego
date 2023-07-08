package logger

type LogFormat string

const (
	LogFormatJSON     LogFormat = "json"
	LogFormatLogPrint LogFormat = "log"
)

type Config struct {
}
