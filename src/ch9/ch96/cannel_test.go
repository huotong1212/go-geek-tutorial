package ch96

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func isCancelled(ch chan int) bool {
	select {
	case v, ok := <-ch:
		fmt.Printf("receive value:%v %v from cannelled channel\n", v, ok)
		return true
	default:
		return false
	}
}

func cancelChannel01(ch chan int) {
	// 这样不行，因为只取消了第一个接受消息的协程，后面的无法取消
	ch <- 1
}

func cancelChannel02(ch chan int) {
	// 利用关闭channel自带的广播机制来实现取消
	close(ch)
}

func TestCancelChannel(t *testing.T) {
	var ch = make(chan int)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int, ch chan int) {
			for !isCancelled(ch) {
				fmt.Println("wait cancel")
				time.Sleep(time.Millisecond * 5)
			}
			wg.Done()
			fmt.Println("channel cancelled")
		}(i, ch)
	}
	cancelChannel02(ch)
	wg.Wait()
}
