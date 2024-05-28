package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/talbs1986/simplego/app/pkg/logger"
)

func Test_LoggerGetShouldBeSingleton(t *testing.T) {
	instance, err := NewSimpleZerolog(logger.DefaultConfig)
	assert.Nil(t, err)
	assert.Equal(t, instance, instance.Get())
}

func Test_LoggerLogShouldCreateNewLogLine(t *testing.T) {
	instance, err := NewSimpleZerolog(logger.DefaultConfig)
	assert.Nil(t, err)
	line := instance.Log()
	actualParent := instance.(*zerologImpl)
	actual := line.(*zerologLog)
	assert.Nil(t, actual.err)
	assert.Empty(t, actual.fields)
	assert.Equal(t, actualParent, actual.parent)
}

func Test_LoggeWithShouldCreateNewLogLineWithInputFields(t *testing.T) {
	instance, err := NewSimpleZerolog(logger.DefaultConfig)
	assert.Nil(t, err)
	expectedFields := logger.LogFields{}
	line := instance.With(&expectedFields)
	actualParent := instance.(*zerologImpl)
	actual := line.(*zerologLog)
	assert.Nil(t, actual.err)
	assert.Equal(t, expectedFields, actual.fields)
	assert.Equal(t, actualParent, actual.parent)
}
