// Package calculator does simple calculations.
package calculator

import "errors"

// Add takes two numbers and returns the result of adding
// them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract toma dos números a y b, y
// devuelve el resultado de restar b de a.
func Subtract(a, b float64) float64 {
	if a > b {
		return a - b
	} else {
		return b - a
	}

}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("disivion por cero no permitida")
	}
	return a / b, nil
}
