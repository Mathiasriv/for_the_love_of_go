// Package calculator does simple calculations.
package calculator

import (
	"errors"
	"fmt"
	"math"
)

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


func Sqrt(input float64)(float64, error){
  
  if input < 0 {
    return 0, fmt.Errorf("no permite numeros negativos %f", input)

  }

  return math.Sqrt(input), nil


}


func AddMany(inputs ...int){
  suma := 0
  for _,num := range inputs{
    suma = suma + num 
   fmt.Println(num)  }
   


}
