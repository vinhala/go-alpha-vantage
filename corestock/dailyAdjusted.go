package corestock

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ga42quy/go-alpha-vantage/connection"
)

// Request for daily stock data
type DailyAdjustedRequest struct {
	Symbol     string
	OutputSize OutputSize
}

// Response of a DailyRequest
type DailyAdjustedResponse struct {
	Request    *DailyAdjustedRequest
	Timeseries map[time.Time]*OHLCVAdjusted
}

func (r *DailyAdjustedRequest) QueryFunction() connection.QueryFunction {
	return connection.TIME_SERIES_DAILY_ADJUSTED
}

func (r *DailyAdjustedRequest) QueryParams() map[string]string {
	params := map[string]string{
		"symbol":   r.Symbol,
		"datatype": "csv",
	}
	if r.OutputSize != "" {
		params["outputsize"] = string(r.OutputSize)
	} else {
		params["outputsize"] = string(COMPACT)
	}

	return params
}

func (r *DailyAdjustedRequest) ParseResponse(response *http.Response) (interface{}, error) {
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Error during Daily Adjusted Request. Status %v", response.StatusCode)
	}
	series, err := parseTimeSeriesAdjustedCSV(response.Body, time.DateOnly, false)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse Daily Adjusted response: %w", err)
	}
	return &DailyAdjustedResponse{
		Request:    r,
		Timeseries: series,
	}, nil
}
