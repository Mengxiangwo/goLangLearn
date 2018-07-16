package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

func eval(a, b int, op string) (int ,error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

func div2(a, b int) (q, r int) {
	return a / b, a % b
}

func div3(a, b int) (q1, r1 int) {
	q1 = a / b
	r1 = a % b

	return
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()

	fmt.Printf("Calling function %s with args " +
		"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(number ...int) int {
	sum := 0
	for i := range number {
		sum += number[i]
	}

	return sum
}

func swap(a, b int) (int, int) {
	b, a = a, b
	return a, b
}


func swap2(a, b *int) {
	*b, *a = *a, *b
}

func main() {
	fmt.Println(div(13, 3))

	q, r := div2(13, 3)
	fmt.Println(q, r)

	q1, r1 := div3(13, 3)
	fmt.Println(q1, r1)

	q2, _ := div3(13, 3)
	fmt.Println(q2)

	/*result, error := eval(3, 4, "X")

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(result)
	}*/

	if result, error := eval(3, 4, "+"); error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(result)
	}

	fmt.Println(apply(pow, 3, 4))

	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5, 6))

	a, b := 3, 4
	fmt.Println(swap(a, b))
	fmt.Println(a, b)

	swap2(&a, &b)
	fmt.Println(a, b)
}
