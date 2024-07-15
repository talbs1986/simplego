package metrics

import (
	"context"
	"time"
)

type mockPusher struct {
	
}

func (m *mockPusher) Push(ctx context.Context) error {
	return nil
}
func (m *mockPusher) Start(interval time.Duration) error {
	return nil
}
func (m *mockPusher) Close(ctx context.Context) error {
	return nil
}
