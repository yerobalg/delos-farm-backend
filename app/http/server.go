package main

import (
	"delos-farm-backend/middlewares"
	"github.com/gin-gonic/gin"
	"os"
	"github.com/joho/godotenv"
	"fmt"
	"delos-farm-backend/bootstrap"
	"delos-farm-backend/farms/repository"
	"delos-farm-backend/farms/service"
	_farmsHandler "delos-farm-backend/farms/handler"
)

func main() {

	//init env variable
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("failed to load env from local file")
	}

	//init database
	db, err := bootstrap.InitDB()
	if err != nil {
		panic(err)
	}
	
	//init router
	engine := gin.Default()
	engine.Use(middlewares.CorsMiddleware())
	router := engine.Group("/api/v1")

	//init farms
	farmsRepository := repository.NewFarmsRepository(db)
	farmsService := service.NewFarmsService(farmsRepository)
	_farmsHandler.NewFarmsHandler(router, farmsService)

	//run server
	engine.Run(":" + os.Getenv("PORT"))
}
