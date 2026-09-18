package config

import (
	"api/model"
	"api/seeder"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbURL := os.Getenv("DATABASE_URL")

	var err error
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot connect to the database", err)
	}

	err = DB.AutoMigrate(&model.User{}, &model.Visit{})
	if err != nil {
		log.Fatal("failed to migrate", err)
	}

	seeder.SeedData(DB)
}
