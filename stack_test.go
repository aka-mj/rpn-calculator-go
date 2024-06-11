package main

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := new(Stack)

	// Test empty stack
	if s.Len() != 0 {
		t.Errorf("s.Len() = %d, want 0", s.Len())
	}
	if !s.Empty() {
		t.Errorf("s.Empty() = false, want true")
	}

	// Test stack with one element
	s.Push(1)
	if s.Len() != 1 {
		t.Errorf("s.Len() = %d, want 1", s.Len())
	}
	if s.Empty() {
		t.Errorf("s.Empty() = true, want false")
	}

	// Test stack with two elements
	s.Push(2)
	if s.Len() != 2 {
		t.Errorf("s.Len() = %d, want 2", s.Len())
	}
	if s.Empty() {
		t.Errorf("s.Empty() = true, want false")
	}

	// Test Pop
	if f, ok := s.Pop(); !ok || f != 2 {
		t.Errorf("s.Pop() = %f, %t, want 2, true", f, ok)
	}
	if s.Len() != 1 {
		t.Errorf("s.Len() = %d, want 1", s.Len())
	}

	if f, ok := s.Pop(); !ok || f != 1 {
		t.Errorf("s.Pop() = %f, %t, want 1, true", f, ok)
	}
	if s.Len() != 0 {
		t.Errorf("s.Len() = %d, want 0", s.Len())
	}
	if !s.Empty() {
		t.Errorf("s.Empty() = false, want true")
	}
	if f, ok := s.Pop(); ok {
		t.Errorf("s.Pop() = %f, %t, want _, false", f, ok)
	}
}
