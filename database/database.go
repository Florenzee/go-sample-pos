package database

import (
	"go-sample-post/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:amartha@tcp(localhost:3306)/schema_go_sample_post?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	database.AutoMigrate(&models.Product{})
	database.AutoMigrate(&models.Customer{})
	database.AutoMigrate(&models.Employee{})
	database.AutoMigrate(&models.Receipt{})
	DB = database
}
