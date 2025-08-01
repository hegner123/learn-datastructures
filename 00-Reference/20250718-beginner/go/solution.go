package main

import "errors"

// Calculator represents a simple calculator for basic arithmetic operations
type Calculator struct{}

// Add performs addition of two integers
func (c *Calculator) Add(a, b int) int {
	// TODO: Implement addition
	panic("method not implemented")
}

// Subtract performs subtraction of two integers
func (c *Calculator) Subtract(a, b int) int {
	// TODO: Implement subtraction
	panic("method not implemented")
}

// Multiply performs multiplication of two integers
func (c *Calculator) Multiply(a, b int) int {
	// TODO: Implement multiplication
	panic("method not implemented")
}

// Divide performs division of two integers
// Returns an error if the divisor is zero
func (c *Calculator) Divide(a, b int) (int, error) {
	// TODO: Implement division with zero check
	return 0, errors.New("method not implemented")
}