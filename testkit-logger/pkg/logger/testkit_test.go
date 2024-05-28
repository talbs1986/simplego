package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/talbs1986/simplego/app/pkg/logger"
)

func Test_LoggerGetShouldBeSingleton(t *testing.T) {
	expected, err := NewSimpleTestKit()
	assert.Nil(t, err)
	actual, err := NewSimpleTestKit()
	assert.Nil(t, err)
	assert.Equal(t, actual, expected)
}

func Test_LoggerLogShouldCreateNewLogLine(t *testing.T) {
	instance, err := NewSimpleTestKit()
	assert.Nil(t, err)
	line := instance.Log()
	actualParent := instance.(*testkitImpl)
	actual := line.(*TestkitLog)
	assert.Nil(t, actual.Err)
	assert.Empty(t, actual.Fields)
	assert.Equal(t, actualParent, actual.parent)
}

func Test_LoggeWithShouldCreateNewLogLineWithInputFields(t *testing.T) {
	instance, err := NewSimpleTestKit()
	assert.Nil(t, err)
	expectedFields := logger.LogFields{}
	line := instance.With(&expectedFields)
	actualParent := instance.(*testkitImpl)
	actual := line.(*TestkitLog)
	assert.Nil(t, actual.Err)
	assert.Equal(t, expectedFields, actual.Fields)
	assert.Equal(t, actualParent, actual.parent)
}
