package handler

import (
	"bytes"
	"delos-farm-backend/domains"
	"delos-farm-backend/domains/mocks"
	"delos-farm-backend/helpers"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var farmsServiceMock = new(mocks.FarmsServiceMock)
var router = gin.Default()
var group = router.Group("api/v1")
var farmsHandler = FarmsHandler{Service: farmsServiceMock}

func TestFarmsHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("CreateSuccess", createSuccess)
}

func createSuccess(t *testing.T) {
	farmInput := domains.FarmsInput{Name: "Farm 1"}
	farm := domains.Farms{Name: farmInput.Name, Slug: "farm_1"}

	farmsServiceMock.Mock.On("Create", &farm).Return(nil)

	group.POST("/farms", farmsHandler.Create)

	code, response := beginTestRequest(
		"POST",
		"/api/v1/farms",
		farmInput,
		t,
	)

	fmt.Println(response.Message)

	assert.Equal(t, http.StatusCreated, code)
	assert.Equal(t, true, response.Success)
}

func beginTestRequest(
	method string,
	url string,
	body domains.FarmsInput,
	tContext *testing.T,
) (int, helpers.Response) {
	recorder := httptest.NewRecorder()

	var request *http.Request
	var err error

	if (body != domains.FarmsInput{}) {
		reqBody, err := json.Marshal(body)
		assert.NoError(tContext, err, "Should not have error while marshalling")
		request, err = http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	} else {
		request, err = http.NewRequest(method, url, nil)
	}

	assert.NoError(tContext, err)

	router.ServeHTTP(recorder, request)
	res := recorder.Result().Body
	defer res.Close()
	response := helpers.Response{}
	json.NewDecoder(res).Decode(&response)
	return recorder.Code, response
}
