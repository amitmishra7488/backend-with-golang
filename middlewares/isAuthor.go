package middlewares

import (
	"fmt"
	"golang-backend/models" // Import your User model
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	userRepository = new(models.UserRepository)
)

func IsAuthor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Extract user ID from token
		userIdFloat64, exists := ctx.Get("userId")
		if !exists {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		userId := int64(userIdFloat64.(float64))
		fmt.Printf("%T", userId)

		// Query user role from database
		userProfile, err := userRepository.GetUserProfile(ctx, userId)
		if err != nil {
			// If there's an error fetching user profile data (e.g., database error), return a 500 Internal Server Error response
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user profile"})
			return
		}

		// Check if user role is 'author'
		fmt.Println(userProfile)
		if userProfile.Role != "author" {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to access this resource"})
			ctx.Abort()
			return
		}

		// Proceed to the next middleware or handler
		ctx.Next()
	}
}
