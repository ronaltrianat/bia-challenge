package repository_test

import (
	"bia-challenge/internal/adapters/repository"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGetDailyConsumptionByMetersIdsAndBetweenDates_HappyPath(t *testing.T) {
	mockDb, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDb.Close()

	dialector := mysql.New(mysql.Config{
		Conn:                      mockDb,
		DriverName:                "mysql",
		SkipInitializeWithVersion: true,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("gorm open error '%s'", err)
	}

	query := `
	SELECT ec.meter_id AS meter_id, date(ec.date) AS date,
		SUM(ec.active_energy) AS active_energy,
		SUM(ec.reactive_energy) AS reactive_energy,
		SUM(ec.capacitive_reactive) AS capacitive_reactive,
		SUM(ec.solar) AS solar
	FROM energy_consumption ec
	WHERE ec.meter_id in (?,?)
	AND ec.date BETWEEN ? AND ?
	GROUP BY meter_id, date(ec.date)
	ORDER BY meter_id, date(ec.date) ASC`

	rows := sqlmock.
		NewRows([]string{"meter_id", "date", "active_energy", "reactive_energy", "capacitive_reactive", "solar"}).
		AddRow(1, time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local), 30021.21, 525515.36, 5874.488, 987.1477).
		AddRow(1, time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local), 30021.21, 525515.36, 5874.488, 987.1477).
		AddRow(2, time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local), 30021.21, 525515.36, 5874.488, 987.1477).
		AddRow(2, time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local), 30021.21, 525515.36, 5874.488, 987.1477)

	metersIDs := []int{1, 2}
	startDate := "2022-01-01 00:00:01"
	endDate := "2023-12-31 23:59:59"

	mock.ExpectQuery(query).WithArgs(metersIDs[0], metersIDs[1], startDate, endDate).WillReturnRows(rows)

	biaRepository := repository.NewMySQLRepository(db)

	expectedResponse := `[{"MeterID":1,"ActiveEnergy":"30021.21","ReactiveEnergy":"525515.36","CapacitiveReactive":"5874.488","Solar":"987.1477","Date":"2021-08-15T14:30:45.0000001-05:00"},{"MeterID":1,"ActiveEnergy":"30021.21","ReactiveEnergy":"525515.36","CapacitiveReactive":"5874.488","Solar":"987.1477","Date":"2021-08-15T14:30:45.0000001-05:00"},{"MeterID":2,"ActiveEnergy":"30021.21","ReactiveEnergy":"525515.36","CapacitiveReactive":"5874.488","Solar":"987.1477","Date":"2021-08-15T14:30:45.0000001-05:00"},{"MeterID":2,"ActiveEnergy":"30021.21","ReactiveEnergy":"525515.36","CapacitiveReactive":"5874.488","Solar":"987.1477","Date":"2021-08-15T14:30:45.0000001-05:00"}]`
	repositoryResponse, err := biaRepository.GetDailyConsumptionByMetersIdsAndBetweenDates(metersIDs, startDate, endDate)
	currentResponse, _ := json.Marshal(repositoryResponse)
	assert.Nil(t, err)
	assert.NotNil(t, repositoryResponse)
	assert.Equal(t, string(currentResponse), expectedResponse)
}

