package main

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	calc := NewCalculator()

	// Test case 1: Basic Addition
	result := calc.Add(5, 3)
	expected := 8
	if result != expected {
		t.Errorf("Add(5, 3) = %d; expected %d", result, expected)
	}

	// Test case 2: Basic Subtraction
	result = calc.Subtract(10, 4)
	expected = 6
	if result != expected {
		t.Errorf("Subtract(10, 4) = %d; expected %d", result, expected)
	}

	// Test case 3: Basic Multiplication
	result = calc.Multiply(7, 2)
	expected = 14
	if result != expected {
		t.Errorf("Multiply(7, 2) = %d; expected %d", result, expected)
	}

	// Test case 4: Basic Division
	result, err := calc.Divide(15, 3)
	expected = 5
	if err != nil {
		t.Errorf("Divide(15, 3) returned error: %v", err)
	}
	if result != expected {
		t.Errorf("Divide(15, 3) = %d; expected %d", result, expected)
	}

	// Test case 5: Division with Remainder
	result, err = calc.Divide(10, 3)
	expected = 3
	if err != nil {
		t.Errorf("Divide(10, 3) returned error: %v", err)
	}
	if result != expected {
		t.Errorf("Divide(10, 3) = %d; expected %d", result, expected)
	}

	// Test case 6: Division by Zero
	_, err = calc.Divide(10, 0)
	if err == nil {
		t.Error("Divide(10, 0) should return error")
	}

	// Test case 7: Negative Numbers
	result = calc.Add(-5, 3)
	expected = -2
	if result != expected {
		t.Errorf("Add(-5, 3) = %d; expected %d", result, expected)
	}

	// Test case 8: Zero Operations
	result = calc.Multiply(0, 5)
	expected = 0
	if result != expected {
		t.Errorf("Multiply(0, 5) = %d; expected %d", result, expected)
	}

	// Test case 9: Large Numbers
	result = calc.Add(1000, 2000)
	expected = 3000
	if result != expected {
		t.Errorf("Add(1000, 2000) = %d; expected %d", result, expected)
	}

	// Test case 10: Subtraction Result Negative
	result = calc.Subtract(3, 5)
	expected = -2
	if result != expected {
		t.Errorf("Subtract(3, 5) = %d; expected %d", result, expected)
	}
}