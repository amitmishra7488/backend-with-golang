package routes

import (
	"golang-backend/controllers"
	"golang-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
    // Initialize the UserController
    userController := new(controllers.UserController)
    bookPuchaseController := new(controllers.BookPurchaseController)

    // Define routes for user related operations
    userRoutes := router.Group("/users")
    {
        // Call CreateUser method on the userController instance
        userRoutes.POST("/signup", userController.CreateUser)
        userRoutes.POST("/login", userController.LoginUser)
        userRoutes.GET("/", userController.GetAllUsers)
        userRoutes.GET("/profile", middlewares.AuthMiddleware(), userController.GetUserProfile) // Add route for fetching user profile
        userRoutes.DELETE("/:id", userController.DeleteUser)
        userRoutes.PUT("/:id", userController.UpdateUser)
        userRoutes.POST("/buy-book/:id", middlewares.AuthMiddleware(), bookPuchaseController.AddBookPurchase)
        // Add more routes for other user operations (e.g., GetUsers, UpdateUser, DeleteUser)
    }
}
