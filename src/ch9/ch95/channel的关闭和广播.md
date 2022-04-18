## channel的关闭

### 生产者和消费者模型
go中协程之间的消息传递一般使用channel
```go
func dataProducer(ch chan string,wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("produce data:" + strconv.Itoa(i))
			ch <- strconv.Itoa(i)
		}(i)
	}
	wg.Done()
}

func dataReceiver(ch chan string,wg *sync.WaitGroup) {
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
	// 十个生产者
	dataProducer(ch,&wg)
	wg.Add(1)
	// 十个消费者
	dataReceiver(ch,&wg)
	wg.Wait()
}
```
存在的问题：
receiver并不知道producer何时停止放数据，特别是有多个receiver的情况下，
这时候我们就需要用到通道的关闭的广播机制

### 通道的关闭
channel的关闭
- 向关闭的channel发送数据，会导致pannic
- v,ok <- ch;ok为bool值，true表示正常接受,false表示通道关闭
- 所有channel的接收者都会在channel关闭时，立刻从阻塞等待中返回，且上述ok值为false.
这个广播机制常被利用，进行向多个订阅者同时发送信号。如：退出信号。
  

例子:
```go
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
```
注意：
1.`close(ch)`这个关闭的操作最好在生产者的协程中进行。