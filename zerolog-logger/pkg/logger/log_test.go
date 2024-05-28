package logger

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	simplego "github.com/talbs1986/simplego/app/pkg/logger"
)

func TestWithShouldCreateNewLogLineWithFields(t *testing.T) {
	instance, _ := NewSimpleZerolog(simplego.DefaultConfig)
	logLine := &zerologLog{
		parent: instance.(*zerologImpl),
		fields: simplego.LogFields{},
		err:    nil,
	}

	expectedFields := simplego.LogFields{
		"stam": 2,
	}
	actual := logLine.With(&expectedFields).(*zerologLog)

	assert.Nil(t, actual.err)
	assert.Equal(t, expectedFields, actual.fields)
	assert.Equal(t, instance, actual.parent)
}

func TestWithShouldReturnSameLineWhenNilFields(t *testing.T) {
	instance, _ := NewSimpleZerolog(simplego.DefaultConfig)
	logLine := &zerologLog{
		parent: instance.(*zerologImpl),
		fields: simplego.LogFields{},
		err:    nil,
	}

	actual := logLine.With(nil).(*zerologLog)
	assert.Equal(t, actual, logLine)
}
func TestWithErr(t *testing.T) {
	logLine := &zerologLog{
		parent: nil,
		fields: simplego.LogFields{},
		err:    nil,
	}

	expected := errors.New("this is a real error")
	actual := logLine.WithErr(expected).(*zerologLog)
	assert.Equal(t, actual, logLine)
	assert.Equal(t, actual.err, &expected)
}
func TestTrace(t *testing.T) {
	instance, _ := NewSimpleZerolog(simplego.DefaultConfig)
	expectedErr := errors.New("new error")
	expectedFields := simplego.LogFields{
		"stam": 3,
	}
	logLine := &zerologLog{
		parent: instance.(*zerologImpl),
		fields: expectedFields,
		err:    &expectedErr,
	}

	logLine.Trace("some kind of msg")
	assert.Equal(t, logLine.err, &expectedErr)
	assert.Equal(t, logLine.fields, expectedFields)
}
func TestDebug(t *testing.T) {
	instance, _ := NewSimpleZerolog(simplego.DefaultConfig)
	expectedErr := errors.New("new error")
	expectedFields := simplego.LogFields{
		"stam": 3,
	}
	logLine := &zerologLog{
		parent: instance.(*zerologImpl),
		fields: expectedFields,
		err:    &expectedErr,
	}

	logLine.Debug("some kind of msg")
	assert.Equal(t, logLine.err, &expectedErr)
	assert.Equal(t, logLine.fields, expectedFields)
}
func TestInfo(t *testing.T) {
	instance, _ := NewSimpleZerolog(simplego.DefaultConfig)
	expectedErr := errors.New("new error")
	expectedFields := simplego.LogFields{
		"stam": 3,
	}
	logLine := &zerologLog{
		parent: instance.(*zerologImpl),
		fields: expectedFields,
		err:    &expectedErr,
	}

	logLine.Info("some kind of msg")
	assert.Equal(t, logLine.err, &expectedErr)
	assert.Equal(t, logLine.fields, expectedFields)
}

func TestWarnWithErrShouldOverrideCurrentErr(t *testing.T) {
	instance, _ := NewSimpleZerolog(simplego.DefaultConfig)
	expectedErr := errors.New("new error")
	expectedFields := simplego.LogFields{
		"stam": 3,
	}
	someErr := errors.New("some error")
	logLine := &zerologLog{
		parent: instance.(*zerologImpl),
		fields: expectedFields,
		err:    &someErr,
	}

	logLine.Warn(expectedErr, "some kind of msg")
	assert.Equal(t, &expectedErr, logLine.err)
	assert.Equal(t, expectedFields, logLine.fields)
}

func TestErrorWithErrShouldOverrideCurrentErr(t *testing.T) {
	instance, _ := NewSimpleZerolog(simplego.DefaultConfig)
	expectedErr := errors.New("new error")
	expectedFields := simplego.LogFields{
		"stam": 3,
	}
	someErr := errors.New("some error")
	logLine := &zerologLog{
		parent: instance.(*zerologImpl),
		fields: expectedFields,
		err:    &someErr,
	}

	logLine.Error(expectedErr, "some kind of msg")
	assert.Equal(t, &expectedErr, logLine.err)
	assert.Equal(t, expectedFields, logLine.fields)
}
