package corestock

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ga42quy/go-alpha-vantage/connection"
)

// Request for monthly adjusted stock data
type MonthlyAdjustedRequest struct {
	Symbol string
}

// Response of a MonthlyAdjustedRequest
type MonthlyAdjustedResponse struct {
	Request    *MonthlyAdjustedRequest
	Timeseries map[time.Time]*OHLCVAdjusted
}

func (r *MonthlyAdjustedRequest) QueryFunction() connection.QueryFunction {
	return connection.TIME_SERIES_MONTHLY_ADJUSTED
}

func (r *MonthlyAdjustedRequest) QueryParams() map[string]string {
	params := map[string]string{
		"symbol":   r.Symbol,
		"datatype": "csv",
	}

	return params
}

func (r *MonthlyAdjustedRequest) ParseResponse(response *http.Response) (interface{}, error) {
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Error during Monthly Adjusted Request. Status %v", response.StatusCode)
	}
	series, err := parseTimeSeriesAdjustedCSV(response.Body, time.DateOnly, true)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse Monthly Adjusted response: %w", err)
	}
	return &MonthlyAdjustedResponse{
		Request:    r,
		Timeseries: series,
	}, nil
}
