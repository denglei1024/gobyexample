package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println("Producing", i)
		ch <- i // 将数据发送到通道
		// time.Sleep(time.Second)
	}
	close(ch)
}

func consumer(ch chan int) {
	for {
		data, ok := <-ch
		if !ok {
			fmt.Println("Channel closed, exiting.")
			return
		}
		fmt.Println("Consuming", data)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	ch := make(chan int, 1) // 创建有缓冲的通道，容量为3

	go producer(ch)
	// go consumer(ch)

	time.Sleep(10 * time.Second)
}
