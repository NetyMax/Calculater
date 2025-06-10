package internel

import "fmt"

func getInput() (float64, float64, string) {
	var a, b float64
	var operation string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	fmt.Print("Введите оператор (+, -, *, /): ")
	fmt.Scanln(&operation)

	return a, b, operation

}

// end
