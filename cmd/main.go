package main

import (
	"fmt"

	"github.com/Max/Calculation/Calculater/internal"
)

func main() {

	a, b, operation := internal.GetInput()

	var result float64
	switch operation {
	case "+":
		result = internal.Addition(a, b)

	case "-":
		result = internal.Subtraction(a, b)

	case "*":
		result = internal.Multiplication(a, b)

	case "/":
		result = internal.Division(a, b)
	default:
		fmt.Println("не известный оператор")
		return
	}
	fmt.Printf("результат: %.2f \n", result)
}
