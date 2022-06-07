package handler

import (
	"delos-farm-backend/domains"
	"delos-farm-backend/domains/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
	"net/http"
	"net/http/httptest"
	// "encoding/json"
	"fmt"
)

var statsServicemock = new(mocks.StatsServiceMock)

func TestStatsHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("GetAllStatsSuccess", func(t *testing.T) {
		stats := []domains.StatsResults{
			{Path: "test 1", APICallCount: "2", UniqueCallCount: "1"},
			{Path: "test 2", APICallCount: "2", UniqueCallCount: "1"},
			{Path: "test 3", APICallCount: "2", UniqueCallCount: "1"},
		}
		limit := "10"
		offset := "0"

		statsServicemock.Mock.On("GetAllStats", limit, offset).Return(stats, nil)

		recorder := httptest.NewRecorder()

		router := gin.Default()
		group := router.Group("api/v1")

		statsHandler := StatsHandler{Service: statsServicemock}

		group.GET("/statistics", statsHandler.GetAll)

		request, err := http.NewRequest(
			"GET", 
			fmt.Sprintf("/api/v1/statistics?limit=%s&offset=%s", limit, offset),
			nil,
		)
		assert.NoError(t, err)

		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
