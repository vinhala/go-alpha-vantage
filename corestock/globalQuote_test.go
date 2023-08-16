package corestock

import (
	"fmt"
	"testing"

	"github.com/ga42quy/go-alpha-vantage/connection"
)

func TestGlobalQuoteRequest(t *testing.T) {
	connection := connection.NewAlphaVantageConnection()
	request := GlobalQuoteRequest{
		Symbol: "IBM",
	}
	response, err := connection.Request(&request)
	if err != nil {
		t.Errorf("Error during Global Quote Request: %v", err)
		return
	}
	if response == nil {
		t.Errorf("Empty response")
		return
	}
	globalQuoteResponse, ok := response.(*GlobalQuoteResponse)
	if !ok {
		t.Errorf("Wrong response type")
		return
	}
	if globalQuoteResponse.Quote.Symbol != "IBM" {
		t.Errorf("Wrong symbol")
		return
	}
	if globalQuoteResponse.Quote.Open == 0 {
		t.Errorf("No open")
		return
	}
	if globalQuoteResponse.Quote.High == 0 {
		t.Errorf("No high")
		return
	}

	fmt.Printf("Symbol: %s\n", globalQuoteResponse.Quote.Symbol)
	fmt.Printf("Open: %f\n", globalQuoteResponse.Quote.Open)
	fmt.Printf("High: %f\n", globalQuoteResponse.Quote.High)
	fmt.Printf("Low: %f\n", globalQuoteResponse.Quote.Low)
	fmt.Printf("Price: %f\n", globalQuoteResponse.Quote.Price)
	fmt.Printf("Latest Trading Day: %s\n", globalQuoteResponse.Quote.LatestTradingDay)
	fmt.Printf("Previous Close: %f\n", globalQuoteResponse.Quote.PreviousClose)
	fmt.Printf("Change: %f\n", globalQuoteResponse.Quote.Change)
	fmt.Printf("Change Percent: %f\n", globalQuoteResponse.Quote.ChangePercent)
}
