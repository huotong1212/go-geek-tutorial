package main

import (
	"ch8/service"
	"fmt"
)

func init() {
	fmt.Println("client init 01")
}

func init() {
	fmt.Println("client init 02")
}

func main() {
	//sub := service.subtract(1,2) // Unexported function 'subtract' usage
	sum := service.Add(1, 2)
	fmt.Println("Sum:", sum)
}
