package v1

import (
	"bookstore/data"
	"bookstore/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
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
	if err := c.ShouldBindJSON(&bookInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	/*data.DB.Model(&book).Select("Name", "Price", "ImagePath").Updates(
	model.Book{
		Name:      bookInput.Name,
		Price:     bookInput.Price,
		ImagePath: bookInput.ImagePath,
	})*/
	log.Println(bookInput)
	data.DB.Model(&book).Update("price", bookInput.Price)
	//data.DB.Model(&book).Updates(bookInput)  todo:why?
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func CreateBook(c *gin.Context) {
	var input BookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book := model.Book{Name: input.Name, Price: input.Price}
	data.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}
