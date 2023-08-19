package ports

import "bia-challenge/internal/core/domain"

type BiaRepository interface {
	GetDailyConsumptionByMetersIdsAndBetweenDates(metersIDs []int, startDate, endDate string) ([]domain.DailyConsumption, error)
	GetWeeklyConsumptionByMetersIdsAndBetweenDates(metersIDs []int, startDate, endDate string) ([]domain.WeeklyConsumption, error)
	GetMonthlyConsumptionByMetersIdsAndBetweenDates(metersIDs []int, startDate, endDate string) ([]domain.MonthlyConsumption, error)
}
