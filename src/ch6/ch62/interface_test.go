package ch62

import (
	"fmt"
	"testing"
)

func Dinner(human Human) {
	human.Eat()
}

func TestInterface(t *testing.T) {
	Dinner(&Student{name: "小明"})
	Dinner(&Teacher{name: "马老师"})
}

func TestInterfaceAssert(t *testing.T) {
	var student Human = &Student{name: "小明同学"}
	var teacher Human = &Teacher{name: "马老师"}

	value, ok := student.(*Student) // 判断student接口变量中存储的是不是Student的指针类型
	t.Log(value, ok)                // &{小明同学 0 } true

	value, ok = teacher.(*Student)
	t.Log(value, ok) // <nil> false
}

func TestSwitchInterfaceAssert(t *testing.T) {
	var student Human = &Student{name: "小明同学"}

	switch t := student.(type) {
	case *Teacher:
		fmt.Println(t, "he is a teacher")
	case *Student:
		fmt.Println(t, "he is a student")
	}
}

func (s *Student) Homework() {
	fmt.Println(s.name, "is doing homework")
}

func TestInterfaceAccessNotExistMethod(t *testing.T) {
	var human Human = &Student{name: "小明"}
	human.Eat()
	// human.Homework()  // 编译错误 Unresolved reference 'Homework'
}

type Teacher struct {
	name   string
	age    int
	gender string
}

func (s *Teacher) Eat() {
	fmt.Println(s.name, "is eating")
}

type Student struct {
	name   string
	age    int
	gender string
}

type Human interface {
	Eat()
}

func (s *Student) Eat() {
	fmt.Println(s.name, "pointer receiver is eating")
}

type Senior interface {
	MathLearn()
}

func (s Student) MathLearn() {
	fmt.Println(s.name, "instance receiver is learning math")
}

func TestInterfaceImpl(t *testing.T) {
	// 如果实现接口时使用的reciver是指针类型，则传入实例会编译不通过
	//var pupil Human = Student{"小哈",6,"男"}
	// 如果实现接口时使用的reciver是指针类型，则传入指针会编译通过
	var human Human = &Student{"小哈 pointer", 6, "男"}

	// 如果实现接口时使用的reciver是实例类型，则传入实例会编译通过
	var seniorInstance Senior = Student{"小哈 Instance", 6, "男"}
	// 如果实现接口时使用的reciver是实例类型，则传入指针会编译通过
	var seniorPointer Senior = &Student{"小哈 pointer", 6, "男"}

	fmt.Println(human, seniorInstance, seniorPointer)

	human.Eat()
	seniorInstance.MathLearn()
	seniorPointer.MathLearn()
}

type College interface {
	Work()
	Social()
}

func (s *Student) Work() {
	fmt.Println(s.name, "pointer is working")
}

//func (s *Student) Social() {
//	fmt.Println(s.name,"pointer is working")
//}

func TestImplInterfaceAllMethod(t *testing.T) {
	// 一个类型必须实现这个接口下的所有方法，才算作是实现了该接口，否则编译会不通过
	//var student College = &Student{name: "小明"}
}
