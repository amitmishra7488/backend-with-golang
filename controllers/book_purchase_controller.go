package controllers

import (
	"context"
	"fmt"
	"golang-backend/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type BookPurchaseController struct{}

var bookPurchaseRepository = new(models.BookPurchaseRepository)



func (m *BookPurchaseController) AddBookPurchase(ctx *gin.Context) {
	var bookPurchase models.BookPurchase
	idStr := ctx.Param("id")

	// Convert the id string to an int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		// If there's an error parsing the id, return a 400 Bad Request response
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Bind JSON data to bookPurchase
	if err := ctx.ShouldBindJSON(&bookPurchase); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the book by ID
	existedBook, err := bookRepository.GetBookById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Book Not Found"})
		return
	}

	// Get the user ID from the context
	userIdFloat64, _ := ctx.Get("userId")
	userId := int64(userIdFloat64.(float64))

	// Generate the purchase ID
	purchaseIDString, err := generatePurchaseID(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Assign values to bookPurchase
	bookPurchase.BookID = id
	bookPurchase.UserId = userId
	bookPurchase.PurchaseID = purchaseIDString
	bookPurchase.PurchaseDate = time.Now()
	bookPurchase.TotalPrice = existedBook.Price * float64(bookPurchase.Quantity)

	// Create the book purchase
	if err := bookPurchaseRepository.NewBookPurchase(ctx, &bookPurchase); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Book purchase added successfully"})
}
func generatePurchaseID(ctx context.Context) (string, error) {
	// Get the last purchase ID
	purchaseID, err := models.LastOrderId(ctx)
	if err != nil {
		return "", err
	}

	// Get the current year and month
	currentYear := time.Now().Year()
	currentMonth := int(time.Now().Month())

	// Format the purchase ID string with leading zeros for single-digit months
	purchaseIDString := fmt.Sprintf("%d-%02d-%d", currentYear, currentMonth, purchaseID)
	return purchaseIDString, nil
}