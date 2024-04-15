package data

import (
	"database/sql"
	"fmt"
)



	func DbInitialize() (*sql.DB, error) {    
	
	db, err := sql.Open("sqlite3", "sqlite3.sqlite")

	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	fmt.Printf("Database up and running")

	return db, nil

}
