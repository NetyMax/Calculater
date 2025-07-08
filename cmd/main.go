package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Max/Calculation/Calculater/internal"
)

type Task struct {
	A, B     float64
	Operator string
}

func main() {

	taskChan := make(chan Task)

	go func() {

		tasks := []Task{
			{A: 5, B: 7, Operator: "+"},
			{A: 8, B: 2, Operator: "-"},
			{A: 2, B: 2, Operator: "*"},
			{A: 4, B: 0, Operator: "/"},
		}
		for _, task := range tasks {
			taskChan <- task
		}
		close(taskChan)
	}()

	go func() {
		for task := range taskChan {
			var result float64
			switch task.Operator {
			case "+":
				result = task.A + task.B
			case "-":
				result = task.A - task.B
			case "/":
				if task.B == 0 {
					fmt.Println("Ошибка: деление на ноль")
					continue
				}
				result = task.A / task.B
			default:
				fmt.Println("Неизвестный оператор:", task.Operator)
				continue
			}
			fmt.Printf("Операция: %.2f %s %.2f = %.2f\n", task.A, task.Operator, task.B, result)
		}
	}()

	go func() {
		http.HandleFunc("/Addition", internal.CalcHandler)
		fmt.Println("Сервер запущен на http://localhost:8080")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatalf("Error when server is power %v", err)
		}
	}()

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
		fmt.Println("не известный оператор!")
		return
	}
	fmt.Printf("результат: %.2f \n", result)

	select {}
}
