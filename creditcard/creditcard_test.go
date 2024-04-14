package creditcard_test

import("testing"
"creditcard")


func TestCard(t *testing.T){
  t.Parallel()

  _ := creditcard.NewCard()
}
