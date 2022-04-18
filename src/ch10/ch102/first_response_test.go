package ch102

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func dataProducer(ch chan int) {
	go func() {
		for i := 0; i < 10; i++ {
			//fmt.Println("produce start:", i)
			ch <- i
			fmt.Println("produce end:", i)
		}
	}()
}

func dataConsumer(ch chan int) {
	go func() {
		for i := 0; i < 10; i++ {
			//fmt.Println("consume start:", i)
			v := <-ch
			fmt.Println("consume get:", v)
			//fmt.Println("consume end:", i)
		}
	}()
}

func TestChannel(t *testing.T) {
	// 可以看到channel中的数据是一个一个按顺序打印的，证明非缓冲channel是生产者生产一个数据，消费者消费一个数据，如果消费者不消费数据，生产者就会hang住，如果生产者不生产数据，消费者就会hang住
	ch := make(chan int)
	dataProducer(ch)
	dataConsumer(ch)
	time.Sleep(time.Second * 2)
}

func runTask(ch chan int, i int) {
	fmt.Println("run Task:", i)
	ch <- i
	fmt.Println("task done:", i)
}

func FirstResponse() int {
	var numOfRoutines int = 10
	//ch := make(chan int)
	ch := make(chan int, numOfRoutines) // 优化
	for i := 0; i < numOfRoutines; i++ {
		go func(ch chan int, i int) {
			runTask(ch, i)
		}(ch, i)
	}

	return <-ch // 这里接收消息应该是一个争抢的过程
}

func TestRoutine(t *testing.T) {
	t.Log(FirstResponse())
	time.Sleep(time.Second * 2)
}

func TestRoutine02(t *testing.T) {
	// 这样会有一个问题，就是虽然FirstResponse返回了，但是因为没有消费者来接收channel中的消息，其他协程会在 ch <- i 这里hang住，等待消费者消费
	// 解决办法：1.将channel改成buffered channel，这样就不会在  ch <- i这里hang住了，而是直接结束协程。
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second * 2)
	t.Log("After:", runtime.NumGoroutine())
}

func runTask02(ch chan int, i int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("routine shut down")
		}
	}()
	fmt.Println("run Task:", i)
	ch <- i
	fmt.Println("task done:", i)
}

func FirstResponse02(ch chan int) int {
	// 当firstResponse从channel中返回以后，关闭channel
	var numOfRoutines int = 10
	//ch := make(chan int, numOfRoutines) // 优化
	for i := 1; i < numOfRoutines; i++ {
		go func(ch chan int, i int) {
			runTask02(ch, i)
		}(ch, i)
	}

	return <-ch // 这里接收消息应该是一个争抢的过程
}

func TestRoutine03(t *testing.T) {
	// 这样会有一个问题，就是虽然FirstResponse返回了，但是因为没有消费者来接收channel中的消息，其他协程会在 ch <- i 这里hang住，等待消费者消费
	// 解决办法：尝试通过关闭ch释放其他协程
	t.Log("Before:", runtime.NumGoroutine())
	ch := make(chan int)
	t.Log("first response:", FirstResponse02(ch))
	close(ch)
	time.Sleep(time.Second * 1)
	//t.Log("After")
	t.Log("After:", runtime.NumGoroutine())
}
