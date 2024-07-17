package scenario

import (
	simplego "github.com/talbs1986/simplego/app/pkg/app"
)

var (
	instance *TestkitLogger
)

// TestkitLogger defines a testkit logger service
type TestkitLogger struct {
}

// NewLoggerTestkit creates a new logger test kit
func NewJobApp() (simplego.ILogger, error) {
	if instance != nil {
		return instance, nil
	}
	instance = &TestkitLogger{}
	return instance, nil
}

func (s *TestkitLogger) Log() simplego.LogLine {
	l := &TestkitLog{
		parent: s,
		Fields: simplego.LogFields{},
	}
	return l
}
func (s *TestkitLogger) With(fields *simplego.LogFields) simplego.LogLine {
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
