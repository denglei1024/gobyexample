package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Key   int
	Value int
	Prev  *Node
	Next  *Node
}

type LRUCache struct {
	Capacity int
	Size     int
	Head     *Node
	Tail     *Node
	Cache    map[int]*Node
}

// NewLRUCache creates a new LRUCache
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		Capacity: capacity,
		Size:     0,
		Head:     nil,
		Tail:     nil,
		Cache:    make(map[int]*Node),
	}
}

// get cache by key
func (c *LRUCache) Get(key int) (int, error) {
	if node, ok := c.Cache[key]; ok {
		c.moveToHead(node)
		return node.Value, nil
	}
	return 0, errors.New("key不存在")
}

// put new value to cache
func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.Cache[key]; ok {
		node.Value = value
		c.moveToHead(node)
	} else {
		// create new node
		newNode := &Node{Value: value, Key: key}
		c.addToHead(newNode)

		if c.Capacity <= c.Size {
			c.removeTail()
		} else {
			c.Size++
		}
		c.Cache[key] = newNode
	}
}

// move to head
func (c *LRUCache) moveToHead(node *Node) {
	c.removeNode(node)
	c.addToHead(node)
}

// remove node
func (c *LRUCache) removeNode(node *Node) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		c.Head = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		c.Tail = node.Prev
	}
	node = nil
}

// add to head
func (c *LRUCache) addToHead(node *Node) {
	if c.Head == nil {
		c.Head = node
		c.Tail = node
	} else {
		node.Next = c.Head
		c.Head.Prev = node
		c.Head = node
	}
}

// remove tail node
func (c *LRUCache) removeTail() {
	if c.Tail != nil {
		delete(c.Cache, c.Tail.Key)
		c.removeNode(c.Tail)
		c.Size--
	}
}

// print lru
func (lru *LRUCache) PrintLRU() {
	currentNode := lru.Head
	for currentNode != nil {
		fmt.Printf("(%d,%d) ", currentNode.Key, currentNode.Value)
		currentNode = currentNode.Next
	}
	fmt.Println()
}

func main() {
	// 创建LRU缓存
	lruCache := NewLRUCache(3)

	// 插入数据
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	lruCache.Put(3, 3)

	// 打印当前LRU缓存
	lruCache.PrintLRU()

	// 查询缓存中的值
	v, _ := lruCache.Get(2)
	fmt.Println("Get(2):", v)

	// 将访问的节点移动到链表头部
	lruCache.PrintLRU()

	// 插入新值，触发淘汰
	lruCache.Put(4, 4)

	// 打印最新的LRU缓存
	lruCache.PrintLRU()
}
