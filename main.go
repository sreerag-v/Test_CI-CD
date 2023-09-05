package main

import (
	"Test/Test-Crud/database"
	"Test/Test-Crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	database.InitDB()

	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	routes.Router(router)
	router.Run("localhost:8080")
}
