package controller

import (
	"delivery-backend/internals/usecase/postgres"
	"delivery-backend/models"
	"github.com/gin-gonic/gin"
)

func GetRestaurants(c *gin.Context) {
	var restaurants []models.Restaurant

	postgres.DB.Find(&restaurants)

	c.JSON(200, gin.H{
		"restaurants": restaurants,
	})
}

func PostRestaurant(c *gin.Context) {
	restaurant := models.Restaurant{Name: "Вкусочка"}

	result := postgres.DB.Create(&restaurant)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"result": restaurant,
	})
}
