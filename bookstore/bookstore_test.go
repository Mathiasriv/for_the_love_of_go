package bookstore_test

import (
	"bookstore"
	"testing"
)

func TestBook(t *testing.T) {

	_ = bookstore.Book{
		Title:  "Nicholas Chuckleby",
		Author: "Charles Dickens",
		Copies: 8,
	}

}
