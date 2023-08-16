package corestock

import (
	"fmt"
	"testing"

	"github.com/ga42quy/go-alpha-vantage/connection"
)

func TestDailyCompact(t *testing.T) {
	connection := connection.NewAlphaVantageConnection()
	request := DailyRequest{
		Symbol:     "IBM",
		OutputSize: COMPACT,
	}
	resp, err := connection.Request(&request)
	if err != nil {
		t.Error(err)
		return
	}
	dailyResp := resp.(*DailyResponse)
	if dailyResp.Request.Symbol != "IBM" {
		t.Error("Wrong symbol")
		return
	}
	if dailyResp.Request.OutputSize != COMPACT {
		t.Error("Wrong output size")
		return
	}
	if len(dailyResp.Timeseries) == 0 {
		t.Error("No data")
		return
	}
	for key, price := range dailyResp.Timeseries {
		fmt.Printf("Time: %s\n", key.Format("2006-01-02"))
		fmt.Printf("Open: %f\n", price.Open)
		fmt.Printf("High: %f\n", price.High)
		fmt.Printf("Low: %f\n", price.Low)
		fmt.Printf("Close: %f\n", price.Close)
		fmt.Printf("Volume: %f\n", price.Volume)
		break
	}
}

func TestDailyFULL(t *testing.T) {
	t.Skip("Skip daily FULL test")
	connection := connection.NewAlphaVantageConnection()
	request := DailyRequest{
		Symbol:     "IBM",
		OutputSize: FULL,
	}
	resp, err := connection.Request(&request)
	if err != nil {
		t.Error(err)
		return
	}
	dailyResp := resp.(*DailyResponse)
	if dailyResp.Request.Symbol != "IBM" {
		t.Error("Wrong symbol")
		return
	}
	if dailyResp.Request.OutputSize != FULL {
		t.Error("Wrong output size")
		return
	}
	if len(dailyResp.Timeseries) == 0 {
		t.Error("No data")
		return
	}
	for key, price := range dailyResp.Timeseries {
		fmt.Printf("Time: %s\n", key.Format("2006-01-02"))
		fmt.Printf("Open: %f\n", price.Open)
		fmt.Printf("High: %f\n", price.High)
		fmt.Printf("Low: %f\n", price.Low)
		fmt.Printf("Close: %f\n", price.Close)
		fmt.Printf("Volume: %f\n", price.Volume)
		break
	}
}
