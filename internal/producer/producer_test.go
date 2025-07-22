package producer_test

import (
	"testing"

	"github.com/Max/Calculation/Calculater/internal/producer"
	"github.com/Max/Calculation/Calculater/internal/task"
)

func TestProduce(t *testing.T) {
	ch := make(chan task.Task, 4)

	go producer.Start(ch)

	expected := []task.Task{
		{A: 5, B: 3, Operator: "+"},
		{A: 10, B: 4, Operator: "-"},
		{A: 9, B: 3, Operator: "/"},
		{A: 8, B: 0, Operator: "/"},
	}
	var received []task.Task
	for t := range ch {
		received = append(received, t)
	}

	if len(received) != len(expected) {
		t.Fatalf("ожидали %d задач, получили %d", len(expected), len(received))
	}

	for i, got := range received {
		want := expected[i]
		if got != want {
			t.Errorf("задача %d: ожидали %+v, получили %+v", i, want, got)
		}
	}
}
