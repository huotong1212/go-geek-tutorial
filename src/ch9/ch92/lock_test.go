package ch92

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	var counter int
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1)
	/**
	可以看到这里counter输出的不是5000，这是因为并发存在争抢的机制，
	可能在A协程中读取的counter为1，但A协程中counter还未进行counter++的操作，B协程就已经开始读取counter了，所以B协程读取的counter值也为1，
	所以虽然两个协程都执行了counter++的操作，但结果还是2，所以counter最终的结果一定是小于等于5000的
	*/
	fmt.Println(counter)
}

// 解决，加锁来保护共享内存（也可以使用读写锁只在写的时候进行保护，加快效率）
func TestLock(t *testing.T) {
	var mut sync.Mutex
	var counter int
	for i := 0; i < 5000; i++ {
		go func() {
			/**
			加入互斥锁，使该内存空间同时只能被一个协程访问，使得每一次的++都是有效的
			*/
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1)
	fmt.Println(counter) // 如果这里出现counter<5000的情况，可能原因是主进程运行结束了，但是协程还没有运行完
}

func TestWait(t *testing.T) {
	// 使用WaitGroup来替代time.Sleep()
	/**
	WaitGroup 只有当所有任务都完成的时候才会继续往下执行
	*/
	var wait sync.WaitGroup
	var mut sync.Mutex
	var counter int
	for i := 0; i < 500000; i++ {
		wait.Add(1)
		go func() {
			defer wait.Done()
			defer mut.Unlock()
			mut.Lock()
			//time.Sleep(time.Nanosecond*1)
			counter++
		}()
	}
	startTime := time.Now().Nanosecond()
	wait.Wait() // 只有当wait中Add的所有任务全部Done之后，Wait才会继续往下执行
	endTime := time.Now().Nanosecond()
	fmt.Println(endTime, startTime)
	fmt.Println("Wait Spent Time:", endTime-startTime)
	fmt.Println(counter)
}
