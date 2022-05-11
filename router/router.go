package router

import (
	"bookstore/auth"
	"bookstore/middleware"
	v1 "bookstore/router/api/v1"

	"github.com/gin-gonic/gin"
)

func NewRouter() {
	engine := gin.Default()
	engine.Static("/static", "./static")

	//engine.POST("/login", v1.Login)
	engine.POST("/auth", auth.GetAuth)

	apiv1 := engine.Group("/api/v1")
	apiv1.Use(middleware.JWTMiddleware())
	{
		// book
		apiv1.GET("/books", v1.GetAllBooks)
		apiv1.GET("/books/:id", v1.GetBookById)
		apiv1.PATCH("/books/:id", v1.UpdateBookById)
		apiv1.POST("/create", v1.CreateBook)
		// login  register
		//apiv1.POST("/login", v1.Login)
		apiv1.POST("/register", v1.Register)
		// cart
		apiv1.GET("/cart", v1.GetCart)
		apiv1.POST("/cart", v1.AddCart)
	}
	engine.Run(":8080")

}
