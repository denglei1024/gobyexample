package main

import "fmt"

func main() {
	for i := 3; i < 10; i++ {
		fmt.Println("i = ", i)
	}

	i := 3
	for i < 10 {
		fmt.Println("i = ", i)
		i++
	}
	for {
		fmt.Println("loop")
		break
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("i = ", i)
	}
}
