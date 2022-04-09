package extend

import (
	"fmt"
	"testing"
)

type Human struct {
	name string
}

func (h *Human) Eat() {
	fmt.Println(h.name, "Human is eating")
}

func (h *Human) Dinner() {
	fmt.Println(h.name, "'s dinner start")
	h.Eat()
}

type Student struct {
	Human
}

func (s *Student) Eat() {
	fmt.Println(s.name, "Student is eating")
}

func TestExtend(t *testing.T) {
	s1 := Student{Human{
		"小明",
	}}

	t.Log(s1)
	// 支持重写
	s1.Eat()
	// 不支持重载
	s1.Dinner()
}

func TestConvert(t *testing.T) {
	//var human Human = Human(Student{Human{"小明"}})  // 不支持这样声明，强转也不行
}

type Teacher struct {
	Human
	name string
}

func (s *Teacher) Eat() {
	fmt.Println(s.name, "Teacher is eating")
}

func TestDupField(t *testing.T) {
	teacher := Teacher{
		Human: Human{"厂长"},
		name:  "马老师",
	}

	teacher.Eat()
	teacher.Dinner()
}

type Father struct {
	name string
}

func (f *Father) Eat() {
	fmt.Println(f.name, "Teacher is eating")
}

type Manager struct {
	Human
	Father
	name string
}

func (s *Manager) Eat() {
	fmt.Println(s.name, "Teacher is eating")
}

func TestDupExtendMethod(t *testing.T) {
	manager := Manager{
		Human:  Human{"厂长"},
		Father: Father{"伍兹"},
		name:   "马老师",
	}

	manager.Eat() // 直接编译不通过  Ambiguous reference 'Eat'
	manager.Dinner()
}

type Pet struct {
	name string
}

type Animal struct {
	name string
}

type Cat struct {
	Pet
	Animal
}

func TestAccessEmbedField(t *testing.T) {
	cat := Cat{Pet{"tom"}, Animal{"timi"}}
	t.Log(cat.Pet.name) // Ambiguous reference 'name'
}
