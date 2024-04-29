package controllers

import (
	"fmt"
	"golang-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	authorRevenueRepository = new(models.AuthorRevenueRepository)
)

type AuthorRevenueController struct{}

func (m AuthorRevenueController) GetAuthorRevenue(ctx *gin.Context) {
	var data []models.AuthorRevenue
	userIdFloat64, _ := ctx.Get("userId")
	userId := int64(userIdFloat64.(float64))

	revenueData, err := authorRevenueRepository.GetRevenue(ctx, userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Book Not Found"})
		return
	}
	data = revenueData
	fmt.Println(revenueData)
	ctx.JSON(http.StatusOK, gin.H{"data": data})

}
