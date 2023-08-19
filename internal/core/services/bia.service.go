package services

import (
	"bia-challenge/internal/core/domain"
	"bia-challenge/internal/core/ports"
)

type biaService struct {
	biaDB   ports.BiaRepository
	factory *energyConsumptionFactory
}

type energyConsumptionInterface interface {
	execute(request *domain.GetEnergyConsumptionRequest) (*domain.GetEnergyConsumptionResponse, error)
}

type energyConsumptionFactory struct {
	strategies map[domain.KindPeriod]energyConsumptionInterface
}

func NewBiaService(biaDB ports.BiaRepository) *biaService {
	consumptionFactory := newEnergyConsumptionFactory(biaDB)
	return &biaService{biaDB: biaDB, factory: consumptionFactory}
}

func (srv *biaService) GetEnergyConsumption(
	request *domain.GetEnergyConsumptionRequest) (*domain.GetEnergyConsumptionResponse, error) {
	return srv.factory.strategies[request.KindPeriod].execute(request)
}

func newEnergyConsumptionFactory(biaDB ports.BiaRepository) *energyConsumptionFactory {
	strategies := make(map[domain.KindPeriod]energyConsumptionInterface)
	strategies[domain.Monthly] = NewConsumptionMonthlyService(biaDB)
	strategies[domain.Daily] = NewConsumptionDailyService(biaDB)
	strategies[domain.Weekly] = NewConsumptionWeeklyService(biaDB)
	return &energyConsumptionFactory{strategies: strategies}
}
