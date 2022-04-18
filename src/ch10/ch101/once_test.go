package ch101

import (
	"fmt"
	"sync"
	"testing"
)

var one sync.Once

func Do() {
	var i = 10
	fmt.Println("Before Do：", i)
	one.Do(func() {
		i += 5
	})
	fmt.Println("After Do：", i)
}

func TestOnce(t *testing.T) {
	for i := 0; i < 3; i++ {
		Do()
	}
}
