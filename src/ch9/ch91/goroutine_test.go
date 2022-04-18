package ch91

import (
	"fmt"
	"testing"
	"time"
)

func TestGoRoutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i, &i)
		}(i)
	}
	// 为了防止主进程比协程先执行完
	time.Sleep(1)
	// 可以看到打印出来的结果是乱序执行的，就表示是并发执行的，不是串行的
}

func TestGoRoutineSharedMemory(t *testing.T) {
	// 测试goroutine争强同一个内存空间的状态
	for i := 0; i < 10; i++ {
		fmt.Println("main:", i, &i)
		go func() {
			fmt.Println(i, &i)
		}()
	}
	time.Sleep(1)
	// 可以看到打印的先是main，然后才是10,10,10
	// 这是因为这样赋值的i的内存是共享的,共享的话就存在一个争抢的情况，其他协程得等主进程访问结束了才能使用这个变量的内存空间
}

func TestGoRoutineSharedMemory2(t *testing.T) {
	// 测试goroutine争强同一个内存空间的状态
	for i := 0; i < 10; i++ {
		fmt.Println("main:", i, &i)
		go func() {
			i++ // 可以看到共享变量内存时是串行递增的，证明被加了锁
			fmt.Println(i, &i)
		}()
	}
	time.Sleep(1)
	// 可以看到打印的先是main，然后才是10,10,10
	// 这是因为这样赋值的i的内存是共享的,共享的话就存在一个争抢的情况，其他协程得等主进程访问结束了才能使用这个变量的内存空间
}
