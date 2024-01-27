package main

import (
	"delivery-backend/internals/controller"
	"delivery-backend/internals/usecase/postgres"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	postgres.ConnectToDB()
	r := gin.Default()
	r.GET("/restaurants", controller.GetRestaurants)
	r.POST("/restaurants", controller.PostRestaurant)
	r.Run()

	fmt.Print("hello world")
}
