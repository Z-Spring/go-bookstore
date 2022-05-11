package v1

import (
	"bookstore/data"
	"bookstore/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type RegisterInput struct {
	Name     string `form:"name" binding:"required,min=1,max=20"`
	Password string `form:"password" binding:"required,min=2,max=50"`
}

func Register(c *gin.Context) {
	uid := "u-" + strconv.FormatInt(time.Now().Unix(), 10)
	/*	name := c.PostForm("name")
		password := c.PostForm("password")*/
	var u RegisterInput
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := model.User{
		Name:     u.Name,
		Password: u.Password,
	}
	data.DB.Create(&user)
	data.DB.Model(&user).Update("uid", uid)

	c.JSON(http.StatusOK, gin.H{"data": u})
}
