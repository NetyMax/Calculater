package consumer

import (
	"fmt"

	"github.com/Max/Calculation/Calculater/internal/task"
)

func Start(taskChan <-chan task.Task) {
	for t := range taskChan {

		var result float64

		switch t.Operator {
		case "+":
			result = t.A + t.B
		case "-":
			result = t.A - t.B
		case "/":
			if t.B == 0 {
				fmt.Println("Ошибка: деление на ноль")
				continue
			}
			result = t.A / t.B
		default:
			fmt.Println("Неизвестный оператор:", t.Operator)
			continue
		}
		fmt.Printf("Операция: %.2f %s %.2f = %.2f\n", t.A, t.Operator, t.B, result)
	}
}
