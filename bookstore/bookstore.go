package bookstore

import "errors"

type Book struct {
	Title, Author string
	Copies        int
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no quedan copias")
	}
	b.Copies--
	return b, nil
}

func GetAllBooks(catalog []Book )[]Book{
  return catalog
}
