package handler

import (
	"delos-farm-backend/domains"
	"delos-farm-backend/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

type FarmsHandler struct {
	service domains.FarmsService
}

func NewFarmsHandler(r *gin.RouterGroup, service domains.FarmsService) {
	handler := &FarmsHandler{service: service}
	r.Group("/farms")
	{
		r.POST("/", handler.Create)
	}
}

func (h *FarmsHandler) Create(c *gin.Context) {
	var input domains.FarmsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, helpers.ResponseFormat(
			"Please fill all required fields",
			false,
			nil,
		))
		return
	}

	farm := domains.Farms{
		Name: input.Name,
		Slug: slug.Make(input.Name),
	}

	if err := h.service.Create(&farm); err != nil {
		if err.Error() == "Farm already exists" {
			c.JSON(http.StatusConflict, helpers.ResponseFormat(
				err.Error(),
				false,
				nil,
			))
		} else {
			c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(
				"Failed to create farm",
				false,
				nil,
			))
		}
		return
	}

	c.JSON(http.StatusCreated, helpers.ResponseFormat(
		"Successfully created farm",
		true,
		farm,
	))
}
