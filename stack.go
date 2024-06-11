// A simple stack using a slice of float64

package main

type Stack struct {
	data []float64
}

// Create a new stack
func (s *Stack) Push(f float64) {
	s.data = append(s.data, f)
}

// Pop the top element from the stack
func (s *Stack) Pop() (float64, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	f := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return f, true
}

// Return the number of elements in the stack
func (s *Stack) Len() int {
	return len(s.data)
}

// Check if the stack is empty
func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

// Clear the stack
func (s *Stack) Clear() {
	s.data = nil
}

