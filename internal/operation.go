package internal

import "fmt"

func Addition(a, b float64) float64 {
	return a + b
}

func Subtraction(a, b float64) float64 {
	return a - b
}

func Multiplication(a, b float64) float64 {
	return a * b
}

func Division(a, b float64) float64 {
	if b == 0 {
		fmt.Println("на ноль делить нельзя")
		return 0
	}

	return a / b
}
