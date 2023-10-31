package main

import "fmt"

func main() {
	//定义变量
	var a string

	fmt.Println("a = ", a)
	fmt.Println("a is empty string", a == "")

	b := "b"
	fmt.Println("b =", b)

	x, y := 1, 2
	fmt.Println("x = ", x, " y = ", y)
}
