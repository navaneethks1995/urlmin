package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/navaneethks1995/urlmin/handler"
	"github.com/navaneethks1995/urlmin/store"
)

func main() {
	r := gin.Default()

	r.POST("/minify", func(ctx *gin.Context) {
		handler.CreateShortUrl(ctx)
	})

	r.GET("/:min_url", func(ctx *gin.Context) {
		handler.RedirectShortUrl(ctx)
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Go URL minifier",
		})
	})
	store.InitializeStorage()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start server due to: %v", err))
	}

}
