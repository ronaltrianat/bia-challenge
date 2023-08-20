package services_test

import (
	"bia-challenge/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

const (
	_DailyConsumptionRepositoryResponseOkFile = "../../../tests/resources/services/daily_consumption_repository_response_ok.json"
	_GetDailyConsumptionServiceResponseOk     = "../../../tests/resources/services/get_daily_consumption_service_response_ok.json"

	_WeeklyConsumptionRepositoryResponseOkFile = "../../../tests/resources/services/weekly_consumption_repository_response_ok.json"
	_GetWeeklyConsumptionServiceResponseOk     = "../../../tests/resources/services/get_weekly_consumption_service_response_ok.json"

	_MonthlyConsumptionRepositoryResponseOkFile = "../../../tests/resources/services/monthly_consumption_repository_response_ok.json"
	_GetMonthlyConsumptionServiceResponseOk     = "../../../tests/resources/services/get_monthly_consumption_service_response_ok.json"
)

type biaRepositoryMockPort struct {
	mock.Mock
}

func (m *biaRepositoryMockPort) GetDailyConsumptionByMetersIdsAndBetweenDates(
	metersIDs []int, startDate, endDate string) ([]domain.DailyConsumption, error) {

	ret := m.Called(metersIDs, startDate, endDate)

	var r0 []domain.DailyConsumption
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]domain.DailyConsumption)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (m *biaRepositoryMockPort) GetWeeklyConsumptionByMetersIdsAndBetweenDates(
	metersIDs []int, startDate, endDate string) ([]domain.WeeklyConsumption, error) {
	ret := m.Called(metersIDs, startDate, endDate)

	var r0 []domain.WeeklyConsumption
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]domain.WeeklyConsumption)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (m *biaRepositoryMockPort) GetMonthlyConsumptionByMetersIdsAndBetweenDates(
	metersIDs []int, startDate, endDate string) ([]domain.MonthlyConsumption, error) {
	ret := m.Called(metersIDs, startDate, endDate)

	var r0 []domain.MonthlyConsumption
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]domain.MonthlyConsumption)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}
