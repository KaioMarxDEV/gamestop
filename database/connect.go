package database

import (
	"fmt"
	"log"

	"github.com/KaioMarxDEV/gamestop/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() error {
	var err error

	dsn := "host=db port=5432 user=gorm password=gorm dbname=gorm sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("database connection failed")
		return err
	}

	fmt.Println("connection opened to database")
	DB.AutoMigrate(&models.User{})
	fmt.Println("migration executed successfully")
	return nil
}
