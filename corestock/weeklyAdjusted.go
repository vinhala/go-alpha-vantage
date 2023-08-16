package corestock

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ga42quy/go-alpha-vantage/connection"
)

// Request for weekly adjusted stock data
type WeeklyAdjustedRequest struct {
	Symbol string
}

// Response of a WeeklyAdjustedRequest
type WeeklyAdjustedResponse struct {
	Request    *WeeklyAdjustedRequest
	Timeseries map[time.Time]*OHLCVAdjusted
}

func (r *WeeklyAdjustedRequest) QueryFunction() connection.QueryFunction {
	return connection.TIME_SERIES_WEEKLY_ADJUSTED
}

func (r *WeeklyAdjustedRequest) QueryParams() map[string]string {
	params := map[string]string{
		"symbol":   r.Symbol,
		"datatype": "csv",
	}

	return params
}

func (r *WeeklyAdjustedRequest) ParseResponse(response *http.Response) (interface{}, error) {
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Error during Weekly Adjusted Request. Status %v", response.StatusCode)
	}
	series, err := parseTimeSeriesAdjustedCSV(response.Body, time.DateOnly, true)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse Weekly Adjusted response: %w", err)
	}
	return &WeeklyAdjustedResponse{
		Request:    r,
		Timeseries: series,
	}, nil
}
