package internal

import (
	"math"
	"testing"
)

const epsilon = 1e-9

func floatsEqual(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}

func TestAddition(t *testing.T) {
	result := Addition(2, 3)
	expected := 5.0

	if !floatsEqual(result, expected) {
		t.Errorf("Addition(2,3): ожидалось %.2f, получилось %.2f", expected, result)
	}
}

func TestSubtraction(t *testing.T) {
	result := Subtraction(6, 4)
	expected := 2.0

	if !floatsEqual(result, expected) {
		t.Errorf("Subtraction(6,4): ожидалось %.2f, получилось %.2f", expected, result)
	}
}

func TestMultiplication(t *testing.T) {
	result := Multiplication(2, 2)
	expected := 4.0

	if !floatsEqual(result, expected) {
		t.Errorf("Multiplication(2,2): ожидалось %.2f, получилось %.2f", expected, result)
	}
}

func TestDivision(t *testing.T) {
	result := Division(8, 2)
	expected := 4.0

	if !floatsEqual(result, expected) {
		t.Errorf("Division(8,2): ожидалось %.2f, получилось %.2f", expected, result)
	}
}

func TestDivisionByZero(t *testing.T) {
	result := Division(5, 0)
	expected := 0.0

	if !floatsEqual(result, expected) {
		t.Errorf("Division(5,0): ожидалось %.2f, получилось %.2f", expected, result)
	}
}
