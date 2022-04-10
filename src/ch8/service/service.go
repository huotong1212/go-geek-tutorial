package service

import "fmt"

func init() {
	fmt.Println("service init 01")
}

func init() {
	fmt.Println("service init 02")
}

func Add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}
