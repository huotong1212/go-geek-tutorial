## Channel

### Channel机制的简单介绍
#### 典型Channel
客户端和服务端都必须同时在线才能完成一次channel的收发工作。
- 如果服务端发送了消息，客户端没有完成接收（可能是掉线或其他原因），服务端会一直处于阻塞状态。
- 如果客户端在接受消息，而服务端还没有发送消息（掉线或者其他原因），客户端会一直处于阻塞状态。

![channel](https://tvax3.sinaimg.cn/large/007uGCBvly1h191mdddhaj30he0fgq5i.jpg)

#### BufferedChannel
只要channel缓存中有消息存在就能完成一次channel的收发工作
- 服务端向channelBuffer（在Buffer还有容量的情况下）中发送了一条消息，可以不用等待客户端是否返回就可以继续往下执行
- 客户端从channelBuffer中接受消息，只要Buffer中存在消息，客户端就可以不用管服务端是否在线而继续往下执行

![bufferedchannel](https://tva2.sinaimg.cn/large/007uGCBvly1h191nnql84j30jq09qgoh.jpg)

### 代码实操
这是一个串行的程序,它的执行时间要0.1s
```go
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
```
out
```go
=== RUN   TestSync
service01 began ...
service01 completed ...
service02 began ...
service02 completed ...
    channel_test.go:27: service01 Done service02 Done
--- PASS: TestSync (0.10s)
PASS
```
这是一个通过channel管道实现的future模式，它的实现只要0.07s
```go
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
	time.Sleep(1) // 这里等待1s，是为了让服务端接收完消息不再阻塞，从而可以输出"service01 exited" 
}
```
out
```go
=== RUN   TestAsync
service02 began ...
service01 began ...
service01 completed ...
service02 completed ...
service01 exited
    channel_test.go:46: service01 Done service02 Done
--- PASS: TestAsync (0.07s)
PASS
```
注意：
1.这里向channel发送ret字符串消息，但是根据以上机制，channel的服务端会在客户端还没有接收消息时阻塞住，所以`service01 exited`并不会马上打印出来。
```go
		ch <- ret
		fmt.Println("service01 exited")
```
2.在客户端接受消息后，服务端才会继续执行，所以`service01 exited`在最后显示
```go
t.Log(<-ch, ret02)
```

一点点改进，
```go
ch := make(chan string, 1)
```
out
```go
=== RUN   TestAsyncBuffer
service02 began ...
service01 began ...
service01 completed ...
service01 exited
service02 completed ...
    channel_test.go:66: service01 Done service02 Done
--- PASS: TestAsyncBuffer (0.06s)
PASS
```
因为使用了 BufferedChannel ，所以服务端不会阻塞，而是会马上继续执行。

思考?如果bufferedChannel管道中的容量小于发送的消息数，服务端是否会继续阻塞
```go
		// 往chan中传入值
		ch <- ret
		fmt.Println("service01 exited")
		ch <- ret
		fmt.Println("service01 exited twice")
```
out
```go
=== RUN   TestAsyncBuffer
service02 began ...
service01 began ...
service01 completed ...
service01 exited
service02 completed ...
service01 exited twice
    channel_test.go:68: service01 Done service02 Done
--- PASS: TestAsyncBuffer (0.05s)
PASS
```
是，会继续阻塞，可以看到`service01 exited twice`并没有被打印出来。

