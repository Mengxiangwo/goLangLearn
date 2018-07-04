package main

import "fmt"

var (
	b1 = "cc"
	b2 = "dd"
	b3 = "ee"
)

func variableFone() {
	a1, a2,  a3 := 5, 6, "hello"

	fmt.Println(a1, a2, a3)
}

func variableClosure() {
	var (
		aa = "this"
		bb = "is"
		cc = "work"
	)

	println(aa, bb, cc)
}

func main() {
	fmt.Println("Hello world")
	variableFone()
	fmt.Println(b1, b2, b3)
	variableClosure()
}