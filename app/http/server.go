package main

import (
	"delos-farm-backend/bootstrap"
	_statsService "delos-farm-backend/stats/service"
	_statsRepository "delos-farm-backend/stats/repository"
	farmsHandler "delos-farm-backend/farms/handler"
	_farmsRepository "delos-farm-backend/farms/repository"
	_farmsService "delos-farm-backend/farms/service"
	"delos-farm-backend/middlewares"
	pondsHandler "delos-farm-backend/ponds/handler"
	_pondsRepository "delos-farm-backend/ponds/repository"
	_pondsService "delos-farm-backend/ponds/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	//init env variable
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("failed to load env from local file")
	}

	//init database
	db, err := bootstrap.InitPostgres()
	if err != nil {
		panic(err)
	}

	//init router
	engine := gin.Default()
	engine.Use(middlewares.CorsMiddleware())
	router := engine.Group("/api/v1")

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Delos Farm API",
			"succes":  true,
			"data":    nil,
		})
	})

	//init stats
	statsRepository := _statsRepository.NewStatsRepository(db)
	statsService := _statsService.NewStatsService(statsRepository)
	statsMiddleware := middlewares.NewStatsMiddleware(statsService)	

	//init farms
	farmsRepository := _farmsRepository.NewFarmsRepository(db)
	farmsService := _farmsService.NewFarmsService(farmsRepository)
	farmsHandler.NewFarmsHandler(router, farmsService, statsMiddleware)

	//init ponds
	pondsRepository := _pondsRepository.NewPondsRepository(db)
	pondsService := _pondsService.NewPondsService(pondsRepository)
	pondsHandler.NewPondsHandler(router, pondsService, statsMiddleware)

	//run server
	engine.Run(":" + os.Getenv("PORT"))
}
