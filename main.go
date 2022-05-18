package main

import (
	"bookstore/router"
	"bookstore/settings"
	"log"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func init() {
	if err := settings.GetJwtSettings(); err != nil {
		log.Println(err)
	}
	if err := settings.GetServerSettings(); err != nil {
		log.Println(err)
	}
	if err := settings.GetRateLimiterSettings(); err != nil {
		log.Println(err)
	}
}

func main() {
	router.NewRouter()
}
