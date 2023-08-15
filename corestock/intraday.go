package corestock

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ga42quy/go-alpha-vantage/connection"
)

// OHLCV is the struct for the open, high, low, close, and volume of a financial instrument
type OHLCV struct {
	Open   float64 `json:"1. open"`
	High   float64 `json:"2. high"`
	Low    float64 `json:"3. low"`
	Close  float64 `json:"4. close"`
	Volume float64 `json:"5. volume"`
}

// IntradayInterval is the interval between two consecutive data points in the time series.
type IntradayInterval string

const (
	ONE_MIN     IntradayInterval = "1min"
	FIVE_MIN    IntradayInterval = "5min"
	FIFTEEN_MIN IntradayInterval = "15min"
	THIRTY_MIN  IntradayInterval = "30min"
	SIXTY_MIN   IntradayInterval = "60min"
)

// Request for intraday stock price time series.
type IntradayRequest struct {
	Interval          IntradayInterval
	Symbol            string
	NotAdjusted       bool
	SkipExtendedHours bool
	//Optional Month in the format YYYY-MM
	Month      string
	OutputSize OutputSize
}

// Response of a IntradayRequest
type IntradayResponse struct {
	Request    *IntradayRequest
	Timeseries map[time.Time]*OHLCV
}

func (r *IntradayRequest) QueryFunction() connection.QueryFunction {
	return connection.TIME_SERIES_INTRADAY
}

func (r *IntradayRequest) QueryParams() map[string]string {
	params := map[string]string{
		"interval": string(r.Interval),
		"symbol":   r.Symbol,
		"datatype": "csv",
	}
	if r.NotAdjusted {
		params["adjusted"] = "false"
	}
	if r.SkipExtendedHours {
		params["extended_hours"] = "false"
	}
	if r.Month != "" {
		params["month"] = r.Month
	}
	if r.OutputSize != "" {
		params["outputsize"] = string(r.OutputSize)
	}

	return params
}

func (r *IntradayRequest) ParseResponse(response *http.Response) (interface{}, error) {
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Error during Intraday Request. Status %v", response.StatusCode)
	}
	series, err := parseTimeSeriesCSV(response.Body, time.DateTime)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse Intraday response: %w", err)
	}
	return &IntradayResponse{
		Request:    r,
		Timeseries: series,
	}, nil
}
