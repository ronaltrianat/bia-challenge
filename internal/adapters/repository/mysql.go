package repository

import (
	"bia-challenge/internal/core/domain"

	"gorm.io/gorm"
)

type mysqlRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) *mysqlRepository {
	return &mysqlRepository{db: db}
}

func (repo *mysqlRepository) GetDailyConsumptionByMetersIdsAndBetweenDates(
	metersIDs []int, startDate, endDate string) ([]domain.DailyConsumption, error) {

	var dailyConsumption []domain.DailyConsumption

	query := `
	SELECT ec.meter_id AS meter_id,
		date(ec.date) AS date,
		SUM(ec.active_energy) AS active_energy,
		SUM(ec.reactive_energy) AS reactive_energy,
		SUM(ec.capacitive_reactive) AS capacitive_reactive,
		SUM(ec.solar) AS solar
	FROM energy_consumption ec
	WHERE ec.meter_id in (?)
	AND ec.date BETWEEN ? AND ?
	GROUP BY meter_id, date(ec.date)
	ORDER BY meter_id, date(ec.date) ASC`

	if trx := repo.db.Raw(query, metersIDs, startDate, endDate).
		Scan(&dailyConsumption); trx.Error != nil {
		return nil, trx.Error
	}

	return dailyConsumption, nil
}

func (repo *mysqlRepository) GetWeeklyConsumptionByMetersIdsAndBetweenDates(
	metersIDs []int, startDate, endDate string) ([]domain.WeeklyConsumption, error) {

	var weeklyConsumption []domain.WeeklyConsumption

	query := `
	SELECT ec.meter_id AS meter_id,
		CONCAT(date(min(ec.date)), ' - ', date(max(ec.date))) AS week,
		SUM(ec.active_energy) AS active_energy,
		SUM(ec.reactive_energy) AS reactive_energy,
		SUM(ec.capacitive_reactive) AS capacitive_reactive,
		SUM(ec.solar) AS solar
	FROM energy_consumption ec
	WHERE ec.meter_id in (?)
	AND ec.date BETWEEN ? AND ?
	GROUP BY meter_id, YEARWEEK(ec.date, 1)
	ORDER BY meter_id, YEARWEEK(ec.date, 1) ASC`

	if trx := repo.db.Raw(query, metersIDs, startDate, endDate).
		Scan(&weeklyConsumption); trx.Error != nil {
		return nil, trx.Error
	}

	return weeklyConsumption, nil
}

func (repo *mysqlRepository) GetMonthlyConsumptionByMetersIdsAndBetweenDates(
	metersIDs []int, startDate, endDate string) ([]domain.MonthlyConsumption, error) {

	var monthlyConsumption []domain.MonthlyConsumption

	query := `
	SELECT ec.meter_id AS meter_id,
		UPPER(CONCAT(MONTHNAME(ec.date), ' ', YEAR(ec.date))) AS month_year,
		SUM(ec.active_energy) AS active_energy,
		SUM(ec.reactive_energy) AS reactive_energy,
		SUM(ec.capacitive_reactive) AS capacitive_reactive,
		SUM(ec.solar) AS solar
	FROM energy_consumption ec
	WHERE ec.meter_id in (?)
	AND ec.date BETWEEN ? AND ?
	GROUP BY meter_id, month_year
	ORDER BY meter_id, month_year ASC`

	if trx := repo.db.Raw(query, metersIDs, startDate, endDate).
		Scan(&monthlyConsumption); trx.Error != nil {
		return nil, trx.Error
	}

	return monthlyConsumption, nil
}
