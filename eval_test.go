package main

import (
	"testing"
)

func TestPerformOPeration(t *testing.T) {
	tests := []struct {
		name string
		a, b float64
		op   string
		want float64
		err  bool
	}{
 		{"add", 1, 2, "+", 3, false},
		{"sub", 1, 2, "-", -1, false},
		{"multiply", 1, 2, "*", 2, false},
		{"divide", 1, 2, "/", 0.5, false},
		{"divide by zero", 1, 0, "/", 0, true},
		{"invalid operation", 1, 2, "invalid", 0, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := performOperation(test.a, test.b, test.op)
			// Check if the error is expected
			if test.err {
				if err == nil {
					t.Errorf("preformOperation(%v, %v, %v) = %v, %v; want _, error", test.a, test.b, test.op, got, err)
				}
				return
			}
			// Unexpected error
			if err != nil {
				t.Errorf("preformOperation(%v, %v, %v) = _, %v; want %v, nil", test.a, test.b, test.op, err, test.want)
			}
			// Check we got the expected result
			if got != test.want {
				t.Errorf("preformOperation(%v, %v, %v) = %v, _; want %v, _", test.a, test.b, test.op, got, test.want)
			}
		})
	}
}



func TestEvalRPN(t *testing.T) {
	tests := []struct {
		name string
		expr string
		want float64
		err  bool
	}{
		{"add", "1 2 +", 3, false},
		{"sub", "1 2 -", -1, false},
		{"multiply", "1 2 *", 2, false},
		{"divide", "1 2 /", 0.5, false},
		{"complex", "1 2 + 3 *", 9, false},
		{"complex", "1 2 3 + *", 5, false},
		{"complex", "1 2 3 + * 4 /", 1.25, false},
		{"single expression", "1", 1, false},
		{"divide by zero", "1 0 /", 0, true},
		{"invalid token", "1 invalid +", 0, true},
		{"not enough operands", "+", 0, true},
		{"empty expression", "", 0, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := evalRPN(test.expr)
			// Check if the error is expected
			if test.err {
				if err == nil {
					t.Errorf("evalRPN(%v) = %v, %v; want _, error", test.expr, got, err)
				}
				return
			}
			// Unexpected error
			if err != nil {
				t.Errorf("evalRPN(%v) = _, %v; want %v, nil", test.expr, err, test.want)
			}
			// Check we got the expected result
			if got != test.want {
				t.Errorf("evalRPN(%v) = %v, _; want %v, _", test.expr, got, test.want)
			}
		})
	}
}

func BenchmarkEvalRPN(b *testing.B) {
	expr := "1 2 + 3 * 4 /"
	for i := 0; i < b.N; i++ {
		evalRPN(expr)
	}
}
