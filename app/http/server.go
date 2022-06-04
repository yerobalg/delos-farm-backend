package main

import (
	"delos-farm-backend/middlewares"
	"github.com/gin-gonic/gin"
	"os"
	"github.com/joho/godotenv"
	"fmt"
	"delos-farm-backend/bootstrap"
	_farmsRepository "delos-farm-backend/farms/repository"
	_farmsService "delos-farm-backend/farms/service"
	farmsHandler "delos-farm-backend/farms/handler"
	_pondsRepository "delos-farm-backend/ponds/repository"
	_pondsService "delos-farm-backend/ponds/service"
	pondsHandler "delos-farm-backend/ponds/handler"

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
	farmsRepository := _farmsRepository.NewFarmsRepository(db)
	farmsService := _farmsService.NewFarmsService(farmsRepository)
	farmsHandler.NewFarmsHandler(router, farmsService)

	//init ponds
	pondsRepository := _pondsRepository.NewPondsRepository(db)
	pondsService := _pondsService.NewPondsService(pondsRepository)
	pondsHandler.NewPondsHandler(router, pondsService)

	//run server
	engine.Run(":" + os.Getenv("PORT"))
}
