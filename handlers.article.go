// handlers.article.go

package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-gin-app/monitor"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()
	monitor.ShowIndexPage.Inc()
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")
}

func showArticleCreationPage(c *gin.Context) {
	monitor.ShowArticleCreationPage.Inc()
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Create New Article"}, "create-article.html")
}

func getArticle(c *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		s := c.PostForm("article_id")
		// Check if the article exists
		if article, err := getArticleByID(articleID); err == nil {
			// Call the render function with the title, article and the name of the
			// template
			monitor.ShowArticlePage.WithLabelValues(s, "").Inc()
			render(c, gin.H{
				"title":   article.Title,
				"payload": article}, "article.html")

		} else {
			// If the article is not found, abort with an error
			monitor.ShowArticlePage.WithLabelValues(s, "article_not_found").Inc()
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		monitor.ShowArticlePage.WithLabelValues("invalid", fmt.Sprintf("invalid article URL: %v",
			c.Param("article_id"))).Inc()
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func createArticle(c *gin.Context) {
	// Obtain the POSTed title and content values
	title := c.PostForm("title")
	content := c.PostForm("content")

	if a, err := createNewArticle(title, content); err == nil {
		// If the article is created successfully, show success message
		monitor.PerformArticleCreation.WithLabelValues("").Inc()
		render(c, gin.H{
			"title":   "Submission Successful",
			"payload": a}, "submission-successful.html")
	} else {
		monitor.PerformArticleCreation.WithLabelValues("error").Inc()
		// if there was an error while creating the article, abort with an error
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
