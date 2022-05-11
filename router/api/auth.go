package api

import (
	"bookstore/auth"
	"bookstore/data"
	"bookstore/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserInfo struct {
	Name     string `form:"name" binding:"required,min=3,max=10"`
	Password string `form:"password" binding:"required,min=1,max=50"`
}

func GetAuth(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	var u UserInfo
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user model.User
	if err := data.DB.Where("name=? and password=?", u.Name, u.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名或密码错误！"})
		return
	}

	token, err := auth.GenerateToken(name, password)
	if err != nil {
		log.Println(err)
	}
	//c.Header("Authorization", "Bearer "+token)
	cookie := http.Cookie{
		Name:  "Authorization",
		Value: token,
	}
	http.SetCookie(c.Writer, &cookie)
	//c.JSON(http.StatusOK, gin.H{"jwt": token})

}
