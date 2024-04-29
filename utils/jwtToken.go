package utils

import (
	"golang-backend/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func SecretKey() []byte {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// Retrieve JWT secret key from environment variables
	jwtSecretKey := os.Getenv("jwtsecretkey")
	if jwtSecretKey == "" {
		panic("JWT secret key environment variable is not set")
	}
	return []byte(jwtSecretKey)
}

// generateJWTToken generates a JWT token with user information
func GenerateJWTToken(user models.User) (string, error) {
	// Create a new token with custom claims
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	// Set custom claims
	claims["id"] = user.UserId
	claims["email"] = user.Email
	// You can add more custom claims as needed

	// Set expiration time for the token (e.g., 24 hours)
	expirationTime := time.Now().Add(24 * time.Hour)
	claims["exp"] = expirationTime.Unix()

	// Retrieve JWT secret key
	jwtSecretKey := SecretKey()

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
