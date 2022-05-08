package v1

import (
	"bookstore/data"
	"bookstore/model"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookInput struct {
	Name  string  `json:"name" form:"name" binding:"required,max=100"`
	Price float64 `json:"price" form:"price" binding:"required,max=300"`
	//ImagePath string  `json:"image_path"`
}

func GetAllBooks(c *gin.Context) {
	var book []model.Book
	data.DB.Find(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func GetBookById(c *gin.Context) {
	var book model.Book
	id := c.Param("id")
	if err := data.DB.Where("id=?", id).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "找不到这本书🍭"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// 暂时只更新价格

func UpdateBookById(c *gin.Context) {
	var (
		book      model.Book
		bookInput BookInput
	)

	id := c.Param("id")
	if err := data.DB.Where("id=?", id).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "找不到这本书🍭"})
		return
	}
	if err := c.ShouldBind(&bookInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(bookInput)
	data.DB.Model(&book).Update("price", bookInput.Price)
	//data.DB.Model(&book).Updates(bookInput)  todo:why?
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func CreateBook(c *gin.Context) {
	var input BookInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(input.Name)
	book := model.Book{Name: input.Name, Price: input.Price}
	data.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}
