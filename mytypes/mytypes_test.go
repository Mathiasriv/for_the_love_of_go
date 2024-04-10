package mytypes_test

import ("testing"
"mytypes")


func TestTwice(t *testing.T){
  t.Parallel()
  input := mytypes.MyInt(9)
  want := mytypes.MyInt(18)
  got := input.Twice()

  if want != got {
    t.Errorf("twice %d: want %d, got %d", input, want, got)
  }
}



func TestMyStringLen(t *testing.T){
  t.Parallel()
  input := mytypes.MyString("hola")
  want := 4
  got := input.Len()
  if want != got{t.Errorf("want != got")}

}
