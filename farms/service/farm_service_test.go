package service

import (
	"delos-farm-backend/domains"
	"delos-farm-backend/farms/repository"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

var Farms = []domains.Farms{{ID: 1, Name: "Farm 1"}, {ID: 2, Name: "Farm 2"}}

var farmRepository = &repository.FarmsRepositoryMock{Mock: mock.Mock{}}
var farmService = NewFarmsService(farmRepository)

func TestCategoryService_GetFound(t *testing.T) {
	farmRepository.Mock.On("Get", uint(1)).Return(Farms[0], nil)

	farm, err := farmService.Get(uint(1))
	assert.Nil(t, err, "should not return error")
	assert.NotNil(t, farm, "Farm should exist")

	assert.Equal(
		t,
		Farms[0].ID,
		farm.ID,
		fmt.Sprintf("Fetching wrong farm, id should be 1 but got %d", farm.ID),
	)
}

func TestCategoryService_GetNotFound(t *testing.T) {
	farmRepository.Mock.On("Get", uint(3)).Return(
		nil, errors.New("Farm not found"),
	)

	farm, err := farmService.Get(uint(3))
	assert.NotNil(t, err, "Should return error")

	isFarmEmpty := reflect.DeepEqual(domains.Farms{}, farm)
	assert.True(t, isFarmEmpty, "Farm object should be empty")
}
