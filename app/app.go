package main

import (
	"fmt"
	"log"
	"mekar/config"
)

func main() {
	db, err := config.Connect()

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected Successfully!")
	}

	router := config.CreateRouter()
	config.InitRouter(router, db)
	config.RunServer(router)
}
