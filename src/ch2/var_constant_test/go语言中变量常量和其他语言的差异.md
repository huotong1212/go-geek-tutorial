# go语言中变量常量和其他语言的差异

## 1.编写go语言中的测试程序

go语言中的测试程序规范

- 源码以_test.go结尾：xxx_test.go
- 测试方法名以Test开头：func TestXXX(t *testing.T){...}

```go
func TestFirstTry(t *testing.T) {
	t.Log("My first try")
}
```

## 2.go语言中声明变量的方式

一行代表一种方式

```go
	var i int
	var a int = 1 
	var b = 1
	var (
		c int = 1
		d     = 1
	)
	e := 1
	f, g := 1, 1
```

注意：和其他语言的不同之处：

- 赋值可以进行自动的类型判断
- 在赋值语句中可以对多个变量进行同时赋值（同python）

## 3.go语言中常量的设置

声明常量

```go
	const JAN = 1
	const (
		FEB = 2
		MAR = 3
		APR = 4
		MAY = 5
	)
```

快速声明连续的常量

```go
	const (
		JUNE = 6 + iota
		JULY
		AUG
		SEPT
		ORC
		NOV
		DEC
	)

	fmt.Println(JULY)  // 7
```
