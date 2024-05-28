package configs

import (
	"context"

	"testing"

	"github.com/stretchr/testify/assert"
)

type mockConfig struct {
	Something string `env:"SOMETHING,default=imnotempty"`
}

func Test_ParseConfigEnvVars(t *testing.T) {
	instance, err := NewGoEnvConfigParser[mockConfig]()
	assert.Nil(t, err)

	actual, err := instance.Parse(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, "imnotempty", actual.Something)
}

func Test_ParseConfigEnvVarsDoesntParseEnvVarsTwice(t *testing.T) {
	instance, err := NewGoEnvConfigParser[mockConfig]()
	assert.Nil(t, err)

	expected, _ := instance.Parse(context.Background())
	actual, _ := instance.Parse(context.Background())
	assert.Equal(t, expected, actual)
}
