package main

import (
	"slices"
	"strings"
)

// TODO: Implement a function that checks if parentheses, brackets, and braces are balanced using a stack
//
// The function should:
// 1. Use a stack to keep track of opening brackets
// 2. For each opening bracket '(', '[', '{', push it onto the stack
// 3. For each closing bracket ')', ']', '}', check if it matches the most recent opening bracket
// 4. Return true if all brackets are properly balanced, false otherwise
//
// Function signature:
// func IsBalanced(expression string) bool
//
// You may need to implement a simple stack or use a slice as a stack.
// Remember: A stack follows LIFO (Last In, First Out) principle.

type Stack struct {
	data   []string
	length int
}

func NewStack() *Stack {
	return &Stack{data: make([]string, 0), length: 0}
}

func (s *Stack) Push(in string) {
	s.data = append(s.data, in)
	s.length++
}

func (s *Stack) Pop() string {
	if s.length < 1 {
		return ""
	}
	out := s.data[s.length-1]
	if s.length > 0 {
		s.data = s.data[:s.length-1]
	} else {
		s.data = make([]string, 0)
	}
	s.length--
	return out
}
func (s *Stack) Peek() string {
	if s.length == 0 {
		return ""
	}
	out := s.data[s.length-1]
	return out
}

func IsBalanced(expression string) bool {
	// TODO: Implement this function
	if expression == "" {
		return true
	}
	in := strings.Split(expression, "")
	s := NewStack()
	valid := []string{"(", "[", "{", ")", "]", "}"}
	for _, input := range in {
		if !slices.Contains(valid, input) {
			return false
		}

		if isOpeningBracket(input) {
			s.Push(input)
		}
		if isClosingBracket(input) {

			if isMatchingPair(s.Peek(), input) {
				s.Pop()
			} else {
				return false
			}
		}
	}
	return s.length == 0
}

// Helper function - you may implement this if needed
func isOpeningBracket(char string) bool {
	// TODO: Check if the character is an opening bracket
	open := []string{"(", "[", "{"}
	return slices.Contains(open, char)
}

// Helper function - you may implement this if needed
func isClosingBracket(char string) bool {
	// TODO: Check if the character is a closing bracket
	closed := []string{")", "]", "}"}
	return slices.Contains(closed, char)
}

func isMatchingPair(open, closed string) bool {
	switch open {
	case "(":
		return closed == ")"
	case "[":
		return closed == "]"
	case "{":
		return closed == "}"
	default:
		return false
	}
}
