package logger

import (
	"errors"

	simplego "github.com/talbs1986/simplego/logger/pkg/logger"
)

var (
	instance *testkitImpl
)

type testkitImpl struct {
}

func NewSimpleTestKit(cfg *simplego.Config) (simplego.ILogger, error) {
	if instance != nil {
		return instance, nil
	}
	if cfg == nil {
		cfg = simplego.DefaultConfig
	}
	instance := &testkitImpl{}
	if cfg.Level == nil {
		cfg.Level = &simplego.DefaultLevel
	}
	if cfg.Format == nil {
		cfg.Format = &simplego.DefaultFormat
	}

	if *cfg.Format == simplego.LogFormatLogPrint {
		return nil, errors.New("unsupported feature") // TODO
	}
	return instance, nil
}

func (s *testkitImpl) Get() simplego.ILogger {
	return instance
}

func (s *testkitImpl) Log() simplego.LogLine {
	l := &TestkitLog{
		parent: s,
		Fields: simplego.LogFields{},
	}
	return l
}
func (s *testkitImpl) With(fields *simplego.LogFields) simplego.LogLine {
	l := &TestkitLog{
		parent: s,
	}
	if fields == nil {
		l.Fields = simplego.LogFields{}
		return l
	}
	l.Fields = *fields
	return l
}
