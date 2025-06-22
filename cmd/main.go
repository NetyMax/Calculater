package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Max/Calculation/Calculater/internal"
)

func main() {

	go func() {
		http.HandleFunc("/Addition", internal.CalcHandler)
		fmt.Println("Сервер запущен на http://localhost:8080")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Println("Ошибка запуска сервера:", err)
		}
	}()

	time.Sleep(300 * time.Millisecond)

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

	select {}
}
