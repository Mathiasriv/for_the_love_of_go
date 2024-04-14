package creditcard
import ("errors")
type card struct{
  number string
}  


func NewCard(number string)(card,error){

  newCard, ok := c{number: number}

  if !ok {
     return card{}, errors.New("error de creacion")  
  }
return newCard,nil  
}

