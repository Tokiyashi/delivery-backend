package controller

import (
	"delivery-backend/app/internals/usecase/postgres"
	"delivery-backend/app/models"
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
	var body struct {
		Name string
	}

	c.Bind(&body)

	restaurant := models.Restaurant{Name: body.Name}

	result := postgres.DB.Create(&restaurant)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"result": restaurant,
	})
}

func PutRestaurant(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		ID     uint `gorm:"primarykey"`
		Name   string
		Rating int16
	}

	c.Bind(&body)

	var result models.Restaurant
	postgres.DB.First(&result, "id = ?", id)

	postgres.DB.Model(&result).Updates(models.Restaurant{Name: body.Name,
		Rating: body.Rating})

	c.JSON(200, gin.H{
		"result": result,
	})

}

func GetRestaurant(c *gin.Context) {
	id := c.Param("id")

	var result models.Restaurant
	postgres.DB.First(&result, "id = ?", id)

	c.JSON(200, gin.H{
		"result": result,
	})
}
