package handler_test

import (
	"bia-challenge/internal/adapters/handler"
	"bia-challenge/internal/core/domain"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type biaServiceMockPort struct {
	mock.Mock
}

func (m *biaServiceMockPort) GetEnergyConsumption(
	request *domain.GetEnergyConsumptionRequest) (*domain.GetEnergyConsumptionResponse, error) {
	ret := m.Called(request)

	var r0 *domain.GetEnergyConsumptionResponse
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*domain.GetEnergyConsumptionResponse)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func TestGetEnergyConsumptionWeekly_HappyPath(t *testing.T) {
	biaServiceMock := new(biaServiceMockPort)

	request := &domain.GetEnergyConsumptionRequest{
		MetersIDs:  []int{1},
		StartDate:  "2022-01-01" + " 00:00:00",
		EndDate:    "2023-12-31" + " 23:59:59",
		KindPeriod: domain.Weekly,
	}

	response := &domain.GetEnergyConsumptionResponse{
		DataGraph: []domain.DataGraph{{
			MeterID:            1,
			Address:            "some address",
			Active:             []decimal.Decimal{decimal.Zero},
			ReactiveInductive:  []decimal.Decimal{decimal.Zero},
			ReactiveCapacitive: []decimal.Decimal{decimal.Zero},
			Exported:           []decimal.Decimal{decimal.Zero},
			Period:             []string{"1"},
		}},
	}

	biaServiceMock.On("GetEnergyConsumption", request).Return(response, nil)

	httpHandler := handler.NewHttpHandler(biaServiceMock)

	req, err := http.NewRequest("GET", "/consumption", nil)
	if err != nil {
		t.Fatalf("fail to create request: %s", err.Error())
	}

	q := req.URL.Query()
	q.Add("start_date", "2022-01-01")
	q.Add("end_date", "2023-12-31")
	q.Add("kind_period", "weekly")
	q.Add("meters_ids", "1")
	req.URL.RawQuery = q.Encode()

	resp := httptest.NewRecorder()
	httpHandler.GetEnergyConsumption(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	currentResponseJson, _ := json.Marshal(response)
	assert.JSONEq(t, string(currentResponseJson), resp.Body.String())
}

func TestGetEnergyConsumption_ErrorService(t *testing.T) {
	biaServiceMock := new(biaServiceMockPort)

	request := &domain.GetEnergyConsumptionRequest{
		MetersIDs:  []int{1},
		StartDate:  "2022-01-01" + " 00:00:00",
		EndDate:    "2023-12-31" + " 23:59:59",
		KindPeriod: domain.Weekly,
	}

	someError := errors.New("some error")
	biaServiceMock.On("GetEnergyConsumption", request).Return(nil, someError)

	httpHandler := handler.NewHttpHandler(biaServiceMock)

	req, err := http.NewRequest("GET", "/consumption", nil)
	if err != nil {
		t.Fatalf("fail to create request: %s", err.Error())
	}

	q := req.URL.Query()
	q.Add("start_date", "2022-01-01")
	q.Add("end_date", "2023-12-31")
	q.Add("kind_period", "weekly")
	q.Add("meters_ids", "1")
	req.URL.RawQuery = q.Encode()

	resp := httptest.NewRecorder()
	httpHandler.GetEnergyConsumption(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	assert.Equal(t, someError.Error(), strings.TrimSuffix(resp.Body.String(), "\n"))
}

func TestGetEnergyConsumption_RequestValidationErrors(t *testing.T) {
	biaServiceMock := new(biaServiceMockPort)
	httpHandler := handler.NewHttpHandler(biaServiceMock)

	type args struct {
		req *http.Request
	}

	tests := []struct {
		name     string
		args     func(t *testing.T) args
		wantCode int
		wantBody string
	}{
		{
			name: "must return http.StatusBadRequest to invalid start date",
			args: func(*testing.T) args {
				req, err := http.NewRequest("GET", "/consumption", nil)
				if err != nil {
					t.Fatalf("fail to create request: %s", err.Error())
				}

				q := req.URL.Query()
				q.Add("start_date", "01-01-2022")
				q.Add("end_date", "2023-12-31")
				q.Add("kind_period", "weekly")
				q.Add("meters_ids", "1")
				req.URL.RawQuery = q.Encode()

				return args{
					req: req,
				}
			},
			wantCode: http.StatusBadRequest,
			wantBody: "{\"status\":\"Invalid request.\",\"error\":\"invalid start date format\"}",
		},
		{
			name: "must return http.StatusBadRequest to invalid end date",
			args: func(*testing.T) args {
				req, err := http.NewRequest("GET", "/consumption", nil)
				if err != nil {
					t.Fatalf("fail to create request: %s", err.Error())
				}

				q := req.URL.Query()
				q.Add("start_date", "2022-06-01")
				q.Add("end_date", "01-06-2023")
				q.Add("kind_period", "weekly")
				q.Add("meters_ids", "1")
				req.URL.RawQuery = q.Encode()

				return args{
					req: req,
				}
			},
			wantCode: http.StatusBadRequest,
			wantBody: "{\"status\":\"Invalid request.\",\"error\":\"invalid end date format\"}",
		},
		{
			name: "must return http.StatusBadRequest to invalid start and end date",
			args: func(*testing.T) args {
				req, err := http.NewRequest("GET", "/consumption", nil)
				if err != nil {
					t.Fatalf("fail to create request: %s", err.Error())
				}

				q := req.URL.Query()
				q.Add("start_date", "02-05-9856")
				q.Add("end_date", "01-06-2023")
				q.Add("kind_period", "weekly")
				q.Add("meters_ids", "1")
				req.URL.RawQuery = q.Encode()

				return args{
					req: req,
				}
			},
			wantCode: http.StatusBadRequest,
			wantBody: "{\"status\":\"Invalid request.\",\"error\":\"invalid start date format : invalid end date format\"}",
		},
		{
			name: "must return http.StatusBadRequest to required meters ids",
			args: func(*testing.T) args {
				req, err := http.NewRequest("GET", "/consumption", nil)
				if err != nil {
					t.Fatalf("fail to create request: %s", err.Error())
				}

				q := req.URL.Query()
				q.Add("start_date", "02-05-9856")
				q.Add("end_date", "01-06-2023")
				q.Add("kind_period", "weekly")
				req.URL.RawQuery = q.Encode()

				return args{
					req: req,
				}
			},
			wantCode: http.StatusBadRequest,
			wantBody: "{\"status\":\"Invalid request.\",\"error\":\"meters_ids are required\"}",
		},
		{
			name: "must return http.StatusBadRequest to invalid meters ids",
			args: func(*testing.T) args {
				req, err := http.NewRequest("GET", "/consumption", nil)
				if err != nil {
					t.Fatalf("fail to create request: %s", err.Error())
				}

				q := req.URL.Query()
				q.Add("start_date", "2023-06-01")
				q.Add("end_date", "2023-07-01")
				q.Add("kind_period", "weekly")
				q.Add("meters_ids", "a")
				req.URL.RawQuery = q.Encode()

				return args{
					req: req,
				}
			},
			wantCode: http.StatusBadRequest,
			wantBody: "{\"status\":\"Invalid request.\",\"error\":\"meters_ids are invalid\"}",
		},
		{
			name: "must return http.StatusBadRequest to required  kind period",
			args: func(*testing.T) args {
				req, err := http.NewRequest("GET", "/consumption", nil)
				if err != nil {
					t.Fatalf("fail to create request: %s", err.Error())
				}

				q := req.URL.Query()
				q.Add("start_date", "02-05-9856")
				q.Add("end_date", "01-06-2023")
				q.Add("meters_ids", "1")
				req.URL.RawQuery = q.Encode()

				return args{
					req: req,
				}
			},
			wantCode: http.StatusBadRequest,
			wantBody: "{\"status\":\"Invalid request.\",\"error\":\"kind_period is required\"}",
		},
		{
			name: "must return http.StatusBadRequest to invalid  kind period",
			args: func(*testing.T) args {
				req, err := http.NewRequest("GET", "/consumption", nil)
				if err != nil {
					t.Fatalf("fail to create request: %s", err.Error())
				}

				q := req.URL.Query()
				q.Add("start_date", "2023-06-01")
				q.Add("end_date", "2023-07-01")
				q.Add("meters_ids", "1")
				q.Add("kind_period", "cxxxx")
				req.URL.RawQuery = q.Encode()

				return args{
					req: req,
				}
			},
			wantCode: http.StatusBadRequest,
			wantBody: "{\"status\":\"Invalid request.\",\"error\":\"kind_period is invalid\"}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)
			resp := httptest.NewRecorder()
			httpHandler.GetEnergyConsumption(resp, tArgs.req)

			assert.Equal(t, tt.wantCode, resp.Code)
			assert.Equal(t, tt.wantBody, strings.TrimSuffix(resp.Body.String(), "\n"))
		})
	}

}
