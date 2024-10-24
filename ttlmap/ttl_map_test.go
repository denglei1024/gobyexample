package ttlmap

import (
	"sync"
	"testing"
	"time"
)

type Item struct {
	value    interface{}
	expireAt int64
}

type TTLMap struct {
	m  map[string]*Item
	mu sync.Mutex
}

func NewTTLMap(size int) (m *TTLMap) {
	m = &TTLMap{m: make(map[string]*Item, size)}
	go func() {
		for now := range time.Tick(time.Second) {
			m.mu.Lock()
			for k, v := range m.m {
				if v.expireAt <= now.Unix() {
					delete(m.m, k)
				}
			}
			m.mu.Unlock()
		}
	}()
	return
}

func (m *TTLMap) Set(key string, value interface{}, ttl int64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m[key] = &Item{value: value, expireAt: time.Now().Unix() + ttl}
}

func (m *TTLMap) Get(key string) (v interface{}, ok bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if item, ok := m.m[key]; ok && item.expireAt > time.Now().Unix() {
		return item.value, true
	}
	return nil, false
}

func (m *TTLMap) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.m[key]; ok {
		delete(m.m, key)
	}
}

func TestTTLMap(t *testing.T) {
	m := NewTTLMap(10)
	m.Set("hello", "world", 5)
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		if v, ok := m.Get("hello"); ok {
			t.Log(v)
		} else {
			t.Log("expired")
		}
	}
}
