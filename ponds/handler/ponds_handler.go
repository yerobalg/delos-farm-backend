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
		api.POST("/", handler.Create)
		api.DELETE("/:id", handler.Delete)
		api.GET("/:id", handler.Get)
		api.GET("/", handler.GetAll)
		api.PUT("/:id", handler.Update)
	}
}

//Create pond handler
func (h *PondsHandler) Create(c *gin.Context) {
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
	}

	//Create the pond, and will return error if insert duplicate name
	if err := h.service.Create(&pond); err != nil {
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
	if err != nil && err.Error() == "Farm not found" {
		c.JSON(http.StatusNotFound, helpers.ResponseFormat(
			"Farm not found",
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

