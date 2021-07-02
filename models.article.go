// models.article.go

package main

import (
	"errors"

	"github.com/go-gin-app/config"
)

// Return a list of all the articles
func getAllArticles() []config.Article {
	// return articleList
	return config.ArticleList
}

// Fetch an article based on the ID supplied
func getArticleByID(id int) (*config.Article, error) {
	for _, a := range config.ArticleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}

// Create a new article with the title and content provided
func createNewArticle(title, content string) (*config.Article, error) {
	// Set the ID of a new article to one more than the number of articles
	a := config.Article{ID: len(config.ArticleList) + 1, Title: title, Content: content}

	// Add the article to the list of articles
	config.ArticleList = append(config.ArticleList, a)

	return &a, nil
}
