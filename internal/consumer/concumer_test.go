package consumer_test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/Max/Calculation/Calculater/internal/consumer"
	"github.com/Max/Calculation/Calculater/internal/task"
)

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w

	f()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestStart(t *testing.T) {

	ch := make(chan task.Task)
	go func() {
		ch <- task.Task{A: 5, B: 3, Operator: "+"}
		ch <- task.Task{A: 10, B: 4, Operator: "-"}
		ch <- task.Task{A: 9, B: 3, Operator: "/"}
		ch <- task.Task{A: 8, B: 0, Operator: "/"}
		ch <- task.Task{A: 2, B: 2, Operator: "%"}
		close(ch)

	}()

	output := captureOutput(func() {
		consumer.Start(ch)
	})
	tests := []struct {
		substr string
	}{
		{substr: "Операция: 5.00 + 3.00 = 8.00"},
		{substr: "Операция: 10.00 - 4.00 = 6.00"},
		{substr: "Операция: 9.00 / 3.00 = 3.00"},
		{substr: "Ошибка: деление на ноль"},
		{substr: "Неизвестный оператор"},
	}

	for _, tt := range tests {
		if !strings.Contains(output, tt.substr) {
			t.Errorf("Ожидали, что строка содержит: %q, но не нашли\n Вывод:\n %s", tt.substr, output)
		}
	}
}
