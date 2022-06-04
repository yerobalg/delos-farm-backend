package handler

import (
	"delos-farm-backend/domains"
	"delos-farm-backend/helpers"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"net/http"
	"strconv"
	"time"
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
		api.GET("/:id", handler.Get)
		api.GET("/", handler.GetAll)
		api.PUT("/:id", handler.Update)
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
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		Name:      input.Name,
		Slug:      slug.Make(input.Name),
	}

	//Create the farm, and will return error if insert duplicate name
	if err := h.service.Create(&farm); err != nil {
		statusCode := http.StatusInternalServerError
		//if error is duplicate key value
		if err.Error() == "Farm already exists" {
			statusCode = http.StatusConflict
		}

		c.JSON(statusCode, helpers.ResponseFormat(
			err.Error(),
			false,
			nil,
		))
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
	farm, err := h.service.Get(uint(idNum))
	if err != nil && err.Error() == "Farm not found" {
		c.JSON(http.StatusNotFound, helpers.ResponseFormat(
			"Farm not found",
			false,
			nil,
		))
		return
	}

	//Delete the farm
	if err := h.service.Delete(&farm); err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(
			"Failed to delete farm",
			false,
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseFormat(
		"Successfully deleted farm",
		true,
		nil,
	))
}

//Get farm by id
func (h *FarmsHandler) Get(c *gin.Context) {
	//Get id from params
	id, _ := c.Params.Get("id")
	idNum, _ := strconv.Atoi(id)

	//Find farm by id, if not found return error
	farm, err := h.service.Get(uint(idNum))
	if err != nil && err.Error() == "Farm not found" {
		c.JSON(http.StatusNotFound, helpers.ResponseFormat(
			"Farm not found",
			false,
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseFormat(
		"Successfully retrieved farm",
		true,
		farm,
	))
}

//Update farm handler
func (h *FarmsHandler) Update(c *gin.Context) {
	//Get id from params
	id, _ := c.Params.Get("id")
	idNum, _ := strconv.Atoi(id)

	//Find farm by id, if not found return error
	farm, err := h.service.Get(uint(idNum))
	if err != nil && err.Error() == "Farm not found" {
		c.JSON(http.StatusNotFound, helpers.ResponseFormat("Farm not found",
			false,
			nil,
		))
		return
	}

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

	farm.Name = input.Name
	farm.Slug = slug.Make(input.Name)
	farm.UpdatedAt = time.Now().Unix()

	//Update farm, and will return error if insert duplicate name
	if err := h.service.Update(&farm); err != nil {
		statusCode := http.StatusInternalServerError

		//if error is duplicate key value
		if err.Error() == "Farm already exists" {
			statusCode = http.StatusConflict
		}

		c.JSON(statusCode, helpers.ResponseFormat(err.Error(), false, nil))
		return
	}

	c.JSON(http.StatusCreated, helpers.ResponseFormat(
		"Successfully updated farm",
		true,
		farm,
	))
}

//Get all farms
func (h *FarmsHandler) GetAll(c *gin.Context) {
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

	//get the farms, and will return error if not found
	_, err := h.service.GetAll(limit, offset)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "No farms found" {
			statusCode = http.StatusNotFound
		}

		c.JSON(statusCode, helpers.ResponseFormat(err.Error(), false, nil))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseFormat(
		"Successfully retrieved farms",
		true,
		time.Now().Unix(),
	))
}
