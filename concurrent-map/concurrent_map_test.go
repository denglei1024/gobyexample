package concurrent_map

import (
	"fmt"
	"sync"
	"testing"
)

func TestUnsafeMap(t *testing.T) {
	// 创建一个线程不安全的map
	myMap := make(map[int]int)

	// 创建一个WaitGroup用于等待所有goroutine完成
	var wg sync.WaitGroup

	// 启动多个goroutine并发访问map
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			myMap[i] = i * 10                               // 向map中写入数据
			fmt.Printf("Key: %d, Value: %d\n", i, myMap[i]) // 读取map中的数据
		}(i)
	}

	// 等待所有goroutine完成
	wg.Wait()

	// 检查map中的值是否正确（这一步是为了模拟测试）
	for i := 0; i < 5; i++ {
		if myMap[i] != i*10 {
			t.Errorf("Expected %d but got %d", i*10, myMap[i])
		}
	}
}

type SaleMap struct {
	mu sync.Mutex
	m  map[int]int
}

func NewSaleMap() *SaleMap {
	return &SaleMap{m: make(map[int]int)}
}

func (s *SaleMap) Set(key int, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func (s *SaleMap) Get(key int) (int, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	value, ok := s.m[key]
	return value, ok
}

func TestConcurrentMap(t *testing.T) {
	// 创建一个线程安全的map
	safeMap := NewSaleMap()
	var wg sync.WaitGroup
	// 启动多个goroutine并发写入和读取数据
	for i := 0; i < 500000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			safeMap.Set(i, i*10) // 向map中写入数据
		}(i)
	}
	// 等待所有goroutine完成
	wg.Wait()

	// 检查map中的值是否正确（这一步是为了模拟测试）
	for i := 0; i < 500000; i++ {
		if value, _ := safeMap.Get(i); value != i*10 {
			t.Errorf("Expected %d but got %d", i*10, value)
		}
	}
}
