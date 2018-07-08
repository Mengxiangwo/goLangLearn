package main

import (
	"io/ioutil"
	"fmt"
)

func eval(a , b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
		fallthrough //不break
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op)
	}

	return result
}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score:%d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}

	return g
}

func main() {
	const filename = "abc.txt"
	/*contents, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}*/

	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(eval(1, 4, "+"))

	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(85),
		grade(95),
		grade(-3),
		)
}
