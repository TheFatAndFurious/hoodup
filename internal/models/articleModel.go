package models

import (
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type Article struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Exerpt string `json:"exerpt"`
	Content string `json:"content"`
	Published bool `json:"published"`
	PublishedAt time.Time `json:"published_at"`
	Author int `json:"author"`
    HasBeenPublished bool `json:"has_been_published"`
    Illustration string `json:"illustration"`
}  

func FormToArticle(form url.Values, existingArticle *Article) (*Article, error) {
    authorID, err := strconv.Atoi(form.Get("author"))
    if err != nil {
        return nil, errors.New("invalid author ID")
    }

    idStr := form.Get("id")
    id := 0
    if idStr!= "" {
        id, err = strconv.Atoi(idStr)
        if err!= nil {
            return nil, errors.New("invalid article ID")
        }
    }

	article := &Article{
        Id: id,
        Title: form.Get("title"),
        Exerpt: form.Get("exerpt"),
        Content: form.Get("content"),
        Published: form.Get("published") == "true",
        Author: authorID,
        Illustration: form.Get("illustration"),
    }
        if article.Published && !article.HasBeenPublished {
            article.PublishedAt = time.Now()
            article.HasBeenPublished = true
        } else if existingArticle != nil {
            article.PublishedAt = existingArticle.PublishedAt
        }
        return article, nil
    } 

func PersistArticle(db *sql.DB, article *Article) error {
	statement, err := db.Prepare("INSERT INTO articles (title, exerpt, content, published, published_at, author, has_been_published, illustration) VALUES (?,?,?,?,?,?,?,?)")
    if err != nil {
        return fmt.Errorf("error preparing staement: %w", err)
    }
    defer statement.Close()

    _, err = statement.Exec(article.Title, article.Exerpt, article.Content, article.Published, article.PublishedAt, article.Author, article.HasBeenPublished, article.Illustration)
    if err!= nil {
        return fmt.Errorf("error executing the statement: %w", err)
    }

    return nil
}

func UpdateArticle(db *sql.DB, article *Article) error {
    statement, err := db.Prepare("UPDATE articles SET title =?, exerpt =?, content =?, published=?, published_at =?, author =?, has_been_published =?, illustration =? WHERE id =?")
    if err != nil {
        return fmt.Errorf("error preparing staement: %w", err)
    }
    defer statement.Close()
    _, err = statement.Exec(article.Title, article.Exerpt, article.Content, article.Published, article.PublishedAt, article.Author, article.HasBeenPublished, article.Illustration, article.Id)
    if err != nil {
        return fmt.Errorf("error executing the statement: %w", err)
    }
    return nil
}
func FetchSingleArticle(db *sql.DB, articleId int) (Article, error) {
    query := "SELECT id, title, content, exerpt, published, published_at, has_been_published, illustration FROM articles WHERE id =?"
    var a Article
    err := db.QueryRow(query, articleId).Scan(&a.Id, &a.Title, &a.Content, &a.Exerpt, &a.Published, &a.PublishedAt, &a.HasBeenPublished, &a.Illustration)
    if err != nil {
        return Article{}, err
    }
    return a, nil
}

func FetchArticleByTag(db *sql.DB, tag string) ([]Article, error) {
    var articles []Article
    query := "SELECT id, title, content, exerpt, published, published_at, has_been_published, illustration FROM articles LEFT JOIN articles_tags on articles.id = articles_tags.article WHERE tag = (?)"
    rows, err := db.Query(query, tag)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
        var a Article
        err := rows.Scan(&a.Id, &a.Title, &a.Content, &a.Exerpt, &a.Published, &a.PublishedAt, &a.HasBeenPublished, &a.Illustration)
        if err != nil {
            return nil, err
        }
        articles = append(articles, a)
    }
    return articles, nil
}

func FetchAllArticles(db *sql.DB) ([]Article, error) {
    var articles []Article
    query := "SELECT id, title, content, exerpt, published, published_at, has_been_published, illustration FROM articles"
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
        var a Article
        err := rows.Scan(&a.Id, &a.Title, &a.Content, &a.Exerpt, &a.Published, &a.PublishedAt, &a.HasBeenPublished, &a.Illustration)
        if err != nil {
            return nil, err
        }
        articles = append(articles, a)
    }
    return articles, nil
}