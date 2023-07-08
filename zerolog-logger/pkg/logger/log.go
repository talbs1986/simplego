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
	for k, v := range *fields {
		l.fields[k] = v
	}
	return l
}
func (l *zerologLog) WithErr(err error) simplego.LogLine {
	l.err = &err
	return l
}
func (l *zerologLog) Trace(msg string) {
	l.checkErrAndLogMsg(l.parent.underyling.Trace(), l.err, msg)
}
func (l *zerologLog) Debug(msg string) {
	l.checkErrAndLogMsg(l.parent.underyling.Trace(), l.err, msg)
}
func (l *zerologLog) Info(msg string) {
	l.checkErrAndLogMsg(l.parent.underyling.Trace(), l.err, msg)
}
func (l *zerologLog) Warn(msg string) {
	l.checkErrAndLogMsg(l.parent.underyling.Trace(), l.err, msg)
}
func (l *zerologLog) Error(err error, msg string) {
	l.checkErrAndLogMsg(l.parent.underyling.Trace(), &err, msg)
}
func (l *zerologLog) Fatal(err error, msg string) {
	l.checkErrAndLogMsg(l.parent.underyling.Fatal(), &err, msg)
}

func (l *zerologLog) checkErrAndLogMsg(underlyingEvent *zerolog.Event, err *error, msg string) {
	underlyingEvent.Fields(l.fields)
	if err != nil {
		underlyingEvent = underlyingEvent.Err(*l.err)
	}
	underlyingEvent.Msg(msg)
}
