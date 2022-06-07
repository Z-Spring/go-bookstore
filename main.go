package main

import (
	"bookstore/router"
	"bookstore/settings"
	"log"
)

var (
	port   string
	mode   string
	config string
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
	if err := settings.GetContextTimeout(); err != nil {
		log.Println(err)
	}
}

func main() {
	router.NewRouter()
}
