package services_test

import (
	"bia-challenge/internal/core/domain"
	"bia-challenge/internal/core/services"
	"encoding/json"
	"errors"
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetEnergyConsumptionMonthly_HappyPath(t *testing.T) {
	biaRepositoryMock := new(biaRepositoryMockPort)

	var repositoryResponse []domain.MonthlyConsumption
	content, _ := os.ReadFile(_MonthlyConsumptionRepositoryResponseOkFile)
	_ = json.Unmarshal(content, &repositoryResponse)

	biaRepositoryMock.On("GetMonthlyConsumptionByMetersIdsAndBetweenDates",
		mock.Anything, mock.Anything, mock.Anything).Return(repositoryResponse, nil)

	biaService := services.NewBiaService(biaRepositoryMock)

	request := &domain.GetEnergyConsumptionRequest{
		MetersIDs:  []int{1, 2},
		StartDate:  "2022-06-01",
		EndDate:    "2023-07-10",
		KindPeriod: domain.Monthly,
	}

	response, err := biaService.GetEnergyConsumption(request)
	sort.Slice(response.DataGraph, func(i, j int) bool {
		return response.DataGraph[i].MeterID < response.DataGraph[j].MeterID
	})

	currentResponseJson, _ := json.Marshal(response)
	expectedResponse, _ := os.ReadFile(_GetMonthlyConsumptionServiceResponseOk)

	assert.JSONEq(t, string(expectedResponse), string(currentResponseJson))
	assert.Nil(t, err)
}

func TestGetEnergyConsumptionMonthly_RepositoryError(t *testing.T) {
	biaRepositoryMock := new(biaRepositoryMockPort)

	someError := errors.New("some error")
	biaRepositoryMock.On("GetMonthlyConsumptionByMetersIdsAndBetweenDates",
		mock.Anything, mock.Anything, mock.Anything).Return(nil, someError)

	biaService := services.NewBiaService(biaRepositoryMock)

	request := &domain.GetEnergyConsumptionRequest{
		MetersIDs:  []int{1, 2},
		StartDate:  "2022-06-01",
		EndDate:    "2023-07-10",
		KindPeriod: domain.Monthly,
	}

	response, err := biaService.GetEnergyConsumption(request)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "some error")
}
