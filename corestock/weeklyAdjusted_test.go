package corestock

import (
	"fmt"
	"testing"

	"github.com/ga42quy/go-alpha-vantage/connection"
)

func TestWeeklyAdjusted(t *testing.T) {
	t.Skip("Skip weekly adjusted test")
	connection := connection.NewAlphaVantageConnection()
	request := WeeklyAdjustedRequest{
		Symbol: "IBM",
	}
	resp, err := connection.Request(&request)
	if err != nil {
		t.Error(err)
		return
	}
	weeklyAdjResp := resp.(*WeeklyAdjustedResponse)
	if weeklyAdjResp.Request.Symbol != "IBM" {
		t.Error("Wrong symbol")
		return
	}
	if len(weeklyAdjResp.Timeseries) == 0 {
		t.Error("No data")
		return
	}
	for key, price := range weeklyAdjResp.Timeseries {
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
