package service

import (
	"delos-farm-backend/domains"
	"delos-farm-backend/domains/mocks"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
	"time"
	"bou.ke/monkey"
)

var Ponds = []domains.Ponds{
	{ID: 1, Name: "Farm 1", Slug: "1_farm_1", FarmID: 1},
	{ID: 2, Name: "Farm 2", Slug: "1_farm_2", FarmID: 1},
}

var pondRepository = &mocks.PondsRepositoryMock{Mock: mock.Mock{}}
var pondService = NewPondsService(pondRepository)

func TestPondsService_GetFound(t *testing.T) {
	pondRepository.Mock.On("Get", uint(1)).Return(Ponds[0], nil)

	pond, err := pondService.Get(uint(1))
	assert.Nil(t, err, "should not return error")
	assert.NotNil(t, pond, "Pond should exist")

	assert.Equal(
		t,
		Ponds[0].ID,
		pond.ID,
		fmt.Sprintf("Fetching wrong pond, id should be 1 but got %d", pond.ID),
	)
}

func TestPondsService_GetNotFound(t *testing.T) {
	pondRepository.Mock.On("Get", uint(3)).Return(
		nil, errors.New("Pond not found"),
	)

	pond, err := pondService.Get(uint(3))
	assert.NotNil(t, err, "Should return error")

	isPondEmpty := reflect.DeepEqual(domains.Ponds{}, pond)
	assert.True(t, isPondEmpty, "Pond object should be empty")
}

func TestPondsService_CreateSuccess(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	})
	name := "Pond 3"
	slug := "pond_3"
	pond := &domains.Ponds{
		Name: name,
		Slug: slug,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		FarmID: 1,
	}
	pondRepository.Mock.On("Create", pond).Return(nil)

	pond, err := pondService.Create(name, slug, 1)
	assert.Nil(t, err, "should not return error")
	assert.NotNil(t, pond, "Pond should exist")
}

func TestPondsService_UpdateSuccess(t *testing.T) {
	updatedPond := &domains.Ponds{
		ID:     3,
		Name:   "Pond 3",
		Slug:   "1_pond_3",
		FarmID: 1,
	}
	pondRepository.Mock.On("Update", updatedPond).Return(nil)

	err := pondService.Update(updatedPond)
	assert.Nil(t, err, "should not return error")
}

func TestPondsService_UpdateAlreadyExists(t *testing.T) {
	pondRepository.Mock.On("Update", &Ponds[0]).Return(
		errors.New("Pond already exists"),
	)

	err := pondService.Update(&Ponds[0])

	assert.NotNil(t, err, "should return pond already exists error")
}

func TestPondsService_DeleteSuccess(t *testing.T) {
	pondRepository.Mock.On("Delete", &Ponds[0]).Return(nil)

	err := pondService.Delete(&Ponds[0])
	assert.Nil(t, err, "should not return error")
}

func TestPondsService_GetAllSuccess(t *testing.T) {
	limit := "2"
	offset := "0"
	pondRepository.Mock.On("GetAll", 2, 0).Return(Ponds, nil)

	ponds, err := pondService.GetAll(limit, offset)
	assert.Nil(t, err, "should not return error")
	assert.NotNil(t, ponds, "Ponds should exist")
	assert.Equal(t, len(Ponds), len(ponds), "Must return 2 ponds")
}

func TestPondsService_GetAllNoData(t *testing.T) {
	limit := "2"
	offset := "5"
	pondRepository.Mock.On("GetAll", 2, 5).Return([]domains.Ponds{}, nil)

	ponds, err := pondService.GetAll(limit, offset)
	
	assert.NotNil(t, err, "should return no ponds found error")
	assert.Equal(t, 0, len(ponds), "Must return 0 ponds")
}
