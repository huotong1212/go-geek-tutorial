package ch97

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		fmt.Println("context has done")
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
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)
	for i := 0; i < 10; i++ {
		go func(i int, ctx context.Context) {
			for !isCancelled(ctx) {
				fmt.Println("wait cancel")
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println("channel cancelled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}
