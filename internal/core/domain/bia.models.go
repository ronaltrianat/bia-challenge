package domain

import "time"

type ConsumptionModel struct {
	MeterID            int
	ActiveEnergy       string
	ReactiveEnergy     string
	CapacitiveReactive string
	Solar              string
}

type DailyConsumption struct {
	ConsumptionModel
	Date time.Time
}

type WeeklyConsumption struct {
	ConsumptionModel
	Week string
}

type MonthlyConsumption struct {
	ConsumptionModel
	MonthYear string
}
