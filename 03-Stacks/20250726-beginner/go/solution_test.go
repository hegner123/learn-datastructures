package main

import (
	"testing"
)

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		expected   bool
	}{
		// Test Case 1: Empty string
		{
			name:       "Empty string",
			expression: "",
			expected:   true,
		},
		// Test Case 2: Single pair of parentheses
		{
			name:       "Single pair of parentheses",
			expression: "()",
			expected:   true,
		},
		// Test Case 3: Single pair of brackets
		{
			name:       "Single pair of brackets",
			expression: "[]",
			expected:   true,
		},
		// Test Case 4: Single pair of braces
		{
			name:       "Single pair of braces",
			expression: "{}",
			expected:   true,
		},
		// Test Case 5: Nested brackets of same type
		{
			name:       "Nested brackets of same type",
			expression: "((()))",
			expected:   true,
		},
		// Test Case 6: Mixed balanced brackets
		{
			name:       "Mixed balanced brackets",
			expression: "([{}])",
			expected:   true,
		},
		// Test Case 7: Unmatched opening bracket
		{
			name:       "Unmatched opening bracket",
			expression: "(()",
			expected:   false,
		},
		// Test Case 8: Unmatched closing bracket
		{
			name:       "Unmatched closing bracket",
			expression: "())",
			expected:   false,
		},
		// Test Case 9: Wrong order of brackets
		{
			name:       "Wrong order of brackets",
			expression: "([)]",
			expected:   false,
		},
		// Test Case 10: Complex nested structure
		{
			name:       "Complex nested structure",
			expression: "{[(){}][(())]}",
			expected:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsBalanced(tt.expression)
			if result != tt.expected {
				t.Errorf("IsBalanced(%q) = %v, expected %v", tt.expression, result, tt.expected)
			}
		})
	}
}