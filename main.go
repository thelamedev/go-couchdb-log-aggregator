package main

import (
	"log"

	"logmotor/pkg"
	"logmotor/pkg/database"
)

func main() {
	config, err := pkg.NewConfig("config.toml")
	if err != nil {
		log.Fatalln(err)
	}

	dbClient, err := database.NewDatabase(&config.CouchDB)
	defer dbClient.Close()

	dbInstance := dbClient.DB(config.CouchDB.Db)

	server := pkg.NewIngestionServer(config.App, dbInstance)
	server.Prepare()

	server.Listen()
}
