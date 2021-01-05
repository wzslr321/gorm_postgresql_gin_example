package controllers

import (
	"github.com/gin-gonic/gin"
	util "github.com/wzslr321/gorm_postgresql/db"
	"github.com/wzslr321/gorm_postgresql/models"
	"net/http"
)

// This function does not save anything to the database,
// It just uses gorm model to create Dummy Data.
func populate() []models.Product{

	post := models.Product{
		Title: "My first post!",
		Description: "Some dummy description",
		Price: 19.99,
		ImageUrl: "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse1.mm.bing.net%2Fth%3Fid%3DOIP.RaHhnyI6Vqnxb5TCF16F7QHaGK%26pid%3DApi&f=1",
	}

	var products []models.Product

	products = append(products,post)

	return products
}

// Fetches all posts created with populate, and displays it as JSON.
func GetPosts(ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{
		"posts":populate(),
	})
}

// Saves new post with with only Title and Description to the database.
func AddPostsToDatabase(ctx *gin.Context){
	util.Db.Create(&models.Product{Title: "Trial post added inline", Description: "Some dummy description once again", Price: 10.99, ImageUrl:  "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse1.mm.bing.net%2Fth%3Fid%3DOIP.RaHhnyI6Vqnxb5TCF16F7QHaGK%26pid%3DApi&f=1"})
	ctx.JSON(
		http.StatusOK,gin.H{
			"status":"posted",
		},
	)
}
