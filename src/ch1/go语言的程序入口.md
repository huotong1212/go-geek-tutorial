## 编写第一个go语言的HelloWorld程序

```go
package main

import "fmt"

func main()  {
	fmt.Println("Hello World")
}
```

注意点：

- go语言中的应用程序的入口

  - 必须是main包下
  - 必须是main方法
  - 但是文件名不一定是main.go
- go语言的main方法不能有返回值,但是可以通过os.Exit来设置返回状态

  ```go
  func main() string {
          fmt.Println("Hello World")
  	return "ok"
  }

  // 报如下错误
  .\HelloWorld.go:5:6: func main must have no arguments and no return values
  ```
  ```go
  func main()  {
  	fmt.Println("Hello World")
  	os.Exit(100)
  }

  // 执行返回
  Process finished with the exit code 100
  ```
- go语言中的main方法不支持传入参数，但是可以通过os.Args来获取命令行参数

  ```go

  func main()  {
  	fmt.Println("Hello World")
  	fmt.Println(os.Args)
  	os.Exit(777)
  }

  // 输出
  D:\GoPros\go-geek-tutorial\src\ch1>go run  HelloWorld.go apple xiaomi
  Hello World
  [C:\Users\huotong\AppData\Local\Temp\go-build1933368431\b001\exe\HelloWorld.exe apple xiaomi]
  exit status 777
  ```
