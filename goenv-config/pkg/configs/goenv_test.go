package configs

import (
	"context"

	"github.com/stretchr/testify/assert"
	"testing"
)

type mockConfig struct {
	Something string `env:"SOMETHING,default=imnotempty"`
}

func Test_ParseConfigEnvVars(t *testing.T) {
	instance, err := NewGoEnvConfigParser[mockConfig]()
	assert.Nil(t, err)

	var cfg mockConfig
	actual, err := instance.Parse(context.Background(), &cfg)
	assert.Nil(t, err)
	assert.Equal(t, "imnotempty", actual.Something)
}

func Test_ParseConfigEnvVarsDoesntParseEnvVarsTwice(t *testing.T) {
	instance, err := NewGoEnvConfigParser[mockConfig]()
	assert.Nil(t, err)

	var cfg mockConfig
	expected, _ := instance.Parse(context.Background(), &cfg)
	actual, _ := instance.Parse(context.Background(), &cfg)
	assert.Equal(t, expected, actual)
}
