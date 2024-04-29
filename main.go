package main

import (
	"golang-backend/db"
	"golang-backend/middlewares"
	"golang-backend/routes" // Import the routes package
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	if err := db.InitDB(); err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	// Initialize the Gin router
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())
	// Define route for the base URL
	r.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, "<h1>Welcome To The Golang-Backend!/<h1>")
	})

	// Setup user routes
	routes.SetupUserRoutes(r)
	routes.SetupAuthorsRoutes(r)
	
	// Run the server on port 8001
	r.Run(":8001")
}
