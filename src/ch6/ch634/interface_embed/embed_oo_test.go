package interface_embed

import (
	"fmt"
	"testing"
)

// Go 支持任意类型的 embed type，当然也包括 interface type，通过组合就可以实现多种不同行为的任意组合，这也是 Go 倡导以更小的单元实现你的代码功能，然后组合它们的理念。

// 首先定义两个行为
type StudentTalk interface {
	talk()
}

type TeacherTalk interface {
	say()
}

// 组合这两个行为
type PersonTalk interface {
	StudentTalk
	TeacherTalk
}

// 通过 embed type 把定义的两个 interface 组合为新的 interface PeopleTalk，此时 PeopleTalk 继承了两个 interface 的 method 集合，也就是 PeopleTalk 拥有了 StudentTalk 和 TeacherTalk 的 method 合集。

// 继承这两个行为，Person相当于实现了PersonTalk接口
type Person struct {
	StudentTalk
	TeacherTalk
}

// Person 也内嵌了 TeacherTalk 和 StudentTalk，对 Person 来说既可以理解成继承了两个 interface 的 method 集合，也可以理解是 Person 拥有两个类型为 TeacherTalk 和 StudentTalk 的 field，它们分别可以被赋值为实现了它们的 struct 的值。

// 实现这两个行为
type Student struct{}

func (s *Student) talk() {
	fmt.Println("Student is talking")
}

type Teacher struct{}

func (t *Teacher) say() {
	fmt.Println("Teacher is saying")
}

// 展示魔力
func meet(personTalk PersonTalk) {
	fmt.Println("====>people meet<====")
	meetStudent(personTalk)
	meetTeacher(personTalk)
}

func meetStudent(s StudentTalk) {
	fmt.Println("====>student meet<====")
	s.talk()
}

func meetTeacher(s TeacherTalk) {
	fmt.Println("====>teacher meet<====")
	s.say()
}

func TestMagic(t *testing.T) {
	var student StudentTalk = &Student{}
	var teacher TeacherTalk = &Teacher{}

	person := Person{student, teacher}
	meet(person)
}
