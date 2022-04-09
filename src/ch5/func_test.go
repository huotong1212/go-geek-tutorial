package ch5

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	// 这里的随机是一致随机性
	return rand.Intn(10), rand.Intn(20)
}

func TestFunc(t *testing.T) {
	n1, n2 := returnMultiValues()
	t.Log(n1, n2)
}

// 使用函数式编程计算时常（类似python中的装饰器）
func Timer(fn func(i int) int) func(i int) int {
	return func(i int) int {
		startTime := time.Now().Second()
		n := fn(i)
		endTime := time.Now().Second()
		fmt.Printf("Spend second:%v s ", endTime-startTime)
		return n
	}
}

func SleepMul(i int) int {
	time.Sleep(time.Second * 10)
	return i * i
}

func TestTimer(t *testing.T) {
	TimeSleep := Timer(SleepMul)
	TimeSleep(9)
}

func MultiParams(ops ...int) int {
	sum := 0
	for _, v := range ops {
		sum += v
	}

	return sum
}

func TestMultiParams(t *testing.T) {
	v1 := MultiParams(1, 2, 3)
	v2 := MultiParams(1, 2, 3, 4, 5)

	t.Log(v1, v2)
}

func TestDefer01(t *testing.T) {
	defer func() {
		fmt.Println("Defer run")
	}()
	t.Log("Started\n")
}

func Defer02() (int, error) {
	defer func() {
		fmt.Println("Defer run")
	}()
	fmt.Println("Started")
	return fmt.Println("Returned")
}

func TestDefer02(t *testing.T) {
	_, _ = Defer02()
}

func Defer03() (int, error) {
	defer func() {
		fmt.Println("Defer run")
	}()
	fmt.Println("Started")
	panic("this is a panic")
	return fmt.Println("Returned")
}

func TestDefer03(t *testing.T) {
	_, _ = Defer03()
}

func TestDefer04(t *testing.T) {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	t.Log("Started\n")
}
