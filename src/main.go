package main

import (
	"friend/config"
	"friend/router"
)

func main() {
	db, _ := config.OpenDB()
	defer db.Close()

	router.HandleRequest(db)
}