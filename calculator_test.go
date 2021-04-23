package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

type testCase struct {
	a, b float64
	want float64
}

type testCaseDivide struct {
	a, b float64
	want float64
	errExcepted bool
}

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 5, want: 6},
		{a: 2, b: -2, want: 0},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 0},
		{a: 1, b: 5, want: 4},
		{a: 2, b: -2, want: -4},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}

}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 5, want: 5},
		{a: 2, b: -2, want: -4},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []testCaseDivide {
		{a: 2, b: 2, want: 1, errExcepted: false},
		{a: 1, b: 5, want: 0.2, errExcepted: false},
		{a: 2, b: -2, want: -1, errExcepted: false},
		{a: 2, b: 0, want: 0, errExcepted: true},
		{a: 3, b: 0, want: 0, errExcepted: true},
		{a: 4, b: 0, want: 0, errExcepted: true},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := err != nil
		if tc.errExcepted != errReceived {
			t.Fatalf("Divide(%f, %f): unexpected error status: %v", tc.a, tc.b, errReceived)
		}
		if !tc.errExcepted && got != tc.want {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	type testCase struct {
		a float64
		want float64
		errExpected bool
	}
	t.Parallel()
	testCases := []testCase {
		{a: 4, want: 2, errExpected: false},
		{a: 100, want: 10, errExpected: false},
		{a: 0, want: 0, errExpected: false},
		{a: -0, want: 0, errExpected: false},
		{a: -4, want: 0, errExpected: true},
		{a: 21, want: 4.582576, errExpected: false},
	}
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("Sqrt(%f): Unexpected error status: %v", tc.a, errReceived)
		}
		if !closeEnough(got, tc.want, 0.0001) {
			t.Errorf("want: %f, got %f, precision 0.0001", tc.want, got)
		}
	}

}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a - b) <= tolerance
}
