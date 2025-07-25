package consumer

import (
	"errors"
	"fmt"

	"github.com/Max/Calculation/Calculater/internal/task"
)

func Start(taskChan <-chan task.Task) {
	for t := range taskChan {
		result, err := DefineOperation(t)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Операция: %.2f %s %.2f = %.2f\n", t.A, t.Operator, t.B, result)
	}
}

func DefineOperation(t task.Task) (float64, error) {
	switch t.Operator {
	case "+":
		return t.A + t.B, nil
	case "-":
		return t.A - t.B, nil
	case "/":
		if t.B == 0 {
			return 0, errors.New("ошибка: деление на ноль")
		}
		return t.A / t.B, nil
	default:
		return 0, fmt.Errorf("неизвестный оператор: %s", t.Operator)
	}
}
