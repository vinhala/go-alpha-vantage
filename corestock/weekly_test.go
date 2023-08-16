package corestock

import (
	"fmt"
	"testing"

	"github.com/ga42quy/go-alpha-vantage/connection"
)

func TestWeekly(t *testing.T) {
	connection := connection.NewAlphaVantageConnection()
	request := WeeklyRequest{
		Symbol: "IBM",
	}
	resp, err := connection.Request(&request)
	if err != nil {
		t.Error(err)
		return
	}
	weeklyResp := resp.(*WeeklyResponse)
	if weeklyResp.Request.Symbol != "IBM" {
		t.Error("Wrong symbol")
		return
	}
	if len(weeklyResp.Timeseries) == 0 {
		t.Error("No data")
		return
	}
	for key, price := range weeklyResp.Timeseries {
		fmt.Printf("Time: %s\n", key.Format("2006-01-02"))
		fmt.Printf("Open: %f\n", price.Open)
		fmt.Printf("High: %f\n", price.High)
		fmt.Printf("Low: %f\n", price.Low)
		fmt.Printf("Close: %f\n", price.Close)
		fmt.Printf("Volume: %f\n", price.Volume)
		break
	}
}
