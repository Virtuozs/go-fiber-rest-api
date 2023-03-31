package config

import (
	"fmt"
	"golang-fiber-crud/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB(config *Config) *gorm.DB {
	sqlinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUsername, config.DBPassword, config.DBName)
	db, err := gorm.Open(postgres.Open(sqlinfo), &gorm.Config{})
	helper.ErrorPanic(err)
	fmt.Println("Connected to database")
	return db
}
