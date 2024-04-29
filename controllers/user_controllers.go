// controllers/user_controller.go
package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-backend/models"
	"golang-backend/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

type UserController struct{}

var (
	userRepository = new(models.UserRepository)
)

// CreateUser controller function
func (m UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	// Attempt to extract user data from the JSON request body
	if err := ctx.BindJSON(&user); err != nil {
		// If there's an error (e.g., invalid JSON format), return a 400 Bad Request response
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(user.Password, user.Email, user.Username)

	// hashing the password
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		panic(err)
	}
	user.Password = string(password)
	// If user.Role is provided, validate and set it; otherwise, use the default value "user"
	if user.Role != "" {
		if err := user.SetRole(user.Role); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		user.Role = models.RoleUser
	}
	// Call the user repository's Create method to save the user data to the database
	if err := userRepository.Create(ctx, &user); err != nil {
		// If there's an error creating the user (e.g., database error), return a 500 Internal Server Error response
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		fmt.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})
}

func (m UserController) LoginUser(ctx *gin.Context) {
	var user models.User
	// Attempt to extract user data from the JSON request body
	if err := ctx.BindJSON(&user); err != nil {
		// If there's an error (e.g., invalid JSON format), return a 400 Bad Request response
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(user)
	userExists, err := userRepository.FindByEmail(ctx, user.Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User Not Found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userExists.Password), []byte(user.Password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Credentials"})
		return
	}
	// Generate JWT token with user information
	token, err := utils.GenerateJWTToken(*userExists)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate authentication token"})
		return
	}

	// Return the JWT token to the client
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (m UserController) GetAllUsers(ctx *gin.Context) {
	var users []*models.User

	users, err := userRepository.QueryAll(ctx)
	if err != nil {
		// If there's an error creating the user (e.g., database error), return a 500 Internal Server Error response
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error Finding users"})
		return
	}
	fmt.Println(users)
	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": users})
}

func (m UserController) DeleteUser(ctx *gin.Context) {
	// Extract the id parameter from the URL
	idStr := ctx.Param("id")

	// Convert the id string to an int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		// If there's an error parsing the id, return a 400 Bad Request response
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Call the user repository's DeleteUserById method to delete the user
	if err := userRepository.DeleteUserById(ctx, id); err != nil {
		// If there's an error deleting the user (e.g., database error), return a 500 Internal Server Error response
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}

	// If the user is successfully deleted, return a 200 OK response
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (m UserController) UpdateUser(ctx *gin.Context) {
	// Extract the id parameter from the URL
	idStr := ctx.Param("id")

	// Convert the id string to an int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		// If there's an error parsing the id, return a 400 Bad Request response
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Attempt to extract the updated user data from the JSON request body
	var updatedUser models.User
	if err := ctx.BindJSON(&updatedUser); err != nil {
		// If there's an error (e.g., invalid JSON format), return a 400 Bad Request response
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the user repository's UpdateUserById method to update the user
	if err := userRepository.UpdateUserById(ctx, id, &updatedUser); err != nil {
		// If there's an error updating the user (e.g., database error), return a 500 Internal Server Error response
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	// If the user is successfully updated, return a 200 OK response
	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (m UserController) GetUserProfile(ctx *gin.Context) {
    // Extract user information from the request context
    userIdFloat64, _ := ctx.Get("userId")
    userId := int64(userIdFloat64.(float64))
    userProfile, err := userRepository.GetUserProfile(ctx, userId)
    if err != nil {
        // If there's an error fetching user profile data (e.g., database error), return a 500 Internal Server Error response
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user profile"})
        return
    }
    // Return user profile data
    ctx.JSON(http.StatusOK, gin.H{"userProfile": userProfile})
}
