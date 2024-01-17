package db

import (
	"fmt"

	"github.com/pedrooctaviocruvinel/eulabs/src/config"
	"github.com/pedrooctaviocruvinel/eulabs/src/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	ProductRepository ProductRepository
}

func NewDatabase(i *gorm.DB) (database Database) {
	return Database{
		ProductRepository: *newProductRepository(i),
	}
}

func ConnectDB() (database *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)
	instance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err = migrateDatabase(instance); err != nil {
		return nil, err
	}

	return instance, nil
}

func migrateDatabase(i *gorm.DB) (err error) {
	if err := i.AutoMigrate(&types.Product{}); err != nil {
		return err
	}

	return nil
}
