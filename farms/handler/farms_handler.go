package handler

import (
	"delos-farm-backend/domains"
	"delos-farm-backend/helpers"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

type FarmsHandler struct {
	service domains.FarmsService
}

func NewFarmsHandler(r *gin.RouterGroup, service domains.FarmsService) {
	handler := &FarmsHandler{service: service}
	api := r.Group("/farms")
	{
		api.POST("/", handler.Create)
		api.DELETE("/:id", handler.Delete)
	}
}

//Create farm handler
func (h *FarmsHandler) Create(c *gin.Context) {
	//validate input
	var input domains.FarmsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, helpers.ResponseFormat(
			"Please fill all required fields",
			false,
			nil,
		))
		return
	}
	
	//create farm entity
	farm := domains.Farms{
		Name: input.Name,
		Slug: slug.Make(input.Name),
	}

	//Create the farm, and will return error if insert duplicate name
	if err := h.service.Create(&farm); err != nil {
		//if error is duplicate key value
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

//Delete farm handler
func (h *FarmsHandler) Delete(c *gin.Context) {
	//Get Id from params
	id, _ := c.Params.Get("id")
	idNum, _ := strconv.Atoi(id) 

	//Find farm by id, if not found return error
	farm, err := h.service.Get(uint(idNum));
	if err != nil && err.Error() == "Farm not found" {
		c.JSON(http.StatusNotFound, helpers.ResponseFormat(
			"Farm not found",
			false,
			nil,
		))
	}

	//Delete the farm
	if err := h.service.Delete(&farm); err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(
			"Failed to delete farm",
			false,
			nil,
		))
	}

	c.JSON(http.StatusOK, helpers.ResponseFormat(
		"Successfully deleted farm",
		true,
		nil,
	))
}

//Get farm by id
func (h* FarmsHandler) Get(c *gin.Context) {
	//Get id from params
	id, _ := c.Params.Get("id")
	idNum, _ := strconv.Atoi(id) 

	//Find farm by id, if not found return error
	farm, err := h.service.Get(uint(idNum));
	if err != nil && err.Error() == "Farm not found" {
		c.JSON(http.StatusNotFound, helpers.ResponseFormat(
			"Farm not found",
			false,
			nil,
		))
	}

	c.JSON(http.StatusOK, helpers.ResponseFormat(
		"Successfully retrieved farm",
		true,
		farm,
	))
}