package database

import (
	"log"
	"os"

	"github.com/go-pg/pg/v10"
)

var Db *pg.DB

func Connect() {
	opts := &pg.Options{
		User : "shouryagautam",
		Password: "shourya",
		Addr : "localhost:5432",
		Database: "tuts",
	}

	Db = pg.Connect(opts)

	if Db == nil {
		log.Printf("Failed to connect to database.\n")
		os.Exit(100)
	}

}
