package data

import (
	"database/sql"
	"fmt"

	"goserver.com/utils"
)



	func DbInitialize() (*sql.DB, error) { 
		
	databasePath := utils.GetEnv("DATABASE_PATH", "../../db/db.sqlite3")
	
	db, err := sql.Open("sqlite3", databasePath)

	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	fmt.Printf("Database up and running")


	return db, nil

}
