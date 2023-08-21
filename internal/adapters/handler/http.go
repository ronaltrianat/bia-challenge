package handler

import (
	"bia-challenge/internal/core/domain"
	"bia-challenge/internal/core/ports"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/go-chi/render"
)

const (
	_ServiceDateFormat = "2006-01-02"
	_StartTimeFormat   = "00:00:00"
	_EndTimeFormat     = "23:59:59"
)

var (
	ErrInvalidStartDateFormat = errors.New("invalid start date format")
	ErrInvalidEndDateFormat   = errors.New("invalid end date format")
	ErrMetersIdsRequired      = errors.New("meters_ids are required")
	ErrInvalidMetersIds       = errors.New("meters_ids are invalid")
	ErrInvalidkindPeriod      = errors.New("kind_period is invalid")
	ErrKindPeriodRequired     = errors.New("kind_period is required")
)

type httpHandler struct {
	biaService ports.BiaServicePort
}

func NewHttpHandler(biaService ports.BiaServicePort) *httpHandler {
	return &httpHandler{biaService: biaService}
}

func (hdl *httpHandler) GetEnergyConsumption(w http.ResponseWriter, r *http.Request) {
	request, err := hdl.buildRequest(r.URL.Query())
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	response, err := hdl.biaService.GetEnergyConsumption(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, response)
}

func (hdl *httpHandler) buildRequest(queryParameters url.Values) (*domain.GetEnergyConsumptionRequest, error) {
	var validationErrors []error

	metersIds := queryParameters["meters_ids"]
	if len(metersIds) == 0 {
		return nil, ErrMetersIdsRequired
	}

	kindPeriod := queryParameters.Get("kind_period")
	if kindPeriod == "" {
		return nil, ErrKindPeriodRequired
	}

	switch domain.KindPeriod(kindPeriod) {
	case domain.Daily, domain.Monthly, domain.Weekly:
	default:
		validationErrors = append(validationErrors, ErrInvalidkindPeriod)
	}

	startDate := queryParameters.Get("start_date")
	endDate := queryParameters.Get("end_date")

	if _, err := time.Parse(_ServiceDateFormat, startDate); err != nil {
		validationErrors = append(validationErrors, ErrInvalidStartDateFormat)
	}

	if _, err := time.Parse(_ServiceDateFormat, endDate); err != nil {
		validationErrors = append(validationErrors, ErrInvalidEndDateFormat)
	}

	ints := make([]int, len(metersIds))

	for i, s := range metersIds {
		if value, err := strconv.Atoi(s); err != nil {
			validationErrors = append(validationErrors, ErrInvalidMetersIds)
		} else {
			ints[i] = value
		}
	}

	request := &domain.GetEnergyConsumptionRequest{
		MetersIDs:  ints,
		StartDate:  startDate + " " + _StartTimeFormat,
		EndDate:    endDate + " " + _EndTimeFormat,
		KindPeriod: domain.KindPeriod(kindPeriod),
	}

	return request, errors.Join(validationErrors...)
}
