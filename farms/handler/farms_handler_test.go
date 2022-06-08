package handler

import (
	"bou.ke/monkey"
	"bytes"
	"delos-farm-backend/domains"
	"delos-farm-backend/domains/mocks"
	"delos-farm-backend/helpers"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var farmsServiceMock = new(mocks.FarmsServiceMock)
var router = gin.Default()
var group = router.Group("api/v1")
var farmsHandler = FarmsHandler{Service: farmsServiceMock}

var Farms = []domains.Farms{
	{
		ID:   0,
		Name: "Farm 0",
		Slug: "farm-0",
	},
	{
		ID:   0,
		Name: "Farm 1",
		Slug: "farm-1",
	},
}

func TestFarmsHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("CreateSuccess", createSuccess)
	t.Run("GetSuccess", getSuccess)
	t.Run("UpdateSuccess", updateSuccess)
	t.Run("DeleteSuccess", deleteSuccess)
	t.Run("GetAllSuccess", getAllSuccess)
}

func createSuccess(t *testing.T) {
	farmInput := domains.FarmsInput{Name: "Farm 1"}

	farmsServiceMock.Mock.On(
		"Create",
		Farms[1].Name,
		Farms[1].Slug,
	).Return(&Farms[1], nil)

	group.POST("/farms", farmsHandler.Create)

	code, response := beginTestRequest(
		"POST",
		"/api/v1/farms",
		farmInput,
		t,
	)

	assert.Equal(t, http.StatusCreated, code)
	assert.Equal(t, true, response.Success)
}

func getSuccess(t *testing.T) {
	farmsServiceMock.Mock.On("Get", uint(0)).Return(Farms[0], nil)

	group.GET("/farms/:id", farmsHandler.Get)

	code, response := beginTestRequest(
		"GET",
		"/api/v1/farms/0",
		domains.FarmsInput{},
		t,
	)

	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, true, response.Success)
}

func updateSuccess(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	})
	Farms[1].UpdatedAt = time.Now().Unix()

	farmsServiceMock.Mock.On("Get", uint(0)).Return(Farms[1], nil)
	farmsServiceMock.Mock.On("Update", &Farms[1]).Return(nil)

	group.PUT("/farms/:id", farmsHandler.Update)

	code, response := beginTestRequest(
		"PUT",
		"/api/v1/farms/1",
		domains.FarmsInput{Name: "Farm 1"},
		t,
	)

	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, true, response.Success)
}

func deleteSuccess(t *testing.T) {
	farmsServiceMock.Mock.On("Get", uint(0)).Return(Farms[0], nil)
	farmsServiceMock.Mock.On("Delete", &Farms[0]).Return(nil)

	group.DELETE("/farms/:id", farmsHandler.Delete)

	code, response := beginTestRequest(
		"DELETE",
		"/api/v1/farms/0",
		domains.FarmsInput{},
		t,
	)

	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, true, response.Success)
}

func getAllSuccess(t *testing.T) {
	farmsServiceMock.Mock.On("GetAll", "10", "0").Return(Farms, nil)

	group.GET("/farms", farmsHandler.GetAll)

	code, response := beginTestRequest(
		"GET",
		"/api/v1/farms?limit=10&offset=0",
		domains.FarmsInput{},
		t,
	)

	assert.Equal(t, http.StatusOK, code)
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
