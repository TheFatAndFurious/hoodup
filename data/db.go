package data

import (
	"database/sql"
	"fmt"

	"goserver.com/internal/utils"
)



	func DbInitialize() (*sql.DB, error) { 
		
	relativePath := utils.GetEnv("DB_PATH", "../sqlite3.sqlite")
	databasePath := utils.GetAbsolutePath(relativePath)
	
	db, err := sql.Open("sqlite3", databasePath)

	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	fmt.Printf("Database up and running")

	return db, nil

	}
