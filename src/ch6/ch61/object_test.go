package ch61

import (
	"fmt"
	"testing"
	"unsafe"
)

// 结构体封装数据
type Student struct {
	name   string
	age    int
	gender string
}

// 结构体的创建和初始化
func TestInitStruct(t *testing.T) {
	s1 := Student{
		name:   "小明",
		age:    6,
		gender: "男",
	}
	s2 := Student{"小张", 6, "男"}
	s3 := Student{}

	s4 := new(Student) // s4是一个指针类型
	s5 := &Student{}   // s5相当于s4

	t.Log(s1, s2, s3)
	t.Log(s4, s5)

	// 给指针类型创建的实例变量复制
	s4.name = "小樱"
	s4.age = 16
	s4.gender = "女"
	t.Log(s4)
}

func (s Student) Study() {
	fmt.Printf("%v instance is Studying\n", unsafe.Pointer(&s.name))
}

func (s *Student) Sport() {
	fmt.Printf("%v pointer is Sport\n", unsafe.Pointer(&s.name))
}

func TestAction(t *testing.T) {
	student := Student{"王小明", 17, "男"}
	fmt.Println("student 的内存地址是：", unsafe.Pointer(&student.name))

	student.Study()
	student.Sport()
}

func TestAccess(t *testing.T) {
	s1 := Student{"小花", 8, "男"}
	s2 := &Student{"小华", 9, "男"}

	// 使用实例成员访问两种方式定义的方法
	fmt.Println(unsafe.Pointer(&s1.name))
	s1.Sport()
	s1.Study()
	// 使用实例成员的指针访问两种方式定义的方法
	fmt.Println(unsafe.Pointer(&s2.name))
	s2.Sport()
	s2.Study()
}

func (s Student) ageGrow() {
	s.age += 1
}

func (s *Student) nameChange() {
	s.name = "Shirley"
}

func TestParamPass(t *testing.T) {
	s1 := Student{"小花", 8, "女"}
	// 对于值传递，任何在 method 内部对 value 做出的改变都不影响调用者看到的 value
	s1.ageGrow()
	s1.nameChange()
	fmt.Println(s1) // {Shirley 8 女}
}
