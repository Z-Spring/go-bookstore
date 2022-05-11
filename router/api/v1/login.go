package v1

import (
	"bookstore/data"
	"bookstore/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	var user model.User
	if err := data.DB.Where("name=? and password=?", name, password).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名或密码错误！"})
		return
	}
	// todo: how to handle redirect? 5.11
	// 外部重定向
	//c.Redirect(http.StatusMovedPermanently, "/auth")

	c.JSON(http.StatusOK, gin.H{"data": user})
	c.Request.URL.Path = "/auth"

}
