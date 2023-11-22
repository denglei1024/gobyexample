package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) Append(data int) {
	newNode := &Node{Data: data, Next: nil}

	if list.Head == nil {
		list.Head = newNode
		return
	}
	lastNode := list.Head
	for lastNode.Next != nil {
		lastNode = lastNode.Next
	}
	lastNode.Next = newNode
}

func (list *LinkedList) Print() {
	currentNode := list.Head
	for currentNode != nil {
		fmt.Printf("%d ", currentNode.Data)
		currentNode = currentNode.Next
	}
	fmt.Println()
}

func main() {
	// 创建链表并添加元素
	list := LinkedList{}
	list.Append(1)
	list.Append(3)
	list.Append(10)

	// 打印链表
	list.Print()
}
