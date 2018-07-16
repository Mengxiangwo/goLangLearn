package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

func variableZeroValue() {
	var a int
	var s string
	//fmt.Println(a, s)
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDedution() {
	var a, b, c, s = 3, 4, 5, "abc"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, 5, "abc"
	s = "abcd"
	fmt.Println(a, b, c, s)
}

func variablePork() {
	var (
		cc = "this"
		dd = "is"
		ee = "work"
	)

	fmt.Println(cc, dd, ee)
}

func euler() {
	/*c := 3 + 4i
	fmt.Println(cmplx.Abs(c))*/
	//fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1)
	//fmt.Println(cmplx.Exp(1i * math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i * math.Pi) + 1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))

	fmt.Println(c)
}

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	//常量定义之后可以不使用
	/*const(
		cpp = 0
		java = 1
		python = 2
		golang = 3
		sss = 4
	)*/
	const(
		cpp = iota
		java
		python
		golang
	)

	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	const(
		f1 = 2 * iota + 1
		f3
		f5
		f7
		f9
	)

	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
	fmt.Println(f1, f3, f5, f7, f9)
}

var (
	aa  = 3
	ss  = "kkk"
	bb = true)

const name = "name"
const (
	a1 = "this"
	a2 = "is"
	a3 = "const"
)

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDedution()
	variableShorter()
	fmt.Println(aa, bb, ss)
	variablePork()
	euler()
	triangle()
	consts()
	enums()
}