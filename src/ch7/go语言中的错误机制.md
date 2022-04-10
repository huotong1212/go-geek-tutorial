## go语言中的错误机制

### 常见的Error使用和处理方式
```go
var AgeTooLarge = errors.New("年龄太大了")
var AgeTooLess = errors.New("年龄太小了")

func SetAge(age int) (int, error) {
	if age > 100 {
		return -1, AgeTooLarge
	}
	if age < 0 {
		return -1, AgeTooLess
	}

	return age, nil
}

func TestError(t *testing.T) {
	if v, err := SetAge(101); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}
```
太简单了，就不多赘述了。
### panic和recover

#### panic
- panic通常用于不可以恢复的错误
- panic在执行前会运行defer指定的内容

例子：
```go
func TestPanic(t *testing.T) {
	defer fmt.Println("defer run")
	fmt.Println("Start")
	panic("panic happened")
	fmt.Println("Ended")
}
```
output
```go
GOROOT=D:\Program Files\Go #gosetup
GOPATH=D:\Program Files\Go;C:\Users\82401\go #gosetup
"D:\Program Files\Go\bin\go.exe" test -c -o C:\Users\82401\AppData\Local\Temp\___TestPanic_in_error_test_go.exe D:/GolandProjects/go-geek-tutorial/src/ch7/error_test.go #gosetup
"D:\Program Files\Go\bin\go.exe" tool test2json -t C:\Users\82401\AppData\Local\Temp\___TestPanic_in_error_test_go.exe -test.v -test.paniconexit0 -test.run ^\QTestPanic\E$ #gosetup
=== RUN   TestPanic
Start
defer run
--- FAIL: TestPanic (0.00s)
panic: panic happened [recovered]
panic: panic happened

goroutine 6 [running]:
testing.tRunner.func1.2({0x252cc0, 0x29e100})
D:/Program Files/Go/src/testing/testing.go:1209 +0x24e
testing.tRunner.func1()
D:/Program Files/Go/src/testing/testing.go:1212 +0x218
panic({0x252cc0, 0x29e100})
D:/Program Files/Go/src/runtime/panic.go:1038 +0x215
command-line-arguments.TestPanic(0x0)
D:/GolandProjects/go-geek-tutorial/src/ch7/error_test.go:34 +0xb8
testing.tRunner(0xc000055ba0, 0x27d430)
D:/Program Files/Go/src/testing/testing.go:1259 +0x102
created by testing.(*T).Run
D:/Program Files/Go/src/testing/testing.go:1306 +0x35a


Process finished with exit code 1
```

#### Exit
- os.Exit也可以退出程序，但Exit退出前不会运行defer
- os.Exit退出不会打印栈的信息

例子：
```go
func TestExit(t *testing.T) {
	defer fmt.Println("defer run")
	fmt.Println("Start")
	os.Exit(-1)
	fmt.Println("Ended")
}
```
output
```go
=== RUN   TestExit
Start


Process finished with exit code 1
```

#### recover
recover和jave中的catch(Throwable t)和python中的 except Exception as e一样，可以捕获所有的panic,可以从pannic中恢复

例子：
```go
func TestRecover(t *testing.T) {
	defer func() {
		if err := recover();err!=nil{
			fmt.Println(err)
		}
	}()
	fmt.Println("Start !!!")
	panic(errors.New("something went wrong"))
}
```
output
```go
=== RUN   TestRecover
Start !!!
something went wrong
--- PASS: TestRecover (0.00s)
PASS
```
可以看到虽然产生了panic，但是程序最后是PASS的状态。

但是，我们并不推荐这样”错误“的使用recover恢复panic的方式：
- 这会形成很多僵尸进程，导致health check失败。
- ”Let it Crash“(失败后重启)往往是我们恢复不确定错误的最好方法。