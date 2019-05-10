package repositories

import (
	"go-wire-test/models"
)

type TestBookRepository struct{}

func NewTestBookRepository() BookRepository {
	return TestBookRepository{}
}

func (br TestBookRepository) GetAllBooks() []models.Book {
	return []models.Book{
		{
			BookID: 1,
			Name:   "The great adventures of a fish.",
		},
		{
			BookID: 2,
			Name:   "Tonight, you sleep with the fishes!",
		},
	}
}
