package main

import (
	"log"

	"FuegoDeQuasar/cmd/config/db"
	server "FuegoDeQuasar/internal/infrastructure/api/starshipCommsResolver"
)

func main() {
	dbInstance, err := db.NewPostgreSQLDB()
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	server.RunServer(dbInstance)
}
