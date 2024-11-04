package main

import (
	"testing"
	"time"
)

func run(done chan int) {
	time.Sleep(time.Second * 3)
	done <- 1
}

func TestChannel(t *testing.T) {
	done := make(chan int)
	go run(done)
	t.Logf("%d", <-done)
	t.Logf("done")
}
