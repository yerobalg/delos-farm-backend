package bootstrap

import (
	"delos-farm-backend/domains"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func InitPostgres() (*gorm.DB, error) {
	dataSourceName := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DBNAME"),
	)

	db, error := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})

	if error != nil {
		return nil, error
	}

	fmt.Println("Successfully connected to database!")

	err := db.AutoMigrate(
		&domains.Farms{},
		&domains.Ponds{},
		&domains.Stats{},
	)

	return db, err
}
