package main

import "testing"

func swap(a, b []int) {
	a, b = b, a
}

func TestSwap(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{6, 7, 8, 9, 10}
	swap(a[:], b[:])
	t.Logf("a: %v, b: %v", a, b)
}

func SliceSwap(a, b *[]int) {
	a, b = b, a
}

// func TestSliceSwap(t *testing.T) {
// 	a := [5]int{1, 2, 3, 4, 5}
// 	b := [5]int{6, 7, 8, 9, 10}
// 	SliceSwap(&a, &b)
// 	t.Logf("a: %v, b: %v", a, b)
// }
