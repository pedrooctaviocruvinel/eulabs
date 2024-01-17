package db

import (
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
	dsn := "eulabsadmin:eulabspassword@tcp(127.0.0.1:3306)/eulabs?charset=utf8mb4&parseTime=True&loc=Local"
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
