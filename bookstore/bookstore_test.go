package bookstore_test

import (
	"bookstore"
	"sort"
	"testing"
  
  "github.com/google/go-cmp/cmp/cmpopts"

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
	catalog := bookstore.Catalog{
		1: {Id: 1, Title: "For the Love of Go"},
		2: {Id: 2, Title: "The Power of Go: Tools"},
	}

	want := []bookstore.Book{
		{Id: 1, Title: "For the Love of Go"},
		{Id: 2, Title: "The Power of Go: Tools"},
	}
	got := catalog.GetAllBooks()
	sort.Slice(got, func(i, j int) bool {
		return got[i].Id < got[j].Id
	})
	 
  if !cmp.Equal(want, got,
cmpopts.IgnoreUnexported(bookstore.Book{})){
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {Id: 1,
			Title: "For the love of go"},
		2: {Id: 2,
			Title: "For the love of go: tools"},
	}
	want := bookstore.Book{Id: 2, Title: "For the love of go: tools"}

	got, err := catalog.GetBook(2)
  if !cmp.Equal(want, got,
cmpopts.IgnoreUnexported(bookstore.Book{})){
		t.Error(cmp.Diff(want, got))
	}

	if err != nil {
		t.Fatal(err)
	}

}

func TestGetBookBadIDReturnsError(t *testing.T) {
	t.Parallel()

	catalog := bookstore.Catalog{}
	_, err := catalog.GetBook(999)
	if err == nil {
		t.Fatal("want error for non-existent ID, got nil")
	}
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:           "Spark joy",
		Author:          "Marie Kondo",
		Copies:          2,
		DiscountPercent: 10,
		PriceCents:      500,
	}
	got := b.NetPriceCents()

	want := 450

	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}

}

func TestSetPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{Title: "For the love of Go", PriceCents: 300}
	want := 200
	err := b.SetPriceCents(want)
	if err != nil {
		t.Fatal(err)
	}
	got := b.PriceCents
	if got != want {
		t.Errorf("want %d != got %d", want, got)
	}

}

func TestSetInvalidPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{Title: "For the love of Go", PriceCents: 300}
	err := b.SetPriceCents(-1)
	if err == nil {
		t.Fatal("want error setting invalid price -1, got nil")
	}

}

func TestSetInvalidCategory(t *testing.T){
  t.Parallel()
  b := bookstore.Book{Title: "For the love of go",}
  err := b.SetCategory(999)
  if err == nil {
    t.Fatal("want error setting invalid category")
  }

  
}

/* func TestSetValidCategory(t *testing.T){
  t.Parallel()
  b := bookstore.Book{Title:"For the love of Go",}
  want := CategoryAutobiography 
  err := b.SetCategory(want)
  if err != nil {
    t.Fatal(err)
  }
  got := b.Category()
  
  if want != got{

    t.Errorf("want category %q, got %q", want, got)
  }


} */

func TestSetCategory(t *testing.T){
  t.Parallel()
  
  b := bookstore.Book{Title: "For the love of go",}
  
cats := []bookstore.Category{
        bookstore.CategoryAutobiography,
        bookstore.CategoryLargePrintRomance,
        bookstore.CategoryParticlePhysics,
    }

    for _, cat := range cats{
      err := b.SetCategory(cat)
     
      if err != nil{
       t.Fatal() 
      }
      got := b.Category()
      if got != cat {
t.Errorf("want category %q, got %q", cat, got)
      }
    }      
  }

