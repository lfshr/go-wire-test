// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"fmt"
	"go-wire-test/repositories"
	"go-wire-test/services"
)

// Injectors from main.go:

func initializeBookStore() BookStore {
	bookRepository := repositories.NewTestBookRepository()
	bookService := services.NewTestBookService(bookRepository)
	bookStore := NewBookStore(bookService)
	return bookStore
}

// main.go:

type BookStore struct {
	BookService services.BookService
}

func NewBookStore(bs services.BookService) BookStore {
	return BookStore{
		BookService: bs,
	}
}

func main() {
	bs := initializeBookStore()
	books := bs.BookService.GetAllBooks()
	fmt.Printf("%v\n", books)
}