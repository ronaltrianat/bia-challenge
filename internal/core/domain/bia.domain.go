package domain

import "github.com/shopspring/decimal"

type KindPeriod string

const (
	Monthly KindPeriod = "monthly"
	Weekly  KindPeriod = "weekly"
	Daily   KindPeriod = "daily"
)

type GetEnergyConsumptionRequest struct {
	MetersIDs  []int
	StartDate  string
	EndDate    string
	KindPeriod KindPeriod
}

type GetEnergyConsumptionResponse struct {
	DataGraph []DataGraph `json:"data_graph"`
}

type DataGraph struct {
	MeterID            int               `json:"meter_id"`
	Address            string            `json:"address"`
	Active             []decimal.Decimal `json:"active"`
	ReactiveInductive  []decimal.Decimal `json:"reactive_inductive"`
	ReactiveCapacitive []decimal.Decimal `json:"reactive_capacitive"`
	Exported           []decimal.Decimal `json:"exported"`
	Period             []string          `json:"period"`
}
