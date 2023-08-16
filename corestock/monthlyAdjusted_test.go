package corestock

import (
	"fmt"
	"testing"

	"github.com/ga42quy/go-alpha-vantage/connection"
)

func TestMonthlyAdjusted(t *testing.T) {
	connection := connection.NewAlphaVantageConnection()
	request := MonthlyAdjustedRequest{
		Symbol: "IBM",
	}
	resp, err := connection.Request(&request)
	if err != nil {
		t.Error(err)
		return
	}
	monthlyAdjResp := resp.(*MonthlyAdjustedResponse)
	if monthlyAdjResp.Request.Symbol != "IBM" {
		t.Error("Wrong symbol")
		return
	}
	if len(monthlyAdjResp.Timeseries) == 0 {
		t.Error("No data")
		return
	}
	for key, price := range monthlyAdjResp.Timeseries {
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
