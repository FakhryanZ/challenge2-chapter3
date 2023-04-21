package main

import (
	"golang-learning-path/go-middleware/database"
	"golang-learning-path/go-middleware/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
