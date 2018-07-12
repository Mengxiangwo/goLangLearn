package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
)

func main() {
	s := "PHP是最好的語言" //UTF-8

	fmt.Println(s)

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	fmt.Println(len(s))

	for i, ch := range s { // ch is rune
		fmt.Print(reflect.TypeOf(ch))
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune Count:", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()
	
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}

	fmt.Println(len(r))
}
