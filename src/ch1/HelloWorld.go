package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(os.Args)
	os.Exit(777)
}
