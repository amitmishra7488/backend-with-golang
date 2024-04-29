// db/db.go
package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var db *bun.DB

func InitDB() error {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("failed to load environment variables: %w", err)
	}

	// Retrieve MongoDB connection string from environment variables
	dbURI := os.Getenv("DB_URI")
	if dbURI == "" {
		return fmt.Errorf("DB_URI environment variable is not set")
	}

	// Create a context for database operations
	ctx := context.Background()

	// Open a PostgreSQL database connection
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dbURI)))
	db = bun.NewDB(pgdb, pgdialect.New())

	// Add query hook for debugging
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	// Check if the database connection is successful
	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	// Log successful connection
	log.Println("Connected to the database")

	return nil
}

func GetDB() *bun.DB {
	return db
}
