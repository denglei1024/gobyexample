package main

import (
	"context"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func(c context.Context) {
		t.Log("a")
		// select 会阻止执行
		select {
		case <-c.Done():
			t.Log("123")
		case <-time.After(3 * time.Second):
			t.Log("456")
		}
		t.Log("w")

	}(ctx)
	time.Sleep(4 * time.Second)
	cancel()
}
