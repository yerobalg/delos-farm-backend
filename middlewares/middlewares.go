package middlewares

import (
	"delos-farm-backend/domains"
	"fmt"
	"github.com/gin-gonic/gin"
	// "strings"
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
		//get ip address by X-Forwarded-For header
		clientIP := c.Request.Header.Get("X-Forwarded-For")

		// if empty, client might be using cloudflare
		// so get IP by CF-Connecting-IP header
		if len(clientIP) == 0 {
			clientIP = c.Request.Header.Get("CF-Connecting-IP")
		}

		// if still empty, client might be behind a proxy like nginx
		// so get IP by X-Real-IP header
		if len(clientIP) == 0 {
			clientIP = c.Request.Header.Get("X-Real-IP")
		}

		// then the 
		if len(clientIP) == 0 {
			clientIP = c.Request.RemoteAddr
		}

		//get path
		path := fmt.Sprintf("%s_%s", c.Request.Method, c.Request.URL.Path)
		fmt.Println(path)

		//get statistics from service
		stats := m.service.GetStatistics(path, clientIP)

		//set statistics to context
		c.Set("stats", stats)
	}
}
