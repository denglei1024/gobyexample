package main

import (
	"fmt"
	"testing"
	"time"
	"unicode"
)

// 计算字符数
func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	fmt.Printf("count is %d\n", count)
	return count
}

func TestGoRoutine1(t *testing.T) {
	go countDigits("1234567890")
	go countDigits("abcdefghij")
	go countDigits("ABCDEFGHIJ")
	go countDigits("1234567890abcdefghijABCDEFGHIJ")
}

func TestGoRoutine2(t *testing.T) {
	go countDigits("1234567890")
	go countDigits("abcdefghij")
	go countDigits("ABCDEFGHIJ")
	go countDigits("1234567890abcdefghijABCDEFGHIJ")
	time.Sleep(3 * time.Second)
}

type field struct{ name string }

func (p *field) print() {
	fmt.Printf("%p %s \n", p, p.name)
}

func Test1(t *testing.T) {
	data1 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data1 {
		go (*field).print(v)
	}

	data2 := []field{{"four"}, {"five"}, {"six"}}
	for _, v := range data2 {
		s := v
		go (*field).print(&s)
	}
	time.Sleep(3 * time.Second)
}
func a(a, b int) (x, y int) {
	defer func() {
		x = a * 10
		y = b * 20
		fmt.Printf("x = %d, y = %d\n", x, y)
	}()
	x = a + 10
	y = b + 20
	a = a - 10
	b = b - 20
	return
}

func Test2(t *testing.T) {
	fmt.Println(a(1, 2))
}

type Person struct {
	Name string
}

// 值拷贝示例
func modifyValue(val int) {
	val = 10
}

func modifyName(person Person) {
	person.Name = "Alice"
}

// 引用传递示例（使用指针）
func modifyReference(person *Person) {
	person.Name = "Alice"
}

func Test4(t *testing.T) {
	// 值传递示例
	a := 5
	modifyValue(a)
	fmt.Println("值拷贝:", a) // 输出: 值拷贝: 5

	person1 := Person{Name: "Bob"}
	modifyName(person1)
	fmt.Println("值传递:", person1.Name) // 输出: 引用传递: Bob

	// 引用传递示例
	person2 := &Person{Name: "Bob"}
	modifyReference(person2)
	fmt.Println("引用传递:", person2.Name) // 输出: 引用传递: Alice
}
