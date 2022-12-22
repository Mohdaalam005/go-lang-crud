package main

import (
	"github.com/joho/godotenv"
	"github.com/mohdaalam005/crud-with-postgres/controller"
	"log"
)

func main() {

	err := godotenv.Load("properties.env")
	CheckNilErr(err)
	log.Println("application started successfully")
	controller.Route()
}

func CheckNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
