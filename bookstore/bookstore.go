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

type Catalog map[int]Book

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no quedan copias")
	}
	b.Copies--
	return b, nil
}

func (c Catalog) GetAllBooks() []Book {
	result := []Book{}
	for _, catalog := range c {
		result = append(result, catalog)
	}
	return result
}

func (c Catalog) GetBook(ID int) (Book, error) {
	b, ok := c[ID]

	if !ok {
		return Book{}, fmt.Errorf("el id %d no existe", ID)
	}
	return b, nil

}

// metodo (solo aplicable a un objeto en particular)
func (b Book) NetPriceCents() int {
	saving := b.PriceCents * b.DiscountPercent / 100
	return b.PriceCents - saving

}

func (b Book) SetPriceCents(NewPrice int){
  b.PriceCents = NewPrice 
}
