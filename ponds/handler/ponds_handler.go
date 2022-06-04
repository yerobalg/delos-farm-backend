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

type PondsHandler struct {
	service domains.PondsService
}

func NewPondsHandler(r *gin.RouterGroup, service domains.PondsService) {
	handler := &PondsHandler{service: service}
	api := r.Group("/ponds")
	{
		api.POST("/:farmId", handler.Create)
		api.DELETE("/:id", handler.Delete)
		api.GET("/:id", handler.Get)
		api.GET("/", handler.GetAll)
		api.PUT("/:id", handler.Update)
	}
}

//Create pond handler
func (h *PondsHandler) Create(c *gin.Context) {
	//get farm id from params
	farmId, _ := c.Params.Get("farmId")
	farmIdNum, err := strconv.Atoi(farmId)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, helpers.ResponseFormat(
			"Invalid farm id",
			false,
			nil,
		))
		return
	}

	//validate input
	var input domains.PondsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, helpers.ResponseFormat(
			"Please fill all required fields",
			false,
			nil,
		))
		return
	}

	//create pond entity
	pond := domains.Ponds{
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		Name:      input.Name,
		Slug:      slug.Make(input.Name),
		FarmID:    uint(farmIdNum),
	}

	//Create the pond, and will return error if insert duplicate name
	if err := h.service.Create(&pond); err != nil {
		statusCode := http.StatusInternalServerError
		//if error is duplicate key value and farm id not found
		if err.Error() == "Pond already exists" {
			statusCode = http.StatusConflict
		} else if err.Error() == "Farm not found" { 
			statusCode = http.StatusNotFound
		}

		c.JSON(statusCode, helpers.ResponseFormat(
			err.Error(),
			false,
			nil,
		))
		return
	}

	c.JSON(http.StatusCreated, helpers.ResponseFormat(
		"Successfully created pond",
		true,
		pond,
	))
}

//Delete pond handler
func (h *PondsHandler) Delete(c *gin.Context) {
	//Get Id from params
	id, _ := c.Params.Get("id")
	idNum, _ := strconv.Atoi(id)

	//Find pond by id, if not found return error
	pond, err := h.service.Get(uint(idNum))
	if err != nil && err.Error() == "Pond not found" {
		c.JSON(http.StatusNotFound, helpers.ResponseFormat(
			"Pond not found",
			false,
			nil,
		))
		return
	}

	//Delete the pond
	if err := h.service.Delete(&pond); err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(
			"Failed to delete pond",
			false,
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseFormat(
		"Successfully deleted pond",
		true,
		nil,
	))
}

//Get pond by id
func (h *PondsHandler) Get(c *gin.Context) {
	//Get id from params
	id, _ := c.Params.Get("id")
	idNum, _ := strconv.Atoi(id)

	//Find pond by id, if not found return error
	pond, err := h.service.Get(uint(idNum))
	if err != nil && err.Error() == "Pond not found" {
		c.JSON(http.StatusNotFound, helpers.ResponseFormat(
			"Pond not found",
			false,
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseFormat(
		"Successfully retrieved pond",
		true,
		pond,
	))
}

//Update pond handler
func (h *PondsHandler) Update(c *gin.Context) {
	//Get id from params
	id, _ := c.Params.Get("id")
	idNum, _ := strconv.Atoi(id)

	//Find pond by id, if not found return error
	pond, err := h.service.Get(uint(idNum))
	if err != nil && err.Error() == "Pond not found" {
		c.JSON(http.StatusNotFound, helpers.ResponseFormat("Pond not found",
			false,
			nil,
		))
		return
	}

	//validate input
	var input domains.PondsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, helpers.ResponseFormat(
			"Please fill all required fields",
			false,
			nil,
		))
		return
	}

	pond.Name = input.Name
	pond.Slug = slug.Make(input.Name)
	pond.UpdatedAt = time.Now().Unix()

	//Update pond, and will return error if insert duplicate name
	if err := h.service.Update(&pond); err != nil {
		statusCode := http.StatusInternalServerError

		//if error is duplicate key value
		if err.Error() == "Pond already exists" {
			statusCode = http.StatusConflict
		}

		c.JSON(statusCode, helpers.ResponseFormat(err.Error(), false, nil))
		return
	}

	c.JSON(http.StatusCreated, helpers.ResponseFormat(
		"Successfully updated pond",
		true,
		pond,
	))
}

//Get all ponds
func (h *PondsHandler) GetAll(c *gin.Context) {
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

	//get the ponds, and will return error if not found
	ponds, err := h.service.GetAll(limit, offset)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "No ponds found" {
			statusCode = http.StatusNotFound
		}

		c.JSON(statusCode, helpers.ResponseFormat(err.Error(), false, nil))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseFormat(
		"Successfully retrieved ponds",
		true,
		ponds,
	))
}
