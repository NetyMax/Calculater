package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/Max/Calculation/Calculater/internal"
)

func main() {

	go func() {
		http.HandleFunc("/Addition", internal.CalcHandler)
		fmt.Println("Сервер запущен на http://localhost:8080")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatalf("Error when server is power %v", err)
		}
	}()

	var wg sync.WaitGroup
	results := make(chan string, 10)
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			a := float64(i)
			b := float64(i * 2)
			result := internal.Addition(a, b)
			results <- fmt.Sprintf("Горутина %d: %.2f + %.2f = %.2f", i, a, b, result)
		}(i)
	}

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
