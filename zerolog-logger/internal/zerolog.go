package logger

import (
	simplego "github.com/talbs1986/simplego/logger/pkg/logger"
)

type zerologImpl struct {
}

func NewSimpleZerolog(cfg *simplego.Config) (simplego.ILogger, error) {
	s := &zerologImpl{}
	return s, nil
}

func (s *zerologImpl) Info() {

}
