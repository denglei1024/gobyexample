package main

import "fmt"

func main() {
	a, b := vals()
	fmt.Println(a, b)
}

func vals() (int, int) {
	return 1, 2
}
