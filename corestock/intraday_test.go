package corestock

import (
	"fmt"
	"testing"
	"time"

	"github.com/ga42quy/go-alpha-vantage/connection"
)

func TestIntraday5MinCompact(t *testing.T) {
	connection := connection.NewAlphaVantageConnection()
	request := IntradayRequest{
		Symbol:     "IBM",
		Interval:   FIVE_MIN,
		OutputSize: COMPACT,
	}
	resp, err := connection.Request(&request)
	if err != nil {
		t.Error(err)
		return
	}
	intradayResp := resp.(*IntradayResponse)
	if intradayResp.Request.Symbol != "IBM" {
		t.Error("Wrong symbol")
		return
	}
	if intradayResp.Request.Interval != FIVE_MIN {
		t.Error("Wrong interval")
		return
	}
	if intradayResp.Request.OutputSize != COMPACT {
		t.Error("Wrong output size")
		return
	}
	if len(intradayResp.Timeseries) == 0 {
		t.Error("No data")
		return
	}
	for key, price := range intradayResp.Timeseries {
		fmt.Printf("Time: %s\n", key.Format("2006-01-02 15:04:05"))
		fmt.Printf("Open: %f\n", price.Open)
		fmt.Printf("High: %f\n", price.High)
		fmt.Printf("Low: %f\n", price.Low)
		fmt.Printf("Close: %f\n", price.Close)
		fmt.Printf("Volume: %f\n", price.Volume)
		break
	}
	//Throttle
	time.Sleep(10 * time.Millisecond)
}

func TestIntraday5MinFull(t *testing.T) {
	t.Skip("Skip intraday full test")
	connection := connection.NewAlphaVantageConnection()
	request := IntradayRequest{
		Symbol:     "IBM",
		Interval:   FIVE_MIN,
		OutputSize: FULL,
	}
	resp, err := connection.Request(&request)
	if err != nil {
		t.Error(err)
		return
	}
	intradayResp := resp.(*IntradayResponse)
	if intradayResp.Request.Symbol != "IBM" {
		t.Error("Wrong symbol")
		return
	}
	if intradayResp.Request.Interval != FIVE_MIN {
		t.Error("Wrong interval")
		return
	}
	if intradayResp.Request.OutputSize != FULL {
		t.Error("Wrong output size")
		return
	}
	if len(intradayResp.Timeseries) == 0 {
		t.Error("No data")
		return
	}
	for key, price := range intradayResp.Timeseries {
		fmt.Printf("Time: %s\n", key.Format("2006-01-02 15:04:05"))
		fmt.Printf("Open: %f\n", price.Open)
		fmt.Printf("High: %f\n", price.High)
		fmt.Printf("Low: %f\n", price.Low)
		fmt.Printf("Close: %f\n", price.Close)
		fmt.Printf("Volume: %f\n", price.Volume)
		break
	}
	//Throttle
	time.Sleep(10 * time.Millisecond)
}

func TestIntraday5MinCustomMonth(t *testing.T) {
	t.Skip("Skip intraday custom month test")
	connection := connection.NewAlphaVantageConnection()
	request := IntradayRequest{
		Symbol:     "IBM",
		Interval:   FIVE_MIN,
		OutputSize: COMPACT,
		Month:      "2020-01",
	}
	resp, err := connection.Request(&request)
	if err != nil {
		t.Error(err)
		return
	}
	intradayResp := resp.(*IntradayResponse)
	if intradayResp.Request.Symbol != "IBM" {
		t.Error("Wrong symbol")
		return
	}
	if intradayResp.Request.Interval != FIVE_MIN {
		t.Error("Wrong interval")
		return
	}
	if intradayResp.Request.OutputSize != COMPACT {
		t.Error("Wrong output size")
		return
	}
	if len(intradayResp.Timeseries) == 0 {
		t.Error("No data")
		return
	}
	for key, price := range intradayResp.Timeseries {
		if key.Month() != 1 {
			t.Error("Wrong month")
			return
		}
		if key.Year() != 2020 {
			t.Error("Wrong year")
			return
		}
		fmt.Printf("Time: %s\n", key.Format("2006-01-02 15:04:05"))
		fmt.Printf("Open: %f\n", price.Open)
		fmt.Printf("High: %f\n", price.High)
		fmt.Printf("Low: %f\n", price.Low)
		fmt.Printf("Close: %f\n", price.Close)
		fmt.Printf("Volume: %f\n", price.Volume)
		break
	}
}
