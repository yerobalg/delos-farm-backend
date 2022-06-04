package bootstrap

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"delos-farm-backend/domains"
	"github.com/go-redis/redis/v8"
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
	);

	return db, err
}

func InitRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to redis!")

	return client, nil
}
