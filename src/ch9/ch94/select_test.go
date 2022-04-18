package ch94

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

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

func TestSelectCase(t *testing.T) {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
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

func spendTime() chan string {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
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
