package logger

import (
	"time"

	simplego "github.com/talbs1986/simplego/logger/pkg/logger"
)

type TestkitLog struct {
	parent *testkitImpl
	Time   time.Time
	Fields simplego.LogFields
	Msg    string
	Lvl    simplego.LogLevel
	Err    *error
}

func (l *TestkitLog) With(fields *simplego.LogFields) simplego.LogLine {
	if fields == nil {
		return l
	}

	newFields := l.Fields
	for k, v := range *fields {
		newFields[k] = v
	}
	newLine := &TestkitLog{
		parent: l.parent,
		Fields: newFields,
		Err:    l.Err,
	}
	return newLine
}
func (l *TestkitLog) WithErr(err error) simplego.LogLine {
	l.Err = &err
	return l
}
func (l *TestkitLog) Trace(msg string) {
	l.checkErrAndLogMsg(simplego.LogLevelTrace, l.Err, msg)
}
func (l *TestkitLog) Debug(msg string) {
	l.checkErrAndLogMsg(simplego.LogLevelTrace, l.Err, msg)
}
func (l *TestkitLog) Info(msg string) {
	l.checkErrAndLogMsg(simplego.LogLevelTrace, l.Err, msg)
}
func (l *TestkitLog) Warn(msg string) {
	l.checkErrAndLogMsg(simplego.LogLevelTrace, l.Err, msg)
}
func (l *TestkitLog) Error(err error, msg string) {
	l.checkErrAndLogMsg(simplego.LogLevelTrace, &err, msg)
}
func (l *TestkitLog) Fatal(err error, msg string) {
	l.checkErrAndLogMsg(simplego.LogLevelTrace, &err, msg)
}

func (l *TestkitLog) checkErrAndLogMsg(lvl simplego.LogLevel, err *error, msg string) {
	l.Time = time.Now()
	l.Lvl = lvl
	l.Msg = msg
	l.Err = err
}
