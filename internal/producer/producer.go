package producer

import (
	"fmt"

	"github.com/Max/Calculation/Calculater/internal/task"
)

func Start(taskChan chan<- task.Task) {
	tasks := []task.Task{
		{A: 5, B: 3, Operator: "+"},
		{A: 10, B: 4, Operator: "-"},
		{A: 9, B: 3, Operator: "/"},
		{A: 8, B: 0, Operator: "/"},
	}

	for _, t := range tasks {
		fmt.Println("Producer отправил:", t)
		taskChan <- t
	}

	close(taskChan)
}

//ууу
