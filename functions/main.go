package main

import "fmt"

func main() {
	sum := sum(1, 2)
	fmt.Println(sum)
}

func sum(a int, b int) int {
	return a + b
}
