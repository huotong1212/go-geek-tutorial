package ch7

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

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

func TestPanic(t *testing.T) {
	defer fmt.Println("defer run")
	fmt.Println("Start")
	panic("panic happened")
	fmt.Println("Ended")
}

func TestExit(t *testing.T) {
	defer fmt.Println("defer run")
	fmt.Println("Start")
	os.Exit(-1)
	fmt.Println("Ended")
}

func TestRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Start !!!")
	panic(errors.New("something went wrong"))
}
