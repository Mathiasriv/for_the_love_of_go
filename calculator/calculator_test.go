package calculator_test
import (
    "calculator"
"testing"
)
func TestAdd(t *testing.T) {
    type testCase struct {
      a , b float64
      want float64}
    
    testCases := []testCase{
      {a: 2, b: 3, want: 5},
      {a: 10, b: 24, want: 34},
      {a: 1, b: 0, want: 1},
    }

    for _,tc := range(testCases){
      got := calculator.Add(tc.a, tc.b)
      if tc.want != got {t.Errorf("want %f, got %f", tc.want, got)}
      
    }
}

func TestSubtract(t *testing.T) {
    type testCase struct {
      a , b float64
      want float64}
    
    testCases := []testCase{
      {a: 3, b: 1, want: 2},
      {a: 12,b: 3, want: 9},
      {a: 6, b: 2, want: 4},
    }

    for _,tc := range(testCases){
      got := calculator.Subtract(tc.a, tc.b)
      if tc.want != got{t.Errorf("want %f, got %f", tc.want, got)}
    }


    
    
}

func TestMultiply(t *testing.T) {

  type testCase struct{
    a,b float64
    want float64
  }

  testCases := []testCase{
    {a: 2, b: 2, want: 4},
    {a: 4, b: 6, want: 24},
    {a: -2, b: -3, want: 6},
  }

  for _, tc := range(testCases){
      
      got := calculator.Multiply(tc.a, tc.b)
      if tc.want != got{t.Errorf("want %f, got %f", tc.want, got)}


  }




  
}
