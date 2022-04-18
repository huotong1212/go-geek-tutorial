package ch101

import (
	"fmt"
	"testing"
	"unsafe"
)

type Cat struct {
	name string
}

func TestPointer(t *testing.T) {
	// &测试
	cat := Cat{}
	t.Log(cat)  // 获取变量
	t.Log(&cat) // 获取变量的指针

	catPointer := &cat              // 定义了一个变量，类型为指针，指向cat的指针
	t.Logf("Type : %T", catPointer) // 打印类型
	t.Log(&catPointer)              // 获取指向这个指针变量的指针
}

func TestUnsafePointer(t *testing.T) {
	// unsafe.pointer 测试
	cat := Cat{}
	// unsafe.Pointer只接收指针变量类型，但是它可以把任意类型的指针变量类型都变成unsafe.Pointer类型
	up := unsafe.Pointer(&cat)
	t.Log(up)               // 显示可寻址的指针值
	t.Logf("Type : %T", up) // 打印类型
}

func TestFmtP(t *testing.T) {
	cat := Cat{}
	// 这个 %p 也只接收指针类型
	fmt.Printf("format point: %p \n", &cat)
}

func TestCompare(t *testing.T) {
	// 三者对比
	cat := Cat{}
	catPointer := &cat

	fmt.Println("&:", &catPointer)                             // 获取这个指向->cat指针->的变量的指针
	fmt.Println("unsafe pointer:", unsafe.Pointer(catPointer)) // 将*Cat变成unsafe.Pointer,然后显示出可寻址的指针值（应该就算原来指针指向的内存地址）
	fmt.Printf("format pointer: %p\n", catPointer)             // 直接打印出这个指针的指向的内存地址

	// 综上，当我们要获取一个指针指向的内存地址时，推荐使用 unsafe.Pointer或者 %p的形式
}
