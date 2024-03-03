package controllers

import (
	"example/go-api-gin-gorm/initializers"
	"example/go-api-gin-gorm/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	//Get data of request body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	log.Println(body)

	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetAllPosts(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPostById(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdateById(c *gin.Context) {

	id := c.Param("id")

	//Get data of request body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	log.Println(body)

	var post models.Post

	initializers.DB.First(&post, id)

	postToBeUpdated := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Model(&post).Updates(postToBeUpdated)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDeleteById(c *gin.Context) {

	id := c.Param("id")

	if id == "" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Missing Id query parameter"})
		return
	}
	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Delete(&models.Post{}, id)

	c.Status(200)

}