func TestGetDailyConsumptionByMetersIdsAndBetweenDates_RepositoryError(t *testing.T) {
	mockDb, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDb.Close()

	dialector := mysql.New(mysql.Config{
		Conn:                      mockDb,
		DriverName:                "mysql",
		SkipInitializeWithVersion: true,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("gorm open error '%s'", err)
	}

	query := `
	SELECT ec.meter_id AS meter_id, date(ec.date) AS date,
		SUM(ec.active_energy) AS active_energy,
		SUM(ec.reactive_energy) AS reactive_energy,
		SUM(ec.capacitive_reactive) AS capacitive_reactive,
		SUM(ec.solar) AS solar
	FROM energy_consumption ec
	WHERE ec.meter_id in (?,?)
	AND ec.date BETWEEN ? AND ?
	GROUP BY meter_id, date(ec.date)
	ORDER BY meter_id, date(ec.date) ASC`

	metersIDs := []int{1, 2}
	startDate := "2022-01-01 00:00:01"
	endDate := "2023-12-31 23:59:59"

	someError := errors.New("some error")
	mock.ExpectQuery(query).WithArgs(metersIDs[0], metersIDs[1], startDate, endDate).WillReturnError(someError)

	biaRepository := repository.NewMySQLRepository(db)

	repositoryResponse, err := biaRepository.GetDailyConsumptionByMetersIdsAndBetweenDates(metersIDs, startDate, endDate)
	assert.NotNil(t, err)
	assert.Nil(t, repositoryResponse)
	assert.EqualError(t, someError, err.Error())
}

func TestGetWeeklyConsumptionByMetersIdsAndBetweenDates_HappyPath(t *testing.T) {
	mockDb, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDb.Close()

	dialector := mysql.New(mysql.Config{
		Conn:                      mockDb,
		DriverName:                "mysql",
		SkipInitializeWithVersion: true,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("gorm open error '%s'", err)
	}

	query := `
	SELECT ec.meter_id AS meter_id,
		CONCAT(date(min(ec.date)), ' - ', date(max(ec.date))) AS week,
		SUM(ec.active_energy) AS active_energy,
		SUM(ec.reactive_energy) AS reactive_energy,
		SUM(ec.capacitive_reactive) AS capacitive_reactive,
		SUM(ec.solar) AS solar
	FROM energy_consumption ec
	WHERE ec.meter_id in (?,?)
	AND ec.date BETWEEN ? AND ?
	GROUP BY meter_id, YEARWEEK(ec.date, 1)
	ORDER BY meter_id, YEARWEEK(ec.date, 1) ASC`

	rows := sqlmock.
		NewRows([]string{"meter_id", "week", "active_energy", "reactive_energy", "capacitive_reactive", "solar"}).
		AddRow(1, "2023-06-01 - 2023-06-04", 30021.21, 525515.36, 5874.488, 987.1477).
		AddRow(1, "2023-06-05 - 2023-06-11", 30021.21, 525515.36, 5874.488, 987.1477).
		AddRow(2, "2023-06-01 - 2023-06-04", 30021.21, 525515.36, 5874.488, 987.1477).
		AddRow(2, "2023-06-05 - 2023-06-11", 30021.21, 525515.36, 5874.488, 987.1477)

	metersIDs := []int{1, 2}
	startDate := "2022-01-01 00:00:01"
	endDate := "2023-12-31 23:59:59"

	mock.ExpectQuery(query).WithArgs(metersIDs[0], metersIDs[1], startDate, endDate).WillReturnRows(rows)

	biaRepository := repository.NewMySQLRepository(db)

	expectedResponse := `[{"MeterID":1,"ActiveEnergy":"30021.21","ReactiveEnergy":"525515.36","CapacitiveReactive":"5874.488","Solar":"987.1477","Week":"2023-06-01 - 2023-06-04"},{"MeterID":1,"ActiveEnergy":"30021.21","ReactiveEnergy":"525515.36","CapacitiveReactive":"5874.488","Solar":"987.1477","Week":"2023-06-05 - 2023-06-11"},{"MeterID":2,"ActiveEnergy":"30021.21","ReactiveEnergy":"525515.36","CapacitiveReactive":"5874.488","Solar":"987.1477","Week":"2023-06-01 - 2023-06-04"},{"MeterID":2,"ActiveEnergy":"30021.21","ReactiveEnergy":"525515.36","CapacitiveReactive":"5874.488","Solar":"987.1477","Week":"2023-06-05 - 2023-06-11"}]`
	repositoryResponse, err := biaRepository.GetWeeklyConsumptionByMetersIdsAndBetweenDates(metersIDs, startDate, endDate)
	currentResponse, _ := json.Marshal(repositoryResponse)
	assert.Nil(t, err)
	assert.NotNil(t, repositoryResponse)
	assert.Equal(t, string(currentResponse), expectedResponse)
}

func TestGetWeeklyConsumptionByMetersIdsAndBetweenDates_RepositoryError(t *testing.T) {
	mockDb, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDb.Close()

	dialector := mysql.New(mysql.Config{
		Conn:                      mockDb,
		DriverName:                "mysql",
		SkipInitializeWithVersion: true,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("gorm open error '%s'", err)
	}

	query := `
	SELECT ec.meter_id AS meter_id,
		CONCAT(date(min(ec.date)), ' - ', date(max(ec.date))) AS week,
		SUM(ec.active_energy) AS active_energy,
		SUM(ec.reactive_energy) AS reactive_energy,
		SUM(ec.capacitive_reactive) AS capacitive_reactive,
		SUM(ec.solar) AS solar
	FROM energy_consumption ec
	WHERE ec.meter_id in (?,?)
	AND ec.date BETWEEN ? AND ?
	GROUP BY meter_id, YEARWEEK(ec.date, 1)
	ORDER BY meter_id, YEARWEEK(ec.date, 1) ASC`

	metersIDs := []int{1, 2}
	startDate := "2022-01-01 00:00:01"
	endDate := "2023-12-31 23:59:59"

	someError := errors.New("some error")
	mock.ExpectQuery(query).WithArgs(metersIDs[0], metersIDs[1], startDate, endDate).WillReturnError(someError)

	biaRepository := repository.NewMySQLRepository(db)

	repositoryResponse, err := biaRepository.GetWeeklyConsumptionByMetersIdsAndBetweenDates(metersIDs, startDate, endDate)
	assert.NotNil(t, err)
	assert.Nil(t, repositoryResponse)
	assert.EqualError(t, someError, err.Error())
}

func TestGetMonthlyConsumptionByMetersIdsAndBetweenDates_HappyPath(t *testing.T) {
	mockDb, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDb.Close()

	dialector := mysql.New(mysql.Config{
		Conn:                      mockDb,
		DriverName:                "mysql",
		SkipInitializeWithVersion: true,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("gorm open error '%s'", err)
	}

	query := `
	SELECT ec.meter_id AS meter_id,
		UPPER(CONCAT(MONTHNAME(ec.date), ' ', YEAR(ec.date))) AS month_year,
		SUM(ec.active_energy) AS active_energy,
		SUM(ec.reactive_energy) AS reactive_energy,
		SUM(ec.capacitive_reactive) AS capacitive_reactive,
		SUM(ec.solar) AS solar
	FROM energy_consumption ec
	WHERE ec.meter_id in (?,?)
	AND ec.date BETWEEN ? AND ?
	GROUP BY meter_id, month_year
	ORDER BY meter_id, month_year ASC`

	rows := sqlmock.
		NewRows([]string{"meter_id", "month_year", "active_energy", "reactive_energy", "capacitive_reactive", "solar"}).
		AddRow(1, "APRIL 2023", 30021.21, 525515.36, 5874.488, 987.1477).
		AddRow(1, "AUGUST 2022", 30021.21, 525515.36, 5874.488, 987.1477).
		AddRow(2, "APRIL 2023", 30021.21, 525515.36, 5874.488, 987.1477).
		AddRow(2, "AUGUST 2022", 30021.21, 525515.36, 5874.488, 987.1477)

	metersIDs := []int{1, 2}
	startDate := "2022-01-01 00:00:01"
	endDate := "2023-12-31 23:59:59"

	mock.ExpectQuery(query).WithArgs(metersIDs[0], metersIDs[1], startDate, endDate).WillReturnRows(rows)

	biaRepository := repository.NewMySQLRepository(db)

	expectedResponse := `[{"MeterID":1,"ActiveEnergy":"30021.21","ReactiveEnergy":"525515.36","CapacitiveReactive":"5874.488","Solar":"987.1477","MonthYear":"APRIL 2023"},{"MeterID":1,"ActiveEnergy":"30021.21","ReactiveEnergy":"525515.36","CapacitiveReactive":"5874.488","Solar":"987.1477","MonthYear":"AUGUST 2022"},{"MeterID":2,"ActiveEnergy":"30021.21","ReactiveEnergy":"525515.36","CapacitiveReactive":"5874.488","Solar":"987.1477","MonthYear":"APRIL 2023"},{"MeterID":2,"ActiveEnergy":"30021.21","ReactiveEnergy":"525515.36","CapacitiveReactive":"5874.488","Solar":"987.1477","MonthYear":"AUGUST 2022"}]`
	repositoryResponse, err := biaRepository.GetMonthlyConsumptionByMetersIdsAndBetweenDates(metersIDs, startDate, endDate)
	currentResponse, _ := json.Marshal(repositoryResponse)
	assert.Nil(t, err)
	assert.NotNil(t, repositoryResponse)
	assert.Equal(t, string(currentResponse), expectedResponse)
}

func TestGetMonthlyConsumptionByMetersIdsAndBetweenDates_RepositoryError(t *testing.T) {
	mockDb, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDb.Close()

	dialector := mysql.New(mysql.Config{
		Conn:                      mockDb,
		DriverName:                "mysql",
		SkipInitializeWithVersion: true,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("gorm open error '%s'", err)
	}

	query := `
	SELECT ec.meter_id AS meter_id,
		UPPER(CONCAT(MONTHNAME(ec.date), ' ', YEAR(ec.date))) AS month_year,
		SUM(ec.active_energy) AS active_energy,
		SUM(ec.reactive_energy) AS reactive_energy,
		SUM(ec.capacitive_reactive) AS capacitive_reactive,
		SUM(ec.solar) AS solar
	FROM energy_consumption ec
	WHERE ec.meter_id in (?,?)
	AND ec.date BETWEEN ? AND ?
	GROUP BY meter_id, month_year
	ORDER BY meter_id, month_year ASC`

	metersIDs := []int{1, 2}
	startDate := "2022-01-01 00:00:01"
	endDate := "2023-12-31 23:59:59"

	someError := errors.New("some error")
	mock.ExpectQuery(query).WithArgs(metersIDs[0], metersIDs[1], startDate, endDate).WillReturnError(someError)

	biaRepository := repository.NewMySQLRepository(db)

	repositoryResponse, err := biaRepository.GetMonthlyConsumptionByMetersIdsAndBetweenDates(metersIDs, startDate, endDate)
	assert.NotNil(t, err)
	assert.Nil(t, repositoryResponse)
	assert.EqualError(t, someError, err.Error())
}
