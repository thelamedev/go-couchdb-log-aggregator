package database

import (
	"context"
	"log"
	"strings"

	"logmotor/pkg"

	kivik "github.com/go-kivik/kivik/v4"
	_ "github.com/go-kivik/kivik/v4/couchdb"
)

const (
	COUCHDB_USER     = "admin"
	COUCHDB_PASSWORD = "admin1234"
)

func NewDatabase(config *pkg.CouchDBConfig) (*kivik.Client, error) {
	client, err := kivik.New("couch", config.Url)
	if err != nil {
		log.Fatalln(err)
	}

	// Create the database if it does not exist
	if err := client.CreateDB(context.Background(), config.Db); err != nil {
		if !strings.Contains(err.Error(), "Precondition Failed") {
			log.Fatalln(err)
		}
	}

	return client, nil
}
