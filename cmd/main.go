package main

import (
	"fmt"
)

func main() {

	a, b, operation := internal.getInput()

	var result float64
	switch operation {
	case "+":
		result = addition(a, b)

	case "-":
		result = subtraction(a, b)

	case "*":
		result = multiplication(a, b)

	case "/":
		result = division(a, b)
	default:
		fmt.Println("не известный оператор")
		return
	}
	fmt.Printf("результат: %.2f/n", result)
}
