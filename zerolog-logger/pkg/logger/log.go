package logger

import (
	"github.com/rs/zerolog"
	simplego "github.com/talbs1986/simplego/logger/pkg/logger"
)

type zerologLog struct {
	parent *zerologImpl
	fields simplego.LogFields
	err    *error
}

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
func (l *zerologLog) WithErr(err error) simplego.LogLine {
	l.err = &err
	return l
}
func (l *zerologLog) Trace(msg string) {
	l.checkErrAndLogMsg(l.parent.underlying.Trace(), l.err, msg)
}
func (l *zerologLog) Debug(msg string) {
	l.checkErrAndLogMsg(l.parent.underlying.Debug(), l.err, msg)
}
func (l *zerologLog) Info(msg string) {
	l.checkErrAndLogMsg(l.parent.underlying.Info(), l.err, msg)
}
func (l *zerologLog) Warn(err error, msg string) {
	l.checkErrAndLogMsg(l.parent.underlying.Warn(), &err, msg)
}
func (l *zerologLog) Error(err error, msg string) {
	l.checkErrAndLogMsg(l.parent.underlying.Error(), &err, msg)
}
func (l *zerologLog) Fatal(err error, msg string) {
	l.checkErrAndLogMsg(l.parent.underlying.Fatal(), &err, msg)
}

func (l *zerologLog) checkErrAndLogMsg(underlyingEvent *zerolog.Event, err *error, msg string) {
	underlyingEvent.Fields(l.fields)
	if err != nil {
		l.err = err
		underlyingEvent = underlyingEvent.Err(*l.err)
	}
	underlyingEvent.Msg(msg)
}
