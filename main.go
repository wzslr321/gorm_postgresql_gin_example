package main

import (
	"github.com/gin-gonic/gin"
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

	_ = r.Run()
}
