package ch101

import "testing"

type Cow struct {
	name string
}
type Dog struct {
	name string
}

func TestSameEmptyStruct(t *testing.T) {
	d1 := Dog{}
	d2 := Dog{}

	t.Logf("d1 pointer:%p, d2 pointer:%p \n", &d1, &d2)
	t.Log(&d1 == &d2)
}

func TestDifferentEmptyStruct(t *testing.T) {
	d1 := Dog{}
	d2 := Cow{}

	t.Logf("d1 pointer:%p, d2 pointer:%p \n", &d1, &d2)
	//t.Log(&d1 == &d2)  // 指针类型不一样不可以比较
}
