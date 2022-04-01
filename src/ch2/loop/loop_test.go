package loop

import (
	"math/rand"
	"runtime"
	"testing"
)

func TestWhile(t *testing.T) {
	n := 1
	for n < 5 {
		// while循环
		n++
	}

	for {
		// 无限循环
		break
	}
}

func TestFor(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Log(i)
	}
	for i, v := range "abc" {
		t.Log(i, v)
	}
}

func TestIf(t *testing.T) {
	// go语言中的if语句支持变量赋值
	if n := 1; n > 0 {
		t.Log(n)
	}
}

func TestSwitchBreak(t *testing.T) {
	// go语言中的case自带break，go中的case可以支持字符串
	switch os := runtime.GOOS; os {
	case "darwin":
		t.Log("OS X")
	case "linux":
		t.Log("Linux.")
	default:
		t.Log("default")
	}
}

func TestSwitchRange(t *testing.T) {
	// go语言中的case可以判断范围
	Num := rand.Int()
	switch {
	case 0 <= Num && Num >= 3:
		t.Log("range 0-3")
	case 4 <= Num && Num >= 9:
		t.Log("range 4-9")
	case Num >= 10:
		t.Log("range 10 -> ")
	case Num < 0:
		t.Log("range <- 0")
	}
}
