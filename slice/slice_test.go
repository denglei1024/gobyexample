package main

import "testing"

// 直接初始化切片
func TestSliceCreate1(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	t.Log(slice)
}

// 使用 make 函数创建切片，指定长度和容量
func TestSliceCreate2(t *testing.T) {
	slice1 := make([]int, 3, 5) // 长度为3，容量为5
	t.Log(slice1)

	slice2 := make([]int, 5) // 长度和容量都是5
	t.Log(slice2)
}

// 通过切片表达式，可以从数组或切片中创建一个新的切片
func TestSliceCreate3(t *testing.T) {
	arr := [4]int{1, 2, 3, 4}
	slice := arr[1:4]

	t.Log(arr)
	t.Log(slice)
}

// 使用`append`函数可以向切片动态添加元素
func TestSliceCreate4(t *testing.T) {
	slice1 := []int{1, 2, 3, 4}
	slice2 := append(slice1, 5, 6, 7)
	t.Log(slice1)
	t.Log(slice2)
}

// 通过使用`copy`函数，你可以从一个切片中复制元素到另一个切片，从而创建一个新的切片
func TestSliceCreate5(t *testing.T) {
	slice1 := []int{1, 2, 3, 4}
	slice2 := make([]int, len(slice1))
	copy(slice2, slice1)
	t.Log(slice2)
}

// 可以使用`[:]`操作符从数组或切片中创建一个新的切片
func TestSliceCreate6(t *testing.T) {
	arr := [4]int{1, 2, 3, 4}
	slice := arr[:]
	t.Log(slice)
}

func TestSliceAppend(t *testing.T) {
	slice := make([]int, 0)
	t.Logf("len:%d cap:%d slice:%v", len(slice), cap(slice), slice)
	for i := 0; i < 100; i++ {
		slice = append(slice, i)
		t.Logf("len:%d cap:%d slice:%v", len(slice), cap(slice), slice)
	}
}
