package _2_map_structure

import "testing"

func TestMapDeclare(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}

	var m2 = map[string]string{}
	m2["xiaoming"] = "小明"

	m3 := make(map[string]string, 10)
	m3["zhishi"] = "芝士"

	t.Log(m1, m2, m3)
}

func TestAccessNotExistingKey(t *testing.T) {
	// 访问不存在的key
	m1 := map[int]int{}
	t.Log(m1[1]) // 0

	if _, ok := m1[2]; ok {
		t.Log("key exists")
	} else {
		t.Log("key not exists")
	}

	m1[3] = 0
	if _, ok := m1[3]; ok {
		t.Log("key exists")
	} else {
		t.Log("key not exists")
	}
}

func TestTravelMap(t *testing.T) {
	m := map[int]int{1: 3, 2: 4, 3: 6}
	for k, v := range m {
		t.Log(k, v)
	}
}
