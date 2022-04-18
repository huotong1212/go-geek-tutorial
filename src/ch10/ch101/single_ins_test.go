package ch101

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

// 单例模式-懒汉式

type EmptyStruct struct {
	name string
}

var once sync.Once
var ins *EmptyStruct

func (e *EmptyStruct) getInstance() *EmptyStruct {
	once.Do(func() {
		fmt.Println("Create obj")
		ins = new(EmptyStruct)
	})
	return ins
}

var singleInstance *EmptyStruct

func getSingleIns() *EmptyStruct {
	//var singleInstance *EmptyStruct
	//once.Do(func() {
	//	fmt.Println("create a single instance")
	//	singleInstance = new(EmptyStruct)
	//})
	singleInstance = new(EmptyStruct)
	return singleInstance
}

//var emptyIns EmptyStruct
//func getEmptyIns() EmptyStruct {
//	once.Do(func() {
//		fmt.Println("create a single instance")
//		emptyIns = EmptyStruct{}
//	})
//	return emptyIns
//}

func TestSingleTon01(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		// 创建10个协程，创建结构体，看看返回的结构体是否相同
		go func() {
			obj := EmptyStruct{}
			ins := obj.getInstance()
			fmt.Println("obj:", obj, "ins:", ins)
			fmt.Println("&obj:", &obj, "&ins:", &ins)
			fmt.Println("obj upointer:", unsafe.Pointer(&obj), "ins upointer:", unsafe.Pointer(&ins))
			//fmt.Printf("obj d: %d ,ins d %d:", obj, ins)
			//fmt.Printf("obj x: %x ,ins x %x \n", obj, ins)
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestSingleTon02(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		// 创建10个协程，创建结构体，看看返回的结构体是否相同
		go func() {
			ins := getSingleIns()
			//fmt.Println("ins:", ins)
			//fmt.Println("ins upointer:", unsafe.Pointer(ins))
			fmt.Printf("ins p %p:\n", ins)
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestSingleTon03(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		// 创建10个协程，创建结构体，看看返回的结构体是否相同
		go func() {
			instance := ins.getInstance()
			fmt.Printf("ins p %p:\n", instance)
			wg.Done()
		}()
	}
	wg.Wait()
}
