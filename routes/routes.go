package routes

import (
	"gin-api-portfolio/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	r.GET("/users", handlers.GetUsers)

	r.GET("/users/:id", handlers.GetUserByID)

	r.POST("/users", handlers.CreateUser)

	r.DELETE("/users/:id", handlers.DeleteUser)

}
