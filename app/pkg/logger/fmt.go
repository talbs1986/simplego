package logger

import (
	"fmt"
	"os"
)

type fmtLogger struct {
	lvl LogLevel
}

func NewFMTLogger(cfg *Config) ILogger {
	if cfg.Level == nil {
		cfg.Level = &DefaultLevel
	}
	return &fmtLogger{*cfg.Level}
}

func (s *fmtLogger) Get() ILogger {
	return s
}

func (s *fmtLogger) Log() LogLine {
	return &fmtLogLine{
		lvl: s.lvl,
	}
}
func (s *fmtLogger) With(fields *LogFields) LogLine {
	return s.Log().With(fields)
}

type fmtLogLine struct {
	fields *LogFields
	lvl    LogLevel
	err    error
}

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
func (l *fmtLogLine) WithErr(err error) LogLine {
	l.err = err
	return l
}

func (l *fmtLogLine) Trace(msg string) {
	if l.skipLvl(LogLevelTrace) {
		return
	}
	imsg := fmt.Sprintf("[Trace] {%v}: %s", l.fields, msg)
	fmt.Fprintln(os.Stdout, imsg)
}
func (l *fmtLogLine) Debug(msg string) {
	if l.skipLvl(LogLevelDebug) {
		return
	}
	imsg := fmt.Sprintf("[Debug] {%v}: %s", l.fields, msg)
	fmt.Fprintln(os.Stdout, imsg)
}
func (l *fmtLogLine) Info(msg string) {
	if l.skipLvl(LogLevelInfo) {
		return
	}
	imsg := fmt.Sprintf("[Info] {%v}: %s", l.fields, msg)
	fmt.Fprintln(os.Stdout, imsg)
}
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
