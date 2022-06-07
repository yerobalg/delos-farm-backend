package handler

import (
	"delos-farm-backend/domains"
	"delos-farm-backend/domains/mocks"
	"delos-farm-backend/helpers"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var statsServicemock = new(mocks.StatsServiceMock)
var router = gin.Default()
var group = router.Group("api/v1")
var statsHandler = StatsHandler{Service: statsServicemock}

func TestStatsHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("GetAllStatsSuccess", getAllStatsSuccess)
}


func getAllStatsSuccess(t *testing.T) {
	stats := []domains.StatsResults{
		{Path: "test 1", APICallCount: "2", UniqueCallCount: "1"},
		{Path: "test 2", APICallCount: "2", UniqueCallCount: "1"},
		{Path: "test 3", APICallCount: "2", UniqueCallCount: "1"},
	}
	limit := "10"
	offset := "0"

	statsServicemock.Mock.On("GetAllStats", limit, offset).Return(stats, nil)

	group.GET("/statistics", statsHandler.GetAll)

	statusCode, response := beginTestRequest(
		"GET",
		fmt.Sprintf("/api/v1/statistics?limit=%s&offset=%s", limit, offset),
		t,
	)

	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, true, response.Success)
}

func beginTestRequest(
	method string,
	url string,
	tContext *testing.T,
) (int, helpers.Response) {
	recorder := httptest.NewRecorder()

	request, err := http.NewRequest(method, url, nil)
	assert.NoError(tContext, err)

	router.ServeHTTP(recorder, request)
	res := recorder.Result().Body
	defer res.Close()
	response := helpers.Response{}
	json.NewDecoder(res).Decode(&response)
	return recorder.Code, response
}
