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

var pondsServiceMock = new(mocks.PondsServiceMock)
var router = gin.Default()
var group = router.Group("api/v1")
var pondsHandler = PondsHandler{Service: pondsServiceMock}

var Ponds = []domains.Ponds{
	{
		ID:   0,
		Name: "Pond 0",
		Slug: "pond-0",
		FarmID: 0,
	},
	{
		ID:   1,
		Name: "Pond 1",
		Slug: "pond-1",
		FarmID: 1,
	},
	{
		ID:   2,
		Name: "Pond 1",
		Slug: "1_pond-1",
		FarmID: 1,
	},
}

func TestPondsHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("CreateSuccess", createSuccess)
	t.Run("GetSuccess", getSuccess)
	t.Run("UpdateSuccess", updateSuccess)
	t.Run("DeleteSuccess", deleteSuccess)
	t.Run("GetAllSuccess", getAllSuccess)
}

func createSuccess(t *testing.T) {
	pondInput := domains.PondsInput{Name: "Pond 1"}

	pondsServiceMock.Mock.On(
		"Create",
		Ponds[2].Name,
		Ponds[2].Slug,
		Ponds[2].FarmID,
	).Return(&Ponds[2], nil)

	group.POST("/ponds/:farmId", pondsHandler.Create)

	code, response := beginTestRequest(
		"POST",
		"/api/v1/ponds/1",
		pondInput,
		t,
	)

	assert.Equal(t, http.StatusCreated, code)
	assert.Equal(t, true, response.Success)
}

func getSuccess(t *testing.T) {
	pondsServiceMock.Mock.On("Get", uint(1)).Return(Ponds[1], nil)

	group.GET("/ponds/:id", pondsHandler.Get)

	code, response := beginTestRequest(
		"GET",
		"/api/v1/ponds/1",
		domains.PondsInput{},
		t,
	)

	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, true, response.Success)
}

func updateSuccess(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	})
	Ponds[1].UpdatedAt = time.Now().Unix()

	pondsServiceMock.Mock.On("Get", uint(1)).Return(Ponds[1], nil)
	pondsServiceMock.Mock.On("Update", &Ponds[1]).Return(nil)

	group.PUT("/ponds/:id", pondsHandler.Update)

	code, response := beginTestRequest(
		"PUT",
		"/api/v1/ponds/1",
		domains.PondsInput{Name: "Pond 1"},
		t,
	)

	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, true, response.Success)
}

func deleteSuccess(t *testing.T) {
	pondsServiceMock.Mock.On("Get", uint(0)).Return(Ponds[0], nil)
	pondsServiceMock.Mock.On("Delete", &Ponds[0]).Return(nil)

	group.DELETE("/ponds/:id", pondsHandler.Delete)

	code, response := beginTestRequest(
		"DELETE",
		"/api/v1/ponds/0",
		domains.PondsInput{},
		t,
	)

	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, true, response.Success)
}

func getAllSuccess(t *testing.T) {
	pondsServiceMock.Mock.On("GetAll", "10", "0").Return(Ponds, nil)

	group.GET("/ponds", pondsHandler.GetAll)

	code, response := beginTestRequest(
		"GET",
		"/api/v1/ponds?limit=10&offset=0",
		domains.PondsInput{},
		t,
	)

	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, true, response.Success)
}

func beginTestRequest(
	method string,
	url string,
	body domains.PondsInput,
	tContext *testing.T,
) (int, helpers.Response) {
	recorder := httptest.NewRecorder()

	var request *http.Request
	var err error

	if (body != domains.PondsInput{}) {
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
