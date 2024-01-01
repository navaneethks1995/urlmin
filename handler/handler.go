package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/navaneethks1995/urlmin/minifier"
	"github.com/navaneethks1995/urlmin/store"
)

// Request model definition
type CreateShortUrlRequest struct {
	OriginalUrl string `json:"original_url" binding:"required"`
	UserId      string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest CreateShortUrlRequest

	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	minUrl := minifier.GenreateMinUrl(creationRequest.OriginalUrl, creationRequest.UserId)
	store.SaveUrlMapping(creationRequest.OriginalUrl, minUrl, creationRequest.UserId)

	host := "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message":      "minified url created successfully",
		"minified_url": host + minUrl,
	})
}

func RedirectShortUrl(c *gin.Context) {
	minUrl := c.Param("min_url")
	originalUrl := store.GetOriginalUrl(minUrl)
	c.Redirect(302, originalUrl)
}
