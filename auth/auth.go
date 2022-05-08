package auth

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type RequestAuth struct {
	Secret string `json:"secret,omitempty"`
}

func GetAuth(c *gin.Context) {
	var requestAuth RequestAuth
	if err := c.ShouldBind(&requestAuth); err != nil {
		log.Println(err)
		return
	}
	token, err := GenerateToken(requestAuth.Secret)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"jwt": token})

}
