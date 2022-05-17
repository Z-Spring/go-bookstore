package router

import (
	"bookstore/global"
	"bookstore/middleware"
	"bookstore/mylog"
	"bookstore/router/api"
	v1 "bookstore/router/api/v1"
	"github.com/gin-gonic/gin"
	"time"
)

func NewRouter() {
	logger, _ := mylog.NewZapProduction()
	engine := gin.New()
	gin.ForceConsoleColor()

	if global.ServerSetting.RunMode == "debug" {
		engine.Use(gin.Logger(), gin.Recovery())
	} else {
		engine.Use(middleware.Log(logger, time.RFC3339, true), gin.Recovery())
	}

	//engine.Use(middleware.LogToEsMiddleware())
	engine.Static("/static", "./static")
	{
		engine.POST("/auth", api.GetAuth)
		engine.POST("/register", v1.Register)
		engine.GET("/books", v1.GetAllBooks)
	}

	apiV1 := engine.Group("/api/v1")
	apiV1.Use(middleware.JWTMiddleware())
	{
		// book
		apiV1.GET("/books/:id", v1.GetBookById)
		apiV1.PATCH("/books/:id", v1.UpdateBookById)
		apiV1.POST("/create", v1.CreateBook)
		// cart
		apiV1.GET("/cart", v1.GetCart)
		apiV1.POST("/cart", v1.AddCart)
	}
	// auto open browser
	/*err := utils.OpenCmd("http://localhost:8080/books")
	if err != nil {
		mylog.Println(err)
		return
	}*/
	engine.Run(":8080")

}
