package models

import (
	"context"
	"golang-backend/db"
)

type AuthorRevenueRepository struct{}

type AuthorRevenue struct {
	AuthorID     int64  `json:"author_id"`
	TotalRevenue int64  `json:"total_revenue"`
	Username     string `json:"username"`
}

func (c AuthorRevenueRepository) GetRevenue(ctx context.Context, id int64) ([]AuthorRevenue, error) {
	query := `WITH authorRevenueTable AS (
        SELECT 
            b.author_id,
            SUM(bp.total_price) AS total_revenue
        FROM 
            books b 
        JOIN 
            book_purchases bp ON b.book_id = bp.book_id
        GROUP BY 
            b.author_id
    )
    SELECT 
        art.author_id,
        art.total_revenue,
        u.username
    FROM 
        authorRevenueTable art
    JOIN 
        users u ON u.id = art.author_id
    WHERE 
        art.author_id = ?`

	rows, err := db.GetDB().QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Define a slice to store the results
	var results []AuthorRevenue

	// Iterate through the rows and extract the data
	for rows.Next() {
		var revenue AuthorRevenue
		if err := rows.Scan(&revenue.AuthorID, &revenue.TotalRevenue, &revenue.Username); err != nil {
			return nil, err
		}
		// Append the result to the results slice
		results = append(results, revenue)
	}

	// Check for any errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Return the results
	return results, nil
}
