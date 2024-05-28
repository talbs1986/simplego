package configs

import (
	"context"
)

type ConfigParser[T interface{}] interface {
	Parse(context.Context) (*T, error)
	Get(context.Context) (*T, error)
}
