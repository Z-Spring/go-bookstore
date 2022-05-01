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
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

/*func RegisterUsers(c *gin.Context) {
	var registerInput RegisterInput
	uid := "u-" + strconv.FormatInt(time.Now().Unix(), 10)
	if err := c.ShouldBindJSON(&registerInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := model.User{
		Name:     registerInput.Name,
		Password: registerInput.Password,
	}
	fmt.Println(registerInput.Name)
	data.DB.Create(&user)
	data.DB.Model(&user).Update("uid", uid)
	c.JSON(http.StatusOK, gin.H{"data": user})
}*/

// TODO 缺少验证环节

func RegisterUsers(c *gin.Context) {
	uid := "u-" + strconv.FormatInt(time.Now().Unix(), 10)
	name := c.PostForm("name")
	password := c.PostForm("password")
	user := model.User{
		Name:     name,
		Password: password,
	}
	data.DB.Create(&user)
	data.DB.Model(&user).Update("uid", uid)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
