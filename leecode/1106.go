package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		var val1, val2 int
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		sum := val1 + val2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{sum, nil}
			tail = head
		} else {
			tail.Next = &ListNode{sum, nil}
			tail = tail.Next
		}
		if l1 == nil && l2 != nil {
			tail.Next = l2
			break
		}
		if l1 != nil && l2 == nil {
			tail.Next = l1
			break
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{carry, nil}
	}
	return
}

func main() {
	// l1 = [2,4,3], l2 = [5,6,4]
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	fmt.Println(addTwoNumbers(l1, l2))
}
