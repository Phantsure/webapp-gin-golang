// handlers.article.go

package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// Call the render function with name of the template to render
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles,
	}, "index.html")
}

func getArticle(c *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		if article, err := getArticleByID(articleID); err == nil {
			// Call the render function with name of the template to render
			render(c, gin.H{
				"title":   article.Title,
				"payload": article,
			}, "article.html")
		} else {
			// If the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// If an invalid article ID is specified in the URL, abort with error
		c.AbortWithError(http.StatusNotFound, err)
	}
}
