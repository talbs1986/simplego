package configs

import (
	"context"
	"fmt"

	envconfig "github.com/sethvargo/go-envconfig"
	simplego "github.com/talbs1986/simplego/configs/pkg/configs"
)

type goenvConfigParserImpl[T interface{}] struct {
	instance *T
}

func NewGoEnvConfigParser[T interface{}]() (simplego.ConfigParser[T], error) {
	instance := &goenvConfigParserImpl[T]{}

	return instance, nil
}

func (s *goenvConfigParserImpl[T]) Parse(ctx context.Context, obj *T) (*T, error) {
	// env vars change only on app boot
	if s.instance != nil {
		return s.instance, nil
	}

	if err := envconfig.Process(ctx, obj); err != nil {
		return nil, fmt.Errorf("simplego configs: parser failed to parse env vars, due to: %w", err)
	}
	s.instance = obj
	return s.instance, nil
}
