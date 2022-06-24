package main

import (
	"flag"
	"log"

	"gitlab.com/jacob-ernst/mets/pkg/models"
)

func main() {
	var dsn string
	flag.StringVar(&dsn, "dsn", "data/dev.db", "dsn for the DB to migrate")
	flag.Parse()
	_, err := models.OpenDB(dsn)
	if err != nil {
		log.Fatalln("could not migrate dsn \"", dsn, "\":", err)
	}
}
