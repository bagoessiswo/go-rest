package controllers

import (
	"github.com/bagoessiswo/go-rest/initializers"
	"github.com/bagoessiswo/go-rest/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"data": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Return data
	c.JSON(200, gin.H{
		"data": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get id off url
	id := c.Param("id")
	// Get the post
	var post models.Post
	initializers.DB.First(&post, id)

	// Return data
	c.JSON(200, gin.H{
		"data": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get id off url
	id := c.Param("id")
	// Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	// Find the post
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Response
	c.JSON(200, gin.H{
		"data": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Delete post
	initializers.DB.Delete(&models.Post{}, id)

	// Response
	c.Status(200)
}
