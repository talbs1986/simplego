package logger

import (
	simplego "github.com/talbs1986/simplego/app/pkg/logger"
)

var (
	instance *testkitImpl
)

type testkitImpl struct {
}

func NewSimpleTestKit() (simplego.ILogger, error) {
	if instance != nil {
		return instance, nil
	}
	instance = &testkitImpl{}
	return instance, nil
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
