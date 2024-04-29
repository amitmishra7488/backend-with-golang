package routes

import (
	"golang-backend/controllers"
	"golang-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupAuthorsRoutes(router *gin.Engine) {
	bookController := new(controllers.BookController)
	authorController := new(controllers.AuthorRevenueController)

	bookRoutes := router.Group("/books") 
	{
		bookRoutes.POST("/add",middlewares.AuthMiddleware(),middlewares.IsAuthor(), bookController.AddBook)
		
	}
	authorRoutes := router.Group("/author")
	{
		authorRoutes.GET("/revenue", middlewares.AuthMiddleware(), middlewares.IsAuthor(), authorController.GetAuthorRevenue)
	}
}