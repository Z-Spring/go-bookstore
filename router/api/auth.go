package api

import (
	"bookstore/auth"
	"bookstore/data"
	"bookstore/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type UserInfo struct {
	Name     string `form:"name" binding:"required,min=3,max=10"`
	Password string `form:"password" binding:"required,min=1,max=50"`
}

func GetAuth(c *gin.Context) {
	//name := c.PostForm("name")
	/*	name := c.DefaultPostForm("name", "qqq")
		password := c.PostForm("password")
	*/
	var u UserInfo
	if err := c.ShouldBind(&u); err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(u.Name)

	var user model.User
	if err := data.DB.Where("name=? and password=?", u.Name, u.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名或密码错误！"})
		return
	}
	// todo 这里可能会出错   orm错误
	//token, err := auth.GenerateToken(name, password)
	token, err := auth.GenerateToken(u.Name, u.Password)
	if err != nil {
		log.Println(err)
	}
	var expires = time.Now().Add(24 * 60 * time.Minute)
	cookie := http.Cookie{
		Name:    "Authorization",
		Value:   token,
		Expires: expires,
		//HttpOnly: true,
	}
	/*userName := http.Cookie{
		Name:    "name",
		Value:   u.Name,
		Expires: expires,
		//HttpOnly: true,
	}*/
	http.SetCookie(c.Writer, &cookie)
	//http.SetCookie(c.Writer, &userName)
	c.JSON(http.StatusOK, gin.H{"data": u.Name})

}
