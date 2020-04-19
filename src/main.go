package main

import (
	"friend/config"
	"friend/controller"
)

func main() {
	db, _ := config.OpenDB()
	defer db.Close()

	controller.HandleRequest(db)
}