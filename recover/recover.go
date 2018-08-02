package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else if str, ok := r.(string); ok {
			fmt.Println("String occurred:", str)
		} else {
			panic(r)
		}
	}()

	//panic(errors.New("this is error"))
	panic("this is string error")
}

func main() {
	tryRecover()
}
