package ports

import "bia-challenge/internal/core/domain"

type BiaServicePort interface {
	GetEnergyConsumption(request *domain.GetEnergyConsumptionRequest) (*domain.GetEnergyConsumptionResponse, error)
}
