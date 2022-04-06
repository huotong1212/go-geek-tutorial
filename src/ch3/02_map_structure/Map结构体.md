## go语言中的Map结构体

### Map的声明和初始化
```go
func TestMapDeclare(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}

	var m2 = map[string]string{}
	m2["xiaoming"] = "小明"

	m3 := make(map[string]string, 10) // 10表示cap容量
	m3["zhishi"] = "芝士"

	t.Log(m1, m2, m3)
}
```
注意：
- m1,m2是常用的两种声明方式
- 因为map的cap也会根据map中增加的元素个数来动态扩容，所以make map的操作常用于性能优化
- cap()方法不能对map使用,m1的len为3,m2初始化时len为0，m3也为0
- 因为len表示存储的元素的个数，而map在使用make创建时无法指定元素，所以10为cap

### 访问Map中不存在的元素
```go
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
```
注意：
- 和其他语言不同，当访问map中不存在的key时，返回的值为值类型的默认值而不是nil
- 可以通过ok来判断key是否存在

### Map的遍历
```go
func TestTravelMap(t *testing.T) {
	m := map[int]int{1: 3, 2: 4, 3: 6}
	for k, v := range m {
		t.Log(k, v)
	}
}
```
输出：
```go
=== RUN   TestTravelMap
    map_test.go:39: 1 3
    map_test.go:39: 2 4
    map_test.go:39: 3 6
--- PASS: TestTravelMap (0.00s)
```