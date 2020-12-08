package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wzslr321/gorm_postgresql/controllers"
	"net/http"
)

func main() {

	r := gin.Default()

	r.GET("/api/index", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK, gin.H{
				"message": "Welcome on index page!",
			},
		)
	})

	r.GET("/api/posts", controllers.GetPosts)

	_ = r.Run()
}
