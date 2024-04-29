package models

import (
	"context"
	"errors"
	"fmt"
	"golang-backend/db"
	"strconv"
	"strings"
	"time"

	"github.com/uptrace/bun"
)

// BookPurchase represents a purchase of a book.
type BookPurchase struct {
	bun.BaseModel  `bun:"table:book_purchases"`
	BookPurchaseID int64     `bun:"book_purchase_id,pk,autoincrement" json:"book_purchase_id"`
	BookID         int64     `bun:"book_id" json:"book_id"`
	PurchaseID     string    `bun:"purchase_id" json:"purchase_id"`
	PurchaseDate   time.Time `bun:"purchase_date" json:"purchase_date"`
	Quantity       int64     `bun:"quantity" json:"quantity"`
	TotalPrice     float64   `bun:"total_price" json:"total_price"`
	UserId         int64     `bun:"user_id" json:"user_id"`
}

type BookPurchaseRepository struct{}

func (m BookPurchaseRepository) NewBookPurchase(ctx context.Context, purchaseData *BookPurchase) error {
	book, err := db.GetDB().NewInsert().Model(purchaseData).Exec(ctx)
	fmt.Println(book)
	return err
}

func LastOrderId(ctx context.Context) (int64, error) {
    var lastPurchaseID string
    query := `SELECT purchase_id FROM book_purchases ORDER BY purchase_date DESC LIMIT 1`
    err := db.GetDB().QueryRowContext(ctx, query).Scan(&lastPurchaseID)
    if err != nil {
        return 0, err
    }

    // Extract the last order number from the purchase ID string
    parts := strings.Split(lastPurchaseID, "-")
    if len(parts) > 3 {
        return 0, errors.New("invalid purchase ID format")
    }
    lastID, err := strconv.ParseInt(parts[2], 10, 64)
    if err != nil {
        return 0, err
    }

    // Increment the obtained lastID by one to get the next purchase ID
    lastID++
    return lastID, nil
}

