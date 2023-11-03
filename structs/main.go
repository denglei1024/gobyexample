package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p := person{name: "John", age: 20}
	fmt.Println(p.name)
	p2 := &p
	p2.name = "John1"
	fmt.Println(p.name)
}
