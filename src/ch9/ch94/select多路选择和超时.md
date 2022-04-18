## select多路选择和超时

### 多路复用

错误示例：

```go
func TestSelectCaseErrorExample(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	var wait sync.WaitGroup
	wait.Add(2)
	go func() {
		defer wait.Done()
		ch1 <- "func01"
		fmt.Println("func01 created")
	}()

	go func() {
		defer wait.Done()
		ch2 <- "func02"
		fmt.Println("func02 created")
	}()


	wait.Wait()

	select {
	case <-ch1:
		fmt.Println("received func01 results")
	case <-ch2:
		fmt.Println("received func02 results")
	default:
		fmt.Println("selected default choice")
	}
}

```

output

```go
fatal error: all goroutines are asleep - deadlock!
```

注意：这会造成死锁，

原因：

```go
ch1 <- "func01"
```

这个管道发送完消息之后，如果客户端没有接收消息会一直hang住

```go
defer wait.Done()
```

wait任务的释放会在协程的最后进行

```go
wait.Wait()
```

如果没有Done，这边就会一直Wait

综上，这三个地方形成了死锁。


解决办法：

1.将wait.Done()放在ch<-之前执行

2.或者将channel改完Buffered Channel,如下

```go
func TestSelectCase(t *testing.T) {
	ch1 := make(chan string,1)
	ch2 := make(chan string,1)
	var wait sync.WaitGroup
	wait.Add(2)
	go func() {
		defer wait.Done()
		ch1 <- "func01"
		fmt.Println("func01 sent")
	}()

	go func() {
		defer wait.Done()
		ch2 <- "func02"
		fmt.Println("func02 sent")
	}()


	wait.Wait()

	select {
	case <-ch1:
		fmt.Println("received func01 results")
	case <-ch2:
		fmt.Println("received func02 results")
	default:
		fmt.Println("selected default choice")
	}

}
```

### select 超时示例

```go
func spendTime() chan string {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Second*2)
		fmt.Println("spendTime ended")
		ch <- "OK"
	}()
	return ch
}

func TestSelectTimeOut(t *testing.T) {
	var ch chan string
	ch = spendTime()

	select {
	case <-ch:
		fmt.Println("Done")
	case <-time.After(time.Second * 1):
		t.Error("time out")
	}
}
```
