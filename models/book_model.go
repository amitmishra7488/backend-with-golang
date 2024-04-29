package models

import (
	"context"
	"errors"
	"golang-backend/db"
	"time"

	"github.com/uptrace/bun"
)

type Book struct {
    bun.BaseModel `bun:"table:books"`
    BookId        int64     `bun:"book_id,pk,autoincrement" json:"book_id"`
    Title         string    `bun:"title" json:"title"`
    AuthorId      int64     `bun:"author_id" json:"author_id"`
    Price         float64   `bun:"price" json:"price"`
    LaunchDate    time.Time `bun:"launch_date" json:"launch_date"`
}


func (b *Book) Validate() error {
	if b.Price < 100 || b.Price > 1000 {
		return errors.New("price must be between 100 and 1000")
	}
	return nil
}

type BookRepository struct{}

func (m BookRepository) AddNewBook(ctx context.Context, book *Book) error {
	_, err := db.GetDB().NewInsert().Model(book).Exec(ctx)
	return err
}

func (m BookRepository) GetBookById(ctx context.Context, bookId int64) (*Book, error) {
	var book Book
	err := db.GetDB().NewSelect().Model(&book).
        Where("book_id = ?", bookId).
        Scan(ctx)
    if err != nil {
        return nil, err
    }
    return &book, nil
}