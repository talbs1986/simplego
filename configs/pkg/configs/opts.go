package configs

func WithConfigParsers[T interface{}](p []ConfigParser[T]) ConfigsOpt[T] {
	return func(s *configsImpl[T]) {
		s.parsers = p
	}
}
