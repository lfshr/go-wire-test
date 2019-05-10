//+build wireinject

package main

import (
	"fmt"
	. "go-wire-test/repositories"
	. "go-wire-test/services"

	"github.com/google/wire"
)

type BookStore struct {
	BookService BookService
}

func NewBookStore(bs BookService) BookStore {
	return BookStore{
		BookService: bs,
	}
}

func main() {
	bs := initializeBookStore()
	books := bs.BookService.GetAllBooks()

	fmt.Printf("%v\n", books)
}

func initializeBookStore() BookStore {
	wire.Build(
		NewTestBookRepository,
		NewTestBookService,
		NewBookStore,
	)
	return BookStore{}
}
