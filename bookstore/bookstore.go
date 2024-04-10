package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	Title, Author   string
	Copies          int
	Id              int
	PriceCents      int
	DiscountPercent int
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no quedan copias")
	}
	b.Copies--
	return b, nil
}

func GetAllBooks(catalog map[int]Book) []Book {
	result := []Book{}
	for _, c := range catalog {
		result = append(result, c)
	}
	return result
}

func GetBook(catalog map[int]Book, ID int) (Book, error) {
	b, ok := catalog[ID]

	if !ok {
		return Book{}, fmt.Errorf("el id %d no existe", ID)
	}
	return b, nil

}
// metodo (solo aplicable a un objeto en particular)
func (b Book)NetPriceCents()int{
  saving := b.PriceCents * b.DiscountPercent / 100
  return b.PriceCents - saving 

}
