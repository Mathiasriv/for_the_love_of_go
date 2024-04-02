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

func TestBuy(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Title:  "Spark joy",
		Author: "Marie Kondo",
		Copies: 2,
	}

	want := 1

	result, err := bookstore.Buy(b)

	got := result.Copies

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}

	if err != nil {
		t.Fatal(err)
	}

}

func TestBuyBookLeftCopies(t *testing.T) {
	b := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Maria Kondo",
		Copies: 0}

	_, err := bookstore.Buy(b)

	if err == nil {
		t.Error("want error buying zero copies, got nill")
	}
}
