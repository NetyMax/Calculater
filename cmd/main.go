package main

import (
	"fmt"
	"net/http"

	"github.com/Max/Calculation/Calculater/internal"
)

func main() {

	http.HandleFunc("/Addition", internal.CalcHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)

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

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
