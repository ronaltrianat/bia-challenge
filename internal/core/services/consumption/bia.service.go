package consumption

import (
	"bia-challenge/internal/core/domain"
	"bia-challenge/internal/core/ports"
)

type biaService struct {
	biaDB            ports.BiaRepositoryPort
	addressesService ports.AddressesServicePort
	factory          *energyConsumptionFactory
}

type energyConsumptionInterface interface {
	execute(request *domain.GetEnergyConsumptionRequest) (*domain.GetEnergyConsumptionResponse, error)
}

type energyConsumptionFactory struct {
	strategies map[domain.KindPeriod]energyConsumptionInterface
}

func NewBiaService(biaDB ports.BiaRepositoryPort, addressesService ports.AddressesServicePort) *biaService {
	consumptionFactory := newEnergyConsumptionFactory(biaDB, addressesService)
	return &biaService{biaDB: biaDB, factory: consumptionFactory, addressesService: addressesService}
}

func (srv *biaService) GetEnergyConsumption(
	request *domain.GetEnergyConsumptionRequest) (*domain.GetEnergyConsumptionResponse, error) {
	return srv.factory.strategies[request.KindPeriod].execute(request)
}

func newEnergyConsumptionFactory(biaDB ports.BiaRepositoryPort,
	addressesService ports.AddressesServicePort) *energyConsumptionFactory {
	strategies := make(map[domain.KindPeriod]energyConsumptionInterface)
	strategies[domain.Monthly] = NewConsumptionMonthlyService(biaDB, addressesService)
	strategies[domain.Daily] = NewConsumptionDailyService(biaDB, addressesService)
	strategies[domain.Weekly] = NewConsumptionWeeklyService(biaDB, addressesService)
	return &energyConsumptionFactory{strategies: strategies}
}
