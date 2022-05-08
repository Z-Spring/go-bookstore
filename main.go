package main

import (
	"bookstore/router"
	"bookstore/settings"
	"log"
)

func init() {
	if err := settings.GetSettings(); err != nil {
		log.Println(err)
	}
}
func main() {
	router.NewRouter()
}
