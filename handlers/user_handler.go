package handlers

import (
	"gin-api-portfolio/database"
	"gin-api-portfolio/models"

	"github.com/gin-gonic/gin"
)

var users = []models.User{
	{ID: 1, Username: "Gbemisola", Email: "gbemisola@gmail.com"},
	{ID: 2, Username: "Barry", Email: "barry@gmail.com"},
	{ID: 3, Username: "Ameera", Email: "ameera@gmail.com"},
}

func GetUsers(c *gin.Context) {

	var users []models.User

	database.DB.Find(&users)

	c.JSON(200, users)
}

func GetUserByID(c *gin.Context) {

	id := c.Param("id")

	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {

		c.JSON(404, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(200, user)
}

func CreateUser(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"error": "Wrong input",
		})
		return
	}

	database.DB.Create(&user)

	c.JSON(201, user)
}

func DeleteUser(c *gin.Context) {

	id := c.Param("id")

	database.DB.Delete(&models.User{}, id)

	c.JSON(200, gin.H{
		"message": "User is deleted",
	})
}
