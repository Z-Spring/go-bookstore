package main

import (
	v1 "bookstore/router/api/v1"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	apiv1 := engine.Group("/api/v1")
	{
		apiv1.GET("/getbooks", v1.GetAllBooks)
		apiv1.POST("/create", v1.CreateBook)
		apiv1.GET("/login", v1.GetAllBooks)
		apiv1.POST("/login", v1.GetAllBooks)
		apiv1.GET("/register", v1.GetAllBooks)
		apiv1.POST("/register", v1.GetAllBooks)
		apiv1.GET("/getcart", v1.GetAllBooks)
	}
	engine.Run()

}
