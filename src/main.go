package main

import (
	"friend/config"
	"friend/router"
)

// @title Friend Management API
// @version 1.0
// @description This is a sample service for managing Friend Management
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8081
// @BasePath /
func main() {
	db, _ := config.OpenDB()
	defer db.Close()

	router.HandleRequest(db)
}