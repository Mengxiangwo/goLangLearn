package main

import (
	"fmt"
	"time"
	"sync"
)

type atomicInt struct {
	value int
	lock sync.Mutex
}

func (a *atomicInt) increment() {
	/*
	//1. mutex lock 作用于當前作用域，也就是當前函數下
	a.lock.Lock()
	defer a.lock.Unlock()

	a.value++
	*/

	//2. 可以使用匿名函數控制作用域的大小
	fmt.Println("safe increment")
	func(){
		a.lock.Lock()
		defer a.lock.Unlock()

		a.value++
	}()
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()

	return int(a.value)
}

func main() {
	var a  atomicInt
	a.increment()

	go func() {
		a.increment()
	}()

	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
