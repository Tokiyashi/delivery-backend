package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func getFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello backend world")
}

const (
	host     = "localhost"
	port     = 5432
	user     = "nikita"
	password = "5434"
	dbname   = "food"
)

func main() {

	ConnectToDB()

	fmt.Print("hello world")
	http.HandleFunc("/food", getFood)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectToDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	rows, err := db.Query(`SELECT "name" FROM "restaurants"`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var name string

		rows.Scan(&name)

		fmt.Println(name)

	}

}
