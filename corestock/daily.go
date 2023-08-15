package corestock

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ga42quy/go-alpha-vantage/connection"
)

// Request for daily stock data
type DailyRequest struct {
	Symbol     string
	OutputSize OutputSize
}

// Response of a DailyRequest
type DailyResponse struct {
	Request    *DailyRequest
	Timeseries map[time.Time]*OHLCV
}

func (r *DailyRequest) QueryFunction() connection.QueryFunction {
	return connection.TIME_SERIES_DAILY
}

func (r *DailyRequest) QueryParams() map[string]string {
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

func (r *DailyRequest) ParseResponse(response *http.Response) (interface{}, error) {
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Error during Daily Request. Status %v", response.StatusCode)
	}
	series, err := parseTimeSeriesCSV(response.Body, time.DateOnly)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse Daily response: %w", err)
	}
	return &DailyResponse{
		Request:    r,
		Timeseries: series,
	}, nil
}
