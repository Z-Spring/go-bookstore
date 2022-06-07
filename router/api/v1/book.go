package v1

import (
	"bookstore/data"
	"bookstore/model"
	"bookstore/utils"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type BookInput struct {
	Name  string  `json:"name" form:"name" binding:"required,max=100"`
	Price float64 `json:"price" form:"price" binding:"required,max=300"`
	//ImagePath string  `json:"image_path"`
}

var ctx = context.Background()
var res = utils.NewRedis()

func GetAllBooks(c *gin.Context) {
	var book []model.Book
	if result, err := res.Get(ctx, "books").Result(); err == nil {
		err := json.Unmarshal([]byte(result), &book)
		if err != nil {
			log.Printf("unmarshal error,%s", err)
		}
		log.Printf("got books from redis: %v", book)
		c.JSON(http.StatusOK, gin.H{"data": book})
		return
	}
	log.Println("can't find books in redis! search books from mysql")
	data.DB.Find(&book)

	data, err2 := json.Marshal(book)
	if err2 != nil {
		log.Println("marshal error", err2)
	}
	err := res.Set(ctx, "books", data, 10*time.Minute).Err()
	if err != nil {
		log.Printf("redis set error:%s\n", err)
	}
	resBook, err := res.Get(ctx, "books").Result()
	if err != nil {
		log.Printf("redis get book error,%s", err)
	}
	err = json.Unmarshal([]byte(resBook), &book)
	if err != nil {
		log.Printf("unmarshal error,%s", err)
	}

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
	log.Println(input.Name)
	book := model.Book{Name: input.Name, Price: input.Price}
	data.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}
