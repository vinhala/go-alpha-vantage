package corestock

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ga42quy/go-alpha-vantage/connection"
)

// Request for Monthly stock data
type MonthlyRequest struct {
	Symbol string
}

// Response of a MonthlyRequest
type MonthlyResponse struct {
	Request    *MonthlyRequest
	Timeseries map[time.Time]*OHLCV
}

func (r *MonthlyRequest) QueryFunction() connection.QueryFunction {
	return connection.TIME_SERIES_MONTHLY
}

func (r *MonthlyRequest) QueryParams() map[string]string {
	params := map[string]string{
		"symbol":   r.Symbol,
		"datatype": "csv",
	}

	return params
}

func (r *MonthlyRequest) ParseResponse(response *http.Response) (interface{}, error) {
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Error during Monthly Request. Status %v", response.StatusCode)
	}
	series, err := parseTimeSeriesCSV(response.Body, time.DateOnly)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse Monthly response: %w", err)
	}
	return &MonthlyResponse{
		Request:    r,
		Timeseries: series,
	}, nil
}
