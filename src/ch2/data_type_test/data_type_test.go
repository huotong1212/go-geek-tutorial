package data_type_test

import (
	"math"
	"testing"
)

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64

	//b = a  //cannot use a (type int32) as type int64 in assignment
	b = int64(a) // ok
	t.Log(a, b)
}

func TestMath(t *testing.T) {
	t.Log(math.MaxInt)
	t.Log(math.MaxFloat64)
	t.Log(math.MaxFloat32)
	// ...
}

func TestPoint(t *testing.T) {
	a := 1
	aPoint := &a

	// aPoint += 1  // 不支持指针运算
	t.Log(a, aPoint)
	t.Logf("%T,%T", a, aPoint)
}
