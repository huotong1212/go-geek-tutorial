package ch93

import (
	"fmt"
	"testing"
	"time"
)

func service01() string {
	fmt.Println("service01 began ...")
	time.Sleep(time.Millisecond * 30)
	fmt.Println("service01 completed ...")
	return "service01 Done"
}

func service02() string {
	fmt.Println("service02 began ...")
	time.Sleep(time.Millisecond * 50)
	fmt.Println("service02 completed ...")
	return "service02 Done"
}

func TestSync(t *testing.T) {
	// 串行执行
	result01 := service01()
	result02 := service02()
	t.Log(result01, result02)
}

func AsyncService() chan string {
	ch := make(chan string)
	go func() {
		ret := service01()
		// 往chan中传入值
		ch <- ret
		fmt.Println("service01 exited")
	}()
	return ch
}

func TestAsync(t *testing.T) {
	// 可以看到它并不是按顺序执行的，而是异步的
	ch := AsyncService()
	ret02 := service02()
	// <-ch 会等待channel中的返回结果
	t.Log(<-ch, ret02)
	time.Sleep(1)
}

func AsyncServiceBuffer() chan string {
	ch := make(chan string, 1)
	go func() {
		ret := service01()
		// 往chan中传入值
		ch <- ret
		fmt.Println("service01 exited")
		ch <- ret
		fmt.Println("service01 exited twice")
	}()
	return ch
}

func TestAsyncBuffer(t *testing.T) {
	// 可以看到它并不是按顺序执行的，而是异步的
	ch := AsyncServiceBuffer()
	ret02 := service02()
	// <-ch 会等待channel中的返回结果
	t.Log(<-ch, ret02)
}
