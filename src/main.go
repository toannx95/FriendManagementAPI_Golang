package main

import (
	"config"
	"controller"
)

func main() {
	db, _ := config.OpenDB()
	defer db.Close()

	controller.HandleRequest(db)
}