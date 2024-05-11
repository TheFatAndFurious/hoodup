package data

import (
	"database/sql"
	"fmt"
)

func InitializeSchema(db *sql.DB) error {
	statements := []string{
		"CREATE TABLE IF NOT EXISTS roles (id INTEGER PRIMARY KEY AUTOINCREMENT, role TEXT)",
		"CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, email TEXT, password TEXT, role TEXT, FOREIGN KEY (role) REFERENCES roles(role))",
		"CREATE TABLE IF NOT EXISTS tags (id INTEGER PRIMARY KEY AUTOINCREMENT, tag TEXT)",
		"CREATE TABLE IF NOT EXISTS articles (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT, content TEXT, exerpt TEXT, published BOOLEAN, published_at DATETIME, author INTEGER, has_been_published BOOLEAN DEFAULT FALSE, illustration TEXT DEFAULT '', FOREIGN KEY (author) REFERENCES users(id))",
		"CREATE TABLE IF NOT EXISTS comments (id INTEGER PRIMARY KEY AUTOINCREMENT, content TEXT, published BOOLEAN, published_at DATETIME, author TEXT, article INTEGER, FOREIGN KEY (article) REFERENCES articles(id))",
		"CREATE TABLE IF NOT EXISTS articles_tags (id INTEGER PRIMARY KEY AUTOINCREMENT, article INTEGER, tag INTEGER, FOREIGN KEY (tag) REFERENCES tags(id), FOREIGN KEY (article) REFERENCES articles(id))",
		
	}

	for _, statement := range statements {
		_, err := db.Exec(statement)
		if err != nil {
			return fmt.Errorf("failed to create table %s: %w", statement, err)
		}

	}
	return nil
}