package service

import (
	"delos-farm-backend/domains"
	"delos-farm-backend/ponds/repository"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

var Ponds = []domains.Ponds{
	{ID: 1, Name: "Farm 1", Slug: "farm_1", FarmID: 1},
	{ID: 2, Name: "Farm 2", Slug: "farm_2", FarmID: 1},
}

var pondRepository = &repository.PondsRepositoryMock{Mock: mock.Mock{}}
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


