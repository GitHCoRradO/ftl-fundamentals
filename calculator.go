// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

func AddMany(inputs ...float64) float64 {
	var acc float64 = 0
	for _, input := range inputs {
		acc = acc + input
	}
	return acc
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return b - a
}

func SubtractMany(inputs ...float64) float64 {
	var acc float64
	for _, input := range inputs {
		acc = acc + input
	}
	return acc
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b != 0 {
		return a / b, nil
	} else {
		return 0, fmt.Errorf("bad input: %f, %f (division by zero is undefined)", a, b)
	}
}

func Sqrt(a float64) (float64, error) {
	if a >= 0 {
		return math.Sqrt(a), nil
	} else {
		return 0, fmt.Errorf("bad input: %f (square root of a negative number is invalid", a)
	}
}
