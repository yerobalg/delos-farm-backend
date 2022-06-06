package middlewares

import (
	"delos-farm-backend/domains"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, authorization, accept, origin, Cache-Control, X-Requested-With, user-info",
		)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.Writer.Header().Set("Content-Type", "application/json")
			c.AbortWithStatus(204)
		} else {
			c.Next()
		}
	}
}

type StatsMiddleware struct {
	service domains.StatsService
}

func NewStatsMiddleware(service domains.StatsService) StatsMiddleware {
	return StatsMiddleware{service: service}
}

func (m StatsMiddleware) GetStatistics() gin.HandlerFunc {
	return func(c *gin.Context) {
		//get ip address
		ip := "ip_" + c.ClientIP()
		fmt.Println(ip)

		//get path
		path := fmt.Sprintf("%s_%s", c.Request.Method, c.Request.URL.Path)
		fmt.Println(path)

		//get statistics from service
		stats := m.service.GetStatistics(path, ip)

		//set statistics to context
		c.Set("stats", stats)
	}
}
