package main

import (
	"log"

	"github.com/pedrooctaviocruvinel/eulabs/src/api"
	"github.com/pedrooctaviocruvinel/eulabs/src/config"
	"github.com/pedrooctaviocruvinel/eulabs/src/db"
	"github.com/pedrooctaviocruvinel/eulabs/src/services/product_services"
)

func main() {
	err := config.LoadEnvironmentVariables("./.env")
	if err != nil {
		log.Fatal(err)
	}

	instance, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	database := db.NewDatabase(instance)

	productServices := product_services.NewProductServices(database.ProductRepository)

	api := api.NewApi(productServices)

	api.Run()
}
