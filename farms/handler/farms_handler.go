package handler

import (
	"delos-farm-backend/domains"
	"delos-farm-backend/helpers"
	"delos-farm-backend/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"net/http"
	"strconv"
	"time"
)

type FarmsHandler struct {
	Service domains.FarmsService
}

func NewFarmsHandler(
	r *gin.RouterGroup,
	Service domains.FarmsService,
	statsMiddleware middlewares.StatsMiddleware,
) {
	handler := &FarmsHandler{Service: Service}
	api := r.Group("/farms")
	{
		api.POST("/", statsMiddleware.GetStatistics(), handler.Create)
		api.DELETE("/:id", statsMiddleware.GetStatistics(), handler.Delete)
		api.GET("/:id", statsMiddleware.GetStatistics(), handler.Get)
		api.GET("/", statsMiddleware.GetStatistics(), handler.GetAll)
		api.PUT("/:id", statsMiddleware.GetStatistics(), handler.Update)
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
	name:= input.Name
	slug:= slug.Make(name)

	//Create the farm, and will return error if insert duplicate name
	farm, err := h.Service.Create(name, slug);
	if  err != nil {
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
	farm, err := h.Service.Get(uint(idNum))
	if err != nil && err.Error() == "Farm not found" {
		c.JSON(http.StatusNotFound, helpers.ResponseFormat(
			"Farm not found",
			false,
			nil,
		))
		return
	}

	//Delete the farm
	if err := h.Service.Delete(&farm); err != nil {
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
	farm, err := h.Service.Get(uint(idNum))
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
	farm, err := h.Service.Get(uint(idNum))
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
	if err := h.Service.Update(&farm); err != nil {
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
	farms, err := h.Service.GetAll(limit, offset)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "No farms found" {
			statusCode = http.StatusNotFound
		}

		c.JSON(statusCode, helpers.ResponseFormat(
			err.Error(),
			false,
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseFormat(
		"Successfully retrieved farms",
		true,
		farms,
	))
}
