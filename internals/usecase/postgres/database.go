package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "nikita"
	password = "5434"
	dbname   = "food"
)

var DB *gorm.DB

func ConnectToDB() {

	dsn := "host=localhost user=nikita password=5434 dbname=food port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB = db

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
