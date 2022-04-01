package demo_test

import (
	"fmt"
	"testing"
)

func TestFirstTry(t *testing.T) {
	t.Log("My first try")
}

// 实现一个斐波拉契数列  1，1，2，3，5，8，13
func TestFib(t *testing.T) {
	a, b := 1, 1
	for i := 0; i < 5; i++ {
		c := a + b
		t.Log(c)
		b = a
		a = c
	}
}

// go语言中变量赋值的方式
func TestVar(t *testing.T) {
	// var
	var i int
	var a int = 1
	var b = 1
	var (
		c int = 1
		d     = 1
	)
	e := 1
	f, g := 1, 1

	fmt.Println(a, b, c, d, e, f, g, i)
}

func TestConstant(t *testing.T) {
	const JAN = 1
	const (
		FEB = 2
		MAR = 3
		APR = 4
		MAY = 5
	)

	const (
		JUNE = 6 + iota
		JULY
		AUG
		SEPT
		ORC
		NOV
		DEC
	)

	fmt.Println(JULY)

}
