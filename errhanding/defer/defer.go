package main

import (
	"fmt"
	"os"
	"bufio"
	"goLangLearn/functional/fib"
	"errors"
)

func tryDefer() {
	for i:=0; i<100; i++ {
		defer fmt.Println(i)

		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0;  i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func writeFile2(filename string) {
	file,err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)

	err = errors.New("this is new errors")
	if err != nil {
		if pathError,ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0;  i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//tryDefer()
	writeFile("fib.txt")
	writeFile2("fib.txt")
}
