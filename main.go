package main

import (
	"fmt"
	"log"

	"github.com/pedrooctaviocruvinel/eulabs/src/db"
)

func main() {
	instance, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	database := db.NewDatabase(instance)

	fmt.Println(database)
}
