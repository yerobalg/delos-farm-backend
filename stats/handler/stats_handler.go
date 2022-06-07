package handler

import (
	"delos-farm-backend/domains"
	"delos-farm-backend/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StatsHandler struct {
	service domains.StatsService
}

func NewStatsHandler(r *gin.RouterGroup, service domains.StatsService,) {
	handler := &StatsHandler{service: service}
	api := r.Group("/statistics")
	api.GET("/", handler.GetAll)
}

func (h *StatsHandler) GetAll(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")

	statistics, err := h.service.GetAllStats(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(
			"Error getting statistics",
			false,
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseFormat(
		"Successfully got statistics",
		true,
		statistics,
	))
}