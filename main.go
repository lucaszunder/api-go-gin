package main

import (
	"github.com/lucaszunder/api-go-gin/database"
	"github.com/lucaszunder/api-go-gin/routes"
)

func main() {
	database.ConnectDatabase()
	routes.HandleRequests()

}
