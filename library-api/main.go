package main

import (
	"fmt"
	"library-api/db"
	"library-api/server"
)

// @title Library API - Swagger Documentation
// @version 1.0.0
// @description Library API for borrowing book
// @termsOfService http://swagger.io/terms/

// @contact.name Gizwanda
// @contact.url http://www.swagger.io/support
// @contact.email adigiz@yahoo.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	err := db.Connect()
	if err != nil {
		fmt.Println("Failed to connect DB", err)
	}
	server.Init()
}
