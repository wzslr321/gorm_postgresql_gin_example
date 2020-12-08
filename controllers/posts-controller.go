package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/wzslr321/gorm_postgresql/models"
	"net/http"
)

var posts []models.Post

func populate() []models.Post{

	postsId := []int64{1,2}
	comments := []models.Comment{
		{
			Body: "Very cool post!",
		},
		{
			Body: "Keep it up!",
		},
	}
	author := &models.Author{Name: "Viktor", Posts:pq.Int64Array(postsId)}
	post := models.Post{
		ID: 1,
		Title: "My first post!",
		Description: "Some dummy description",
		Author: *author,
		Comments:comments,
	}


	posts = append(posts,post)

	return posts
}

func GetPosts(ctx *gin.Context){
	populate()
	ctx.JSON(http.StatusOK,gin.H{
		"posts":posts,
	})
}
