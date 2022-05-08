package v1

import (
	"bookstore/data"
	"bookstore/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Login(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	log.Println(name)
	var user model.User
	if err := data.DB.Where("name=? and password=?", name, password).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名或密码错误！"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
