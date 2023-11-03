package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) SayHi() {
	fmt.Println("hello, ", p.name)
}

func main() {
	p := &person{name: "John", age: 12}
	p.SayHi()
}
