package corestock

import (
	"fmt"
	"testing"

	"github.com/ga42quy/go-alpha-vantage/connection"
)

func TestDailyAdjustedCompact(t *testing.T) {
	t.Skip("Skip DailyAdjustedCompact because it's a premium endpoint")
	connection := connection.NewAlphaVantageConnection()
	request := DailyAdjustedRequest{
		Symbol:     "IBM",
		OutputSize: COMPACT,
	}
	resp, err := connection.Request(&request)
	if err != nil {
		t.Error(err)
		return
	}
	dailyAdjResp := resp.(*DailyAdjustedResponse)
	if dailyAdjResp.Request.Symbol != "IBM" {
		t.Error("Wrong symbol")
		return
	}
	if dailyAdjResp.Request.OutputSize != COMPACT {
		t.Error("Wrong output size")
		return
	}
	if len(dailyAdjResp.Timeseries) == 0 {
		t.Error("No data")
		return
	}
	for key, price := range dailyAdjResp.Timeseries {
		fmt.Printf("Time: %s\n", key.Format("2006-01-02"))
		fmt.Printf("Open: %f\n", price.Open)
		fmt.Printf("High: %f\n", price.High)
		fmt.Printf("Low: %f\n", price.Low)
		fmt.Printf("Close: %f\n", price.Close)
		fmt.Printf("AdjustedClose: %f\n", price.AdjustedClose)
		fmt.Printf("Volume: %f\n", price.Volume)
		fmt.Printf("DividendAmount: %f\n", price.DividendAmount)
		fmt.Printf("SplitCoefficient: %f\n", price.SplitCoefficient)
		break
	}
}

func TestDailyAdjustedFull(t *testing.T) {
	t.Skip("Skip DailyAdjustedFull because it's a premium endpoint")
	connection := connection.NewAlphaVantageConnection()
	request := DailyAdjustedRequest{
		Symbol:     "IBM",
		OutputSize: FULL,
	}
	resp, err := connection.Request(&request)
	if err != nil {
		t.Error(err)
		return
	}
	dailyAdjResp := resp.(*DailyAdjustedResponse)
	if dailyAdjResp.Request.Symbol != "IBM" {
		t.Error("Wrong symbol")
		return
	}
	if dailyAdjResp.Request.OutputSize != FULL {
		t.Error("Wrong output size")
		return
	}
	if len(dailyAdjResp.Timeseries) == 0 {
		t.Error("No data")
		return
	}
	for key, price := range dailyAdjResp.Timeseries {
		fmt.Printf("Time: %s\n", key.Format("2006-01-02"))
		fmt.Printf("Open: %f\n", price.Open)
		fmt.Printf("High: %f\n", price.High)
		fmt.Printf("Low: %f\n", price.Low)
		fmt.Printf("Close: %f\n", price.Close)
		fmt.Printf("AdjustedClose: %f\n", price.AdjustedClose)
		fmt.Printf("Volume: %f\n", price.Volume)
		fmt.Printf("DividendAmount: %f\n", price.DividendAmount)
		fmt.Printf("SplitCoefficient: %f\n", price.SplitCoefficient)
		break
	}
}
