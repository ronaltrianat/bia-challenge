package services

import (
	"bia-challenge/internal/core/domain"
	"bia-challenge/internal/core/ports"

	"github.com/shopspring/decimal"
)

type consumptionMonthlyService struct {
	biaDB ports.BiaRepository
}

func NewConsumptionMonthlyService(biaDB ports.BiaRepository) *consumptionMonthlyService {
	return &consumptionMonthlyService{biaDB: biaDB}
}

func (srv *consumptionMonthlyService) execute(
	request *domain.GetEnergyConsumptionRequest) (*domain.GetEnergyConsumptionResponse, error) {

	monthlyConsumption, err := srv.biaDB.GetMonthlyConsumptionByMetersIdsAndBetweenDates(
		request.MetersIDs, request.StartDate, request.EndDate)
	if err != nil {
		return nil, err
	}

	consumptions := make(map[int]domain.DataGraph)
	for _, v := range monthlyConsumption {
		var dataGraph domain.DataGraph

		if value, found := consumptions[v.MeterID]; found {
			dataGraph = value
		} else {
			dataGraph = domain.DataGraph{MeterID: v.MeterID}
		}

		dataGraph.Period = append(dataGraph.Period, v.MonthYear)

		if value, err := decimal.NewFromString(v.ActiveEnergy); err == nil {
			dataGraph.Active = append(dataGraph.Active, value)
		}

		if value, err := decimal.NewFromString(v.CapacitiveReactive); err == nil {
			dataGraph.ReactiveCapacitive = append(dataGraph.ReactiveCapacitive, value)
		}

		if value, err := decimal.NewFromString(v.ReactiveEnergy); err == nil {
			dataGraph.ReactiveInductive = append(dataGraph.ReactiveInductive, value)
		}

		if value, err := decimal.NewFromString(v.Solar); err == nil {
			dataGraph.Exported = append(dataGraph.Exported, value)
		}

		consumptions[v.MeterID] = dataGraph
	}

	var dataGraphList []domain.DataGraph
	for _, v := range consumptions {
		dataGraphList = append(dataGraphList, v)
	}

	response := &domain.GetEnergyConsumptionResponse{DataGraph: dataGraphList}

	return response, nil
}
