package ch62

import (
	"fmt"
	"testing"
)

type doNothing interface {
}

func do(i doNothing) {
	nothing, ok := i.(doNothing)
	fmt.Printf("%v,type t: %T ,is doNothing:%v %v \n", i, i, nothing, ok)
}

type School struct {
	name string
}

func TestEmptyInterface(t *testing.T) {
	do(1)
	do("abc")
	do(School{"苏州科技大学"})
	do(&School{"苏州大学"})
}

func doSliceInterface(s []interface{}) {
	fmt.Println(s)
}

func TestSliceInterface(t *testing.T) {
	s1 := []int{1, 2, 3}
	t.Log(s1)
	//doSliceInterface(s1)  // 编译错误 cannot use s1 (type []int) as type []interface {} in argument to doSliceInterface

	var i []interface{}
	for _, v := range s1 {
		i = append(i, v)
	}

	doSliceInterface(i)
}
func doVariable(params ...interface{}) {
	fmt.Println(params...)

	for _, v := range params {
		fmt.Println(v)
	}
}

func TestVariableParams01(t *testing.T) {
	doVariable(1, "abc", 1.1)
}

func TestVariableParams02(t *testing.T) {
	doVariable([]int{5, 6, 7}, "abc", 1.1)
}

func variableInt(is ...int) {
	//fmt.Println(is...)  // Cannot use 'is' (type []int) as type []interface{}
}

func TestAssert(t *testing.T) {
	var s []interface{} = []interface{}{1, 2, 3}
	//i,flag := s.([]interface{}) // (non-interface type []interface{} on left)
	v1, ok1 := interface{}(s).(interface{})
	v2, ok2 := interface{}(s).([]interface{})
	t.Log(v1, ok1)
	t.Log(v2, ok2)

	str := "小明爱吃鱼"
	inv, ok3 := interface{}(str).(interface{})
	strv, ok4 := interface{}(str).(string)
	sliceinv, ok5 := interface{}(str).([]interface{})
	t.Log(inv, ok3)
	t.Log(strv, ok4)
	t.Log(sliceinv, ok5)

	// value, ok := em.(T) em必须是接口类型的变量，如果不是需要我们显示强转
}
