package main

import (
	"main/config"
	"main/controller"
)

func main() {
	db, _ := config.OpenDB()
	defer db.Close()

	controller.HandleRequest(db)
}