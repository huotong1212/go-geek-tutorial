package _3_map_set

import "testing"

func TestMapWithFunction(t *testing.T) {
	m1 := map[int]func(op int) int{}
	m1[1] = func(op int) int {
		return op
	}
	m1[2] = func(op int) int {
		return op * op
	}
	m1[3] = func(op int) int {
		return op * op * op
	}

	t.Log(m1[1](3), m1[2](3), m1[3](3))
}

func TestSetWithMap(t *testing.T) {
	mySet := map[string]bool{}
	// 添加元素
	mySet["China"] = true
	// 判断元素是否存在
	t.Log(mySet["China"])
	t.Log(mySet["America"])

	// 去除元素
	delete(mySet, "China")
	t.Log(mySet)

	// 获取元素的个数
	t.Log("before append:", len(mySet))
	mySet["French"] = true
	t.Log("after append:", len(mySet))
}
