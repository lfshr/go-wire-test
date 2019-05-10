package repositories

import "go-wire-test/models"

type BookRepository interface {
	GetAllBooks() []models.Book
}
