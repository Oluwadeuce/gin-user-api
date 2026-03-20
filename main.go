package main

import (
	"github.com/gin-gonic/gin"

	"gin-api-portfolio/database"
	"gin-api-portfolio/routes"
)

func main() {

	r := gin.Default()

	database.Database()

	routes.Routes(r)

	r.Run(":8080")
}