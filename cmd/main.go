package main

import (
	"ecommerce-boiler/cmd/server"
	//_ "boilerplate_golang/docs"
	"ecommerce-boiler/pkg/database"
	"log"
)

func main() {
	var DBConnection = database.NewConnection()
	err := server.Run(DBConnection)
	if err != nil {
		log.Fatal(err)
		return
	}
	server.Injection()
}
