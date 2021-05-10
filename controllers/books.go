package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/selithrarion/go-gin/models"
	"net/http"
)

func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func FindBook(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

type CreateBookDTO struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func CreateBook(c *gin.Context) {
	var body CreateBookDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: body.Title, Author: body.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

type UpdateBookDTO struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func UpdateBook(c *gin.Context) {
	//var book models.Book
	//if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found"})
	//	return
	//}
	//
	//var body UpdateBookDTO
	//if err := c.ShouldBindJSON(&body); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//models.DB.Model(&book).Updates(body)
	//c.JSON(http.StatusOK, gin.H{"data": book})

	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found!"})
		return
	}

	var body UpdateBookDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(&models.Book{Title: body.Title, Author: body.Author})

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found!"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
