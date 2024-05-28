package logger

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	simplego "github.com/talbs1986/simplego/app/pkg/logger"
)

func TestWithShouldCreateNewLogLineWithFields(t *testing.T) {
	instance, _ := NewSimpleTestKit()
	logLine := &TestkitLog{
		parent: instance.(*testkitImpl),
		Fields: simplego.LogFields{},
		Err:    nil,
	}

	expectedFields := simplego.LogFields{
		"stam": 2,
	}
	actual := logLine.With(&expectedFields).(*TestkitLog)

	assert.Nil(t, actual.Err)
	assert.Nil(t, logLine.Time)
	assert.Equal(t, expectedFields, actual.Fields)
	assert.Equal(t, instance, actual.parent)
}

func TestWithShouldReturnSameLineWhenNilFields(t *testing.T) {
	instance, _ := NewSimpleTestKit()
	logLine := &TestkitLog{
		parent: instance.(*testkitImpl),
		Fields: simplego.LogFields{},
		Err:    nil,
	}

	actual := logLine.With(nil).(*TestkitLog)
	assert.Equal(t, actual, logLine)
}
func TestWithErr(t *testing.T) {
	logLine := &TestkitLog{
		parent: nil,
		Fields: simplego.LogFields{},
		Err:    nil,
	}

	expected := errors.New("this is a real error")
	actual := logLine.WithErr(expected).(*TestkitLog)
	assert.Equal(t, actual, logLine)
	assert.Equal(t, actual.Err, &expected)
}
func TestTrace(t *testing.T) {
	instance, _ := NewSimpleTestKit()
	expectedErr := errors.New("new error")
	expectedFields := simplego.LogFields{
		"stam": 3,
	}
	logLine := &TestkitLog{
		parent: instance.(*testkitImpl),
		Fields: expectedFields,
		Err:    &expectedErr,
	}

	expectedMsg := "some kind of msg"
	now := time.Now()
	logLine.Trace(expectedMsg)
	assert.LessOrEqual(t, now.Unix(), logLine.Time.Unix())
	assert.Equal(t, logLine.Msg, expectedMsg)
	assert.Equal(t, logLine.Err, &expectedErr)
	assert.Equal(t, logLine.Fields, expectedFields)
}
func TestDebug(t *testing.T) {
	instance, _ := NewSimpleTestKit()
	expectedErr := errors.New("new error")
	expectedFields := simplego.LogFields{
		"stam": 3,
	}
	logLine := &TestkitLog{
		parent: instance.(*testkitImpl),
		Fields: expectedFields,
		Err:    &expectedErr,
	}

	expectedMsg := "some kind of msg"
	now := time.Now()
	logLine.Debug(expectedMsg)
	assert.LessOrEqual(t, now.Unix(), logLine.Time.Unix())
	assert.Equal(t, logLine.Msg, expectedMsg)
	assert.Equal(t, logLine.Err, &expectedErr)
	assert.Equal(t, logLine.Fields, expectedFields)
}
func TestInfo(t *testing.T) {
	instance, _ := NewSimpleTestKit()
	expectedErr := errors.New("new error")
	expectedFields := simplego.LogFields{
		"stam": 3,
	}
	logLine := &TestkitLog{
		parent: instance.(*testkitImpl),
		Fields: expectedFields,
		Err:    &expectedErr,
	}

	expectedMsg := "some kind of msg"
	now := time.Now()
	logLine.Info(expectedMsg)
	assert.LessOrEqual(t, now.Unix(), logLine.Time.Unix())
	assert.Equal(t, logLine.Msg, expectedMsg)
	assert.Equal(t, logLine.Err, &expectedErr)
	assert.Equal(t, logLine.Fields, expectedFields)
}

func TestWarnWithErrShouldOverrideCurrentErr(t *testing.T) {
	instance, _ := NewSimpleTestKit()
	expectedErr := errors.New("new error")
	expectedFields := simplego.LogFields{
		"stam": 3,
	}
	someErr := errors.New("some error")
	logLine := &TestkitLog{
		parent: instance.(*testkitImpl),
		Fields: expectedFields,
		Err:    &someErr,
	}

	expectedMsg := "some kind of msg"
	now := time.Now()
	logLine.Warn(expectedErr, expectedMsg)
	assert.LessOrEqual(t, now.Unix(), logLine.Time.Unix())
	assert.Equal(t, logLine.Msg, expectedMsg)
	assert.Equal(t, &expectedErr, logLine.Err)
	assert.Equal(t, expectedFields, logLine.Fields)
}

func TestErrorWithErrShouldOverrideCurrentErr(t *testing.T) {
	instance, _ := NewSimpleTestKit()
	expectedErr := errors.New("new error")
	expectedFields := simplego.LogFields{
		"stam": 3,
	}
	someErr := errors.New("some error")
	logLine := &TestkitLog{
		parent: instance.(*testkitImpl),
		Fields: expectedFields,
		Err:    &someErr,
	}

	expectedMsg := "some kind of msg"
	now := time.Now()
	logLine.Error(expectedErr, expectedMsg)
	assert.LessOrEqual(t, now.Unix(), logLine.Time.Unix())
	assert.Equal(t, logLine.Msg, expectedMsg)
	assert.Equal(t, &expectedErr, logLine.Err)
	assert.Equal(t, expectedFields, logLine.Fields)
}
