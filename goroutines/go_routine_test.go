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
