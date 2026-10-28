package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase(config *Config) *gorm.DB {
	dns := fmt.Sprintf("DBhost=%s DBuser=%s DBpass=%s DBport=%s DBname=%s DBssl=%s",
		config.DB_HOST, config.DB_USER, config.DB_PASS, config.DB_PORT, config.DB_NAME, config.DBSSL_Mode)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db
}
