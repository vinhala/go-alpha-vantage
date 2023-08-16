package corestock

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ga42quy/go-alpha-vantage/connection"
)

// Request for weekly stock data
type WeeklyRequest struct {
	Symbol string
}

// Response of a WeeklyRequest
type WeeklyResponse struct {
	Request    *WeeklyRequest
	Timeseries map[time.Time]*OHLCV
}

func (r *WeeklyRequest) QueryFunction() connection.QueryFunction {
	return connection.TIME_SERIES_WEEKLY
}

func (r *WeeklyRequest) QueryParams() map[string]string {
	params := map[string]string{
		"symbol":   r.Symbol,
		"datatype": "csv",
	}

	return params
}

func (r *WeeklyRequest) ParseResponse(response *http.Response) (interface{}, error) {
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Error during Weekly Request. Status %v", response.StatusCode)
	}
	series, err := parseTimeSeriesCSV(response.Body, time.DateOnly)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse Weekly response: %w", err)
	}
	return &WeeklyResponse{
		Request:    r,
		Timeseries: series,
	}, nil
}
