package bookstore_test

import (
	"bookstore"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
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
func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := map[int]bookstore.Book{
		1: {Id: 1, Title: "For the Love of Go"},
		2: {Id: 2, Title: "The Power of Go: Tools"},
	}

	want := []bookstore.Book{
		{Id: 1, Title: "For the Love of Go"},
		{Id: 2, Title: "The Power of Go: Tools"},
	}
	got := bookstore.GetAllBooks(catalog)
	sort.Slice(got, func(i, j int) bool {
		return got[i].Id < got[j].Id
	})
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := map[int]bookstore.Book{
		1: {Id: 1,
			Title: "For the love of go"},
		2: {Id: 2,
			Title: "For the love of go: tools"},
	}
	want := bookstore.Book{Id: 2, Title: "For the love of go: tools"}

	got, err := bookstore.GetBook(catalog, 2)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

	if err != nil {
		t.Fatal(err)
	}

}

func TestGetBookBadIDReturnsError(t *testing.T) {
	t.Parallel()
	catalog := map[int]bookstore.Book{}
	_, err := bookstore.GetBook(catalog, 999)
	if err == nil {
		t.Fatal("want error for non-existent ID, got nil")
	}
}
