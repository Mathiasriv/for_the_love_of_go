package main

import ("calculator"
"fmt") 


func main(){
  result := calculator.Add(4,2)
  fmt.Println(result)
  calculator.AddMany(1,2,3)
  su := calculator.Add(2,2)
  fmt.Println(su)

}
