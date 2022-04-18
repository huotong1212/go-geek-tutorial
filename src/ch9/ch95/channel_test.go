package ch95

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func dataProducer(ch chan string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("produce data:" + strconv.Itoa(i))
			ch <- strconv.Itoa(i)
		}(i)
	}
	wg.Done()
}

func dataReceiver(ch chan string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		go func() {
			data := <-ch
			fmt.Println("receiver data:", data)
		}()
	}
	wg.Done()
}

func TestChannel(t *testing.T) {
	ch := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}

func dataProducerSignal(ch chan string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			ch <- strconv.Itoa(i)
			fmt.Println("produce data:" + strconv.Itoa(i))
			wg.Done()
		}(i)
	}
	// 这里我们关闭channel
	wg.Done()
	wg.Wait()
	close(ch)
}

func dataReceiverSignal(ch chan string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			data, ok := <-ch
			if ok {
				fmt.Println("receiver data:", data)
			} else {
				fmt.Println("channel has been closed")
			}
			wg.Done()
		}()
	}
	wg.Done()
}

func TestChannelSignal(t *testing.T) {
	ch := make(chan string, 20)
	wg := sync.WaitGroup{}
	wg.Add(1)
	dataProducerSignal(ch, &wg)
	wg.Add(1)
	dataReceiverSignal(ch, &wg)
	wg.Wait()
}

func dataProducerOfficial(ch chan string, wg *sync.WaitGroup) {
	// 一个生产者
	go func() {
		for i := 0; i < 10; i++ {
			ch <- strconv.Itoa(i)
			fmt.Println("produce data:" + strconv.Itoa(i))
		}
		// 这里我们关闭channel
		close(ch)
		wg.Done()
	}()
	// 执行关闭channel这里要放在生产者之中，否则会报 send on closed channel 向一个已关闭的通道发送消息
	//close(ch)
}

func dataReceiverOfficial(ch chan string, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Println("receiver data:", data)
			} else {
				fmt.Println("channel has been closed")
				break
			}
		}
		wg.Done()
	}()
}

func TestChannelOfficial(t *testing.T) {
	ch := make(chan string)
	wg := sync.WaitGroup{}
	// 一个生产者
	wg.Add(1)
	dataProducerOfficial(ch, &wg)
	// 两个消费者
	wg.Add(1)
	dataReceiverOfficial(ch, &wg)
	wg.Add(1)
	dataReceiverOfficial(ch, &wg)
	wg.Wait()
}
