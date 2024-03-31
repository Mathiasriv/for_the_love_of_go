package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 2, b: 3, want: 5},
		{a: 10, b: 24, want: 34},
		{a: 1, b: 0, want: 1},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}

	}
}

func TestSubtract(t *testing.T) {
	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 3, b: 1, want: 2},
		{a: 12, b: 3, want: 9},
		{a: 6, b: 2, want: 4},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}

}

func TestMultiply(t *testing.T) {

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 4, b: 6, want: 24},
		{a: -2, b: -3, want: 6},
	}

	for _, tc := range testCases {

		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}

	}

}

func TestDivide(t *testing.T) {

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 3, b: 3, want: 1},
		{a: 10, b: 5, want: 2},
		{a: 9, b: 3, want: 3},
	}

	for _, tc := range testCases {

		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}
		if tc.want != got {
			t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)

		}

		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Divide(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}

	}

}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Divide(1, 0)
	if err == nil {
		t.Error("want error for invalid input, got nil")
	}
}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}
