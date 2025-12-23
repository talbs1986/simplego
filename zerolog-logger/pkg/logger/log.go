package logger

import (
	"fmt"

	"github.com/rs/zerolog"
	simplego "github.com/talbs1986/simplego/app/pkg/logger"
)

type zerologLog struct {
	parent *zerologImpl
	fields simplego.LogFields
	err    *error
}

// With creates a new log line and appends the fields
func (l *zerologLog) With(fields *simplego.LogFields) simplego.LogLine {
	if fields == nil {
		return l
	}

	newFields := l.fields
	for k, v := range *fields {
		newFields[k] = v
	}
	newLine := &zerologLog{
		parent: l.parent,
		fields: newFields,
		err:    l.err,
	}
	return newLine
}

// WithErr appends an error to the log line
func (l *zerologLog) WithErr(err error) simplego.LogLine {
	l.err = &err
	return l
}
func (l *zerologLog) Trace(msg string, args ...any) {
	l.checkErrAndLogMsg(l.parent.underlying.Trace(), l.err, msg, args...)
}
func (l *zerologLog) Debug(msg string, args ...any) {
	l.checkErrAndLogMsg(l.parent.underlying.Debug(), l.err, msg, args...)
}
func (l *zerologLog) Info(msg string, args ...any) {
	l.checkErrAndLogMsg(l.parent.underlying.Info(), l.err, msg, args...)
}
func (l *zerologLog) Warn(err error, msg string, args ...any) {
	l.checkErrAndLogMsg(l.parent.underlying.Warn(), &err, msg, args...)
}
func (l *zerologLog) Error(err error, msg string, args ...any) {
	l.checkErrAndLogMsg(l.parent.underlying.Error(), &err, msg, args...)
}
func (l *zerologLog) Fatal(err error, msg string, args ...any) {
	l.checkErrAndLogMsg(l.parent.underlying.Fatal(), &err, msg, args...)
}

func (l *zerologLog) checkErrAndLogMsg(underlyingEvent *zerolog.Event, err *error, msg string, args ...any) {
	underlyingEvent.Fields(map[string]interface{}(l.fields))
	if err != nil {
		l.err = err
		underlyingEvent = underlyingEvent.Err(*l.err)
	}
	actualMsg := msg
	if len(args) > 0 {
		actualMsg = fmt.Sprintf(msg, args...)
	}
	underlyingEvent.Msg(actualMsg)
}
