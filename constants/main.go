package main

import (
	"fmt"
	"math"
	"time"
)

const name = "test"

func main() {
	const timeout = time.Second * 30
	fmt.Println("timeout: ", timeout)

	fmt.Println("name is ", name)

	const n = 100000

	const d = 3e20 / n
	fmt.Println(math.Sin(d))
}
