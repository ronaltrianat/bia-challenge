package ports

import "bia-challenge/internal/core/domain"

type BiaService interface {
	GetEnergyConsumption(request *domain.GetEnergyConsumptionRequest) (*domain.GetEnergyConsumptionResponse, error)
}
