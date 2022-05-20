package router

import (
	"bookstore/global"
	"bookstore/limiter"
	"bookstore/middleware"
	"bookstore/mylog"
	"bookstore/router/api"
	v1 "bookstore/router/api/v1"
	"context"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func NewRouter() {
	// rateLimiter init
	var m = limiter.MyLimiter{
		Limiter:      rate.NewLimiter(global.RateLimiterSetting.RateLimit, global.RateLimiterSetting.Buckets),
		LastGetToken: time.Now(),
		RoutePath:    global.RateLimiterSetting.RoutePath,
		RoutePathLimiter: limiter.RoutePathLimiter{
			//LimiterBuckets: make(map[string]int),
			LimiterBuckets: sync.Map{},
		},
	}

	logger, _ := mylog.NewZapProduction()
	engine := gin.New()
	gin.ForceConsoleColor()

	if global.ServerSetting.RunMode == "debug" {
		engine.Use(gin.Logger(), gin.Recovery())
	} else {
		engine.Use(middleware.Log(logger, time.RFC3339, true), gin.Recovery())
	}

	engine.Use(middleware.RateLimiter(&m))
	engine.Use(middleware.RouteTimeOut(1 * time.Minute))
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

	server := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server listen err:%s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("shutting down server...")

	ctx, channel := context.WithTimeout(context.Background(), 5*time.Second)
	defer channel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown error")
	}
	log.Fatal("server exiting...")

}
