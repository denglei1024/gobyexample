package main

import "fmt"

func main() {
	total := sum(1, 2, 3)
	fmt.Println(total)

	// 将切片的值展开为一个值列表
	total = sum([]int{1, 2, 3}...)
	fmt.Println(total)
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
