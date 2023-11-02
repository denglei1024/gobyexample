package main

import "fmt"

func main() {
	fc := closures()
	fc()
}

func closures() func() {
	x := 10
	return func() {
		fmt.Println(x)
	}
}
