package logger

import (
	"fmt"
	"os"
)

// Config defines the configuration object of the FMT logger
type Config struct {
	Level  *LogLevel
	Format *LogFormat
}

type fmtLogger struct {
	lvl LogLevel
}

// NewFMTLogger creates a new logger implemented by FMT
func NewFMTLogger(cfg *Config) ILogger {
	if cfg == nil {
		cfg = DefaultConfig
	}
	if cfg.Level == nil {
		cfg.Level = &DefaultLevel
	}
	return &fmtLogger{*cfg.Level}
}

// Log creates a new log line
func (s *fmtLogger) Log() LogLine {
	return &fmtLogLine{
		lvl: s.lvl,
	}
}

// With creates a new log line with the fields
func (s *fmtLogger) With(fields *LogFields) LogLine {
	return s.Log().With(fields)
}

type fmtLogLine struct {
	fields *LogFields
	lvl    LogLevel
	err    error
}

// With appends log fields to the current log line
func (l *fmtLogLine) With(fields *LogFields) LogLine {
	if fields == nil {
		return l
	}
	if l.fields == nil {
		l.fields = &LogFields{}
	}
	for k, v := range *fields {
		(*l.fields)[k] = v
	}
	return l
}

// WithErr appends an error to the log line
func (l *fmtLogLine) WithErr(err error) LogLine {
	l.err = err
	return l
}

// Trace writes a trace log line
func (l *fmtLogLine) Trace(msg string) {
	if l.skipLvl(LogLevelTrace) {
		return
	}
	imsg := fmt.Sprintf("[Trace] {%v}: %s", l.fields, msg)
	fmt.Fprintln(os.Stdout, imsg)
}

// //Debug writes a debug log line
func (l *fmtLogLine) Debug(msg string) {
	if l.skipLvl(LogLevelDebug) {
		return
	}
	imsg := fmt.Sprintf("[Debug] {%v}: %s", l.fields, msg)
	fmt.Fprintln(os.Stdout, imsg)
}

// Info writes an info log line
func (l *fmtLogLine) Info(msg string) {
	if l.skipLvl(LogLevelInfo) {
		return
	}
	imsg := fmt.Sprintf("[Info] {%v}: %s", l.fields, msg)
	fmt.Fprintln(os.Stdout, imsg)
}

// Warn writes a warning log line
func (l *fmtLogLine) Warn(err error, msg string) {
	if l.skipLvl(LogLevelWarn) {
		return
	}
	var actualErr = err
	if l.err != nil {
		actualErr = fmt.Errorf("%s , with err: %s", err, l.err)
	}
	imsg := fmt.Sprintf("[Warn] {%v}: %s, due to: %s", l.fields, msg, actualErr)
	fmt.Fprintln(os.Stdout, imsg)
}

// Error writes an error log line
func (l *fmtLogLine) Error(err error, msg string) {
	if l.skipLvl(LogLevelError) {
		return
	}
	var actualErr = err
	if l.err != nil {
		actualErr = fmt.Errorf("%s , with err: %s", err, l.err)
	}
	imsg := fmt.Sprintf("[Error] {%v}: %s, due to: %s", l.fields, msg, actualErr)
	fmt.Fprintln(os.Stderr, imsg)
}

// Fatal writes a fatal log line
func (l *fmtLogLine) Fatal(err error, msg string) {
	var actualErr = err
	if l.err != nil {
		actualErr = fmt.Errorf("%s , with err: %s", err, l.err)
	}
	imsg := fmt.Sprintf("[Fatal] {%v}: %s, due to: %s", l.fields, msg, actualErr)
	fmt.Fprintln(os.Stderr, imsg)
	panic(imsg)
}

func (l *fmtLogLine) skipLvl(lvl LogLevel) bool {
	return lvl < l.lvl
}
