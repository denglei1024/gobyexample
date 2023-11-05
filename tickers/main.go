package main

import (
	"fmt"
	"time"
)

func main() {
	ticket := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case b := <-done:
				fmt.Println("b is", b)
				return
			case t := <-ticket.C:
				fmt.Println("tick at ", t)
			}
		}
	}()

	time.Sleep(2600 * time.Millisecond)
	ticket.Stop()
	done <- true
	fmt.Println("ticket stopped")
}
