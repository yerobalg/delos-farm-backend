package handler

import (
	"delos-farm-backend/domains"
	"delos-farm-backend/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StatsHandler struct {
	Service domains.StatsService
}

func NewStatsHandler(r *gin.RouterGroup, Service domains.StatsService,) {
	handler := &StatsHandler{Service: Service}
	api := r.Group("/statistics")
	api.GET("/", handler.GetAll)
}

func (h *StatsHandler) GetAll(c *gin.Context) {
	//Get query params
	limit, isLimitExist := c.GetQuery("limit")
	offset, isOffsetExist := c.GetQuery("offset")

	//the default for limit is 100 and offset is 0
	if !isLimitExist {
		limit = "100"
	}
	if !isOffsetExist {
		offset = "0"
	}

	//Get statistics from services
	statistics, err := h.Service.GetAllStats(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(
			"Error getting statistics",
			false,
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseFormat(
		"Successfully retrieved statistics",
		true,
		statistics,
	))
}