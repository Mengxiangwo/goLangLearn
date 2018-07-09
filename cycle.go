package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

func convertToBin(n int) string {
	result := ""

	if n == 0 {
		return "0"
	}

	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}

	return result

}

func forfunc() int {
	sum := 0

	for i := 1; i <= 100; i++ {
		sum += i
	}

	return sum
}

func printFile(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan()  {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("forever")
	}
}

func main() {
	fmt.Println(forfunc())

	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(0),
	)

	printFile("abc.txt")

	forever()
}
