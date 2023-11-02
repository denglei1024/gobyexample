package main

import (
	"maps"
	"testing"
)

func TestCreateMap(t *testing.T) {
	// 使用 make 创建 map
	myMap := make(map[string]int)

	// 赋值
	myMap["key1"] = 1
	myMap["key2"] = 2
	t.Log(myMap)

	// 读取
	val1 := myMap["key1"]
	t.Log(val1)

	// 删除
	delete(myMap, "key1")
	t.Log(myMap)
}

func TestGetNotExists(t *testing.T) {
	myMap := map[string]int{"key1": 1, "key2": 2}
	value, b := myMap["key3"]
	t.Log(value, b)
}

func Test3(t *testing.T) {
	myMap1 := map[string]int{"key1": 1, "key2": 2}
	myMap2 := map[string]int{"key1": 1, "key2": 2}
	t.Log(maps.Equal(myMap1, myMap2))
}
