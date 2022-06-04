package middlewares

import (
	"context"
	"delos-farm-backend/domains"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Middleware struct{}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, authorization, accept, origin, Cache-Control, X-Requested-With, user-info")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.Writer.Header().Set("Content-Type", "application/json")
			c.AbortWithStatus(204)
		} else {
			c.Next()
		}
	}
}

func StatsMiddleware(redisClient *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiCountKey := c.Request.Method + "_" + c.Request.URL.Path
		uniqueCallCountKey := "ip_" + c.ClientIP()

		apiCountRes, err := redisClient.Incr(
			context.Background(), apiCountKey,
		).Result()
		if err != nil {
			c.Next()
			return
		}

		uniqueCallCountRes, err := redisClient.Incr(
			context.Background(), uniqueCallCountKey,
		).Result()
		if err != nil {
			c.Next()
			return
		}

		stats := domains.Stats{
			APICount: apiCountRes,
			UniqueCallCount: uniqueCallCountRes,
		}

		c.Set("stats", stats)
	}
}
