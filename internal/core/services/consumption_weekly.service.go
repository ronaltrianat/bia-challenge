package services

import (
	"bia-challenge/internal/core/domain"
	"bia-challenge/internal/core/ports"

	"github.com/shopspring/decimal"
)

type consumptionWeeklyService struct {
	biaDB ports.BiaRepository
}

func NewConsumptionWeeklyService(biaDB ports.BiaRepository) *consumptionWeeklyService {
	return &consumptionWeeklyService{biaDB: biaDB}
}

func (srv *consumptionWeeklyService) execute(
	request *domain.GetEnergyConsumptionRequest) (*domain.GetEnergyConsumptionResponse, error) {

	weeklyConsumption, err := srv.biaDB.GetWeeklyConsumptionByMetersIdsAndBetweenDates(
		request.MetersIDs, request.StartDate, request.EndDate)
	if err != nil {
		return nil, err
	}

	consumptions := make(map[int]domain.DataGraph)
	for _, v := range weeklyConsumption {
		var dataGraph domain.DataGraph

		if value, found := consumptions[v.MeterID]; found {
			dataGraph = value
		} else {
			dataGraph = domain.DataGraph{MeterID: v.MeterID}
		}

		dataGraph.Period = append(dataGraph.Period, v.Week)

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
