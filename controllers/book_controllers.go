package controllers

import (
	"golang-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type BookController struct{}

var (
	bookRepository = new(models.BookRepository)
)

// Add a new book
func (m *BookController) AddBook(ctx *gin.Context) {
	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userIdFloat64, _ := ctx.Get("userId")
	userId := int64(userIdFloat64.(float64))
	book.AuthorId = userId
	if book.LaunchDate.IsZero() {
		book.LaunchDate = time.Now()
	}
	if err := bookRepository.AddNewBook(ctx, &book); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Book added successfully"})
}
