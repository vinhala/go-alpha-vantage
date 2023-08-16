package corestock

import (
	"fmt"
	"testing"

	"github.com/ga42quy/go-alpha-vantage/connection"
)

func TestMonthly(t *testing.T) {
	connection := connection.NewAlphaVantageConnection()
	request := MonthlyRequest{
		Symbol: "IBM",
	}
	resp, err := connection.Request(&request)
	if err != nil {
		t.Error(err)
		return
	}
	monthlyResp := resp.(*MonthlyResponse)
	if monthlyResp.Request.Symbol != "IBM" {
		t.Error("Wrong symbol")
		return
	}
	if len(monthlyResp.Timeseries) == 0 {
		t.Error("No data")
		return
	}
	for key, price := range monthlyResp.Timeseries {
		fmt.Printf("Time: %s\n", key.Format("2006-01-02"))
		fmt.Printf("Open: %f\n", price.Open)
		fmt.Printf("High: %f\n", price.High)
		fmt.Printf("Low: %f\n", price.Low)
		fmt.Printf("Close: %f\n", price.Close)
		fmt.Printf("Volume: %f\n", price.Volume)
		break
	}
}
