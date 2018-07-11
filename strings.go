package main

import "fmt"

func main() {
	s := "PHP是最好的语言"
	r := []rune(s)

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}

	fmt.Println(len(r))

	fmt.Println()
}
