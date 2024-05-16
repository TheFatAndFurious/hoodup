package integration

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"goserver.com/internal/handlers"
	"goserver.com/internal/models"
)

// Set up a test database and insert a user for login testing


func setupTestDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	// Create articles table
	createTableSQL := `CREATE TABLE articles (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"title" TEXT,
		"content" TEXT,
		"exerpt" TEXT,
		"published" BOOLEAN,
		"published_at" DATETIME,
		"author" INTEGER,
		"has_been_published" BOOLEAN
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}

	// Insert sample articles
	insertArticleSQL := `INSERT INTO articles (title, content, exerpt, published, author, has_been_published) VALUES (?, ?, ?, ?, ?, ?), (?, ?, ?, ?, ?, ?)`
	_, err = db.Exec(insertArticleSQL,
		"First Article", "This is the first article.", "Excerpt", false, 1, false,
		"Second Article", "This is the second article.", "Excerpt", false, 1, false)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestGetArticles(t *testing.T) {
	// Step 1: Setup the test database
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to setup test database: %v", err)
	}
	defer db.Close()

	// Step 2: Initialize the router and define the route
	router := mux.NewRouter()
	router.HandleFunc("/articles", handlers.DisplayAllArticlesHandler(db)).Methods("GET")

	// Step 3: Simulate an HTTP request to the articles endpoint
	fmt.Print("++++++++++++++++++++++++++++++++")
	req, err := http.NewRequest("GET", "/articles", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)


	// Step 4: Verify the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var articles []models.Article
	if err := json.NewDecoder(rr.Body).Decode(&articles); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	if len(articles) != 2 {
		t.Errorf("Handler returned wrong number of articles: got %v want %v", len(articles), 2)
	}

	expectedTitles := []string{"First Article", "Second Article"}
	for i, article := range articles {
		if article.Title != expectedTitles[i] {
			t.Errorf("Handler returned wrong article title: got %v want %v", article.Title, expectedTitles[i])
		}
	}
}
