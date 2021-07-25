package main

import (
	"log"

	"example.com/mymodule/postgres/src/models"
)

func main() {
	db, err := models.InitDB()
	if err != nil {
		log.Println(db)
	}
}
