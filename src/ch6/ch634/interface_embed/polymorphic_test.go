package interface_embed

import (
	"fmt"
	"testing"
)

// go语言中通过接口实现多态测试
type Pet interface {
	Eat()
}

type Dog struct{}

func (d *Dog) Eat() {
	fmt.Println("dog is eating")
}

type Cat struct{}

func (d *Cat) Eat() {
	fmt.Println("cat is eating")
}

func Dinner(pet Pet) {
	pet.Eat()
}

func TestPolymorphic(t *testing.T) {
	var dog Pet = new(Dog)
	var cat Pet = new(Cat)

	Dinner(dog) // dog is eating
	Dinner(cat) // cat is eating
}
