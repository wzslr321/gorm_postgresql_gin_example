package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	util "github.com/wzslr321/gorm_postgresql/db"
	"github.com/wzslr321/gorm_postgresql/models"
)

func addPost(ctx *gin.Context){
	postsId := []int64{1,2}
	author := &models.Author{Name: "Viktor", Posts:pq.Int64Array{postsId}}
	util.Db.Create(&models.Post{Title: "My first post!", Description: "Some dummy description"})
}
