package corestock

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/ga42quy/go-alpha-vantage/connection"
)

type GlobalQuote struct {
	Symbol           string
	Open             float64
	High             float64
	Low              float64
	Price            float64
	Volume           int64
	LatestTradingDay time.Time
	PreviousClose    float64
	Change           float64
	ChangePercent    float64
}

type globalQuoteRaw struct {
	Symbol           string `json:"01. symbol"`
	Open             string `json:"02. open"`
	High             string `json:"03. high"`
	Low              string `json:"04. low"`
	Price            string `json:"05. price"`
	Volume           string `json:"06. volume"`
	LatestTradingDay string `json:"07. latest trading day"`
	PreviousClose    string `json:"08. previous close"`
	Change           string `json:"09. change"`
	ChangePercent    string `json:"10. change percent"`
}

type globalQuoteResponsePayload struct {
	QuoteRaw globalQuoteRaw `json:"Global Quote"`
}

func parseGlobalQuoteJSON(body io.Reader) (GlobalQuote, error) {
	var payload globalQuoteResponsePayload
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&payload)
	if err != nil {
		return GlobalQuote{}, fmt.Errorf("Failed to parse Global Quote JSON: %w", err)
	}
	//Convert fields
	globalQuote := GlobalQuote{}
	globalQuote.Symbol = payload.QuoteRaw.Symbol
	open, err := strconv.ParseFloat(payload.QuoteRaw.Open, 64)
	if err != nil {
		return GlobalQuote{}, fmt.Errorf("Failed to parse open: %w", err)
	}
	globalQuote.Open = open

	high, err := strconv.ParseFloat(payload.QuoteRaw.High, 64)
	if err != nil {
		return GlobalQuote{}, fmt.Errorf("Failed to parse high: %w", err)
	}
	globalQuote.High = high

	low, err := strconv.ParseFloat(payload.QuoteRaw.Low, 64)
	if err != nil {
		return GlobalQuote{}, fmt.Errorf("Failed to parse low: %w", err)
	}
	globalQuote.Low = low

	price, err := strconv.ParseFloat(payload.QuoteRaw.Price, 64)
	if err != nil {
		return GlobalQuote{}, fmt.Errorf("Failed to parse price: %w", err)
	}
	globalQuote.Price = price

	volume, err := strconv.ParseInt(payload.QuoteRaw.Volume, 10, 64)
	if err != nil {
		return GlobalQuote{}, fmt.Errorf("Failed to parse volume: %w", err)
	}
	globalQuote.Volume = volume

	latestTradingDay, err := time.Parse(time.DateOnly, payload.QuoteRaw.LatestTradingDay)
	if err != nil {
		return GlobalQuote{}, fmt.Errorf("Failed to parse latestTradingDay: %w", err)
	}
	globalQuote.LatestTradingDay = latestTradingDay

	prevClose, err := strconv.ParseFloat(payload.QuoteRaw.PreviousClose, 64)
	if err != nil {
		return GlobalQuote{}, fmt.Errorf("Failed to parse previous close: %w", err)
	}
	globalQuote.PreviousClose = prevClose

	change, err := strconv.ParseFloat(payload.QuoteRaw.Change, 64)
	if err != nil {
		return GlobalQuote{}, fmt.Errorf("Failed to parse change: %w", err)
	}
	globalQuote.Change = change

	changePct, err := strconv.ParseFloat(payload.QuoteRaw.ChangePercent[:len(payload.QuoteRaw.ChangePercent)-1], 64)
	if err != nil {
		return GlobalQuote{}, fmt.Errorf("Failed to parse change percent: %w", err)
	}
	globalQuote.ChangePercent = changePct

	return globalQuote, nil
}

// Request for a single latest stock quote
type GlobalQuoteRequest struct {
	Symbol string
}

// Response of a GlobalQuoteRequest
type GlobalQuoteResponse struct {
	Request *GlobalQuoteRequest
	Quote   GlobalQuote
}

func (r *GlobalQuoteRequest) QueryFunction() connection.QueryFunction {
	return connection.GLOBAL_QUOTE
}

func (r *GlobalQuoteRequest) QueryParams() map[string]string {
	params := map[string]string{
		"symbol":   r.Symbol,
		"datatype": "json",
	}

	return params
}

func (r *GlobalQuoteRequest) ParseResponse(response *http.Response) (interface{}, error) {
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Error during Global Quote Request. Status %v", response.StatusCode)
	}
	quote, err := parseGlobalQuoteJSON(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse Global Quote response: %w", err)
	}
	return &GlobalQuoteResponse{
		Request: r,
		Quote:   quote,
	}, nil
}
