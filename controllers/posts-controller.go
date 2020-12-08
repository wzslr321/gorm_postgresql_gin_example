package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	util "github.com/wzslr321/gorm_postgresql/db"
	"github.com/wzslr321/gorm_postgresql/models"
	"net/http"
)

// This function does not save anything to the database,
// It just uses gorm model to create Dummy Data.
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

	var posts []models.Post

	posts = append(posts,post)

	return posts
}

// Fetches all posts created with populate, and displays it as JSON.
func GetPosts(ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{
		"posts":populate(),
	})
}

// Saves new post with with only Title and Description to the database.
func AddPostsToDatabase(ctx *gin.Context){
	util.Db.Create(&models.Post{Title: "Trial post added inline", Description: "Some dummy description once again"})
	ctx.JSON(
		http.StatusOK,gin.H{
			"status":"posted",
		},
	)
}
