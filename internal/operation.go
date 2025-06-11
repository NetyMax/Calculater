package internal

import "fmt"

func addition(a, b float64) float64 {
	return a + b
}

func subtraction(a, b float64) float64 {
	return a - b
}

func multiplication(a, b float64) float64 {
	return a * b
}

func division(a, b float64) float64 {
	if b == 0 {
		fmt.Println("на ноль делить нельзя")
		return 0
	}

	return a / b
}
