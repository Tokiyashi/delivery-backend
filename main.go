package main

import (
	"delivery-backend/app/internals/controller"
	"delivery-backend/app/internals/usecase/postgres"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	postgres.ConnectToDB()

	r := gin.Default()
	r.GET("/restaurants", controller.GetRestaurants)
	r.POST("/restaurants", controller.PostRestaurant)
	r.PUT("/restaurants/:id", controller.PutRestaurant)
	r.GET("/restaurants/:id", controller.GetRestaurant)

	fmt.Println("Server has started successfully! 🚀")

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
