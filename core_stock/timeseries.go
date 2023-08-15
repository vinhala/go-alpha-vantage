package corestock

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"time"
)

// Columns of csv timeseries
const (
	timestamp = iota
	open
	high
	low
	close
	volume
)

// OHLCV is the struct for the open, high, low, close, and volume of a financial instrument
type OHLCV struct {
	Open   float64 `json:"1. open"`
	High   float64 `json:"2. high"`
	Low    float64 `json:"3. low"`
	Close  float64 `json:"4. close"`
	Volume int64   `json:"5. volume"`
}

// A single entry in a prices time series
type PricesTimeSeriesEntry struct {
	Timestamp time.Time
	PriceData OHLCV
}

func parseCSVRecord(row []string) (*PricesTimeSeriesEntry, error) {
	entry := &PricesTimeSeriesEntry{}
	ohlcv := &OHLCV{}

	t, err := time.Parse(time.DateTime, row[timestamp])
	if err != nil {
		return nil, fmt.Errorf("error parsing timestamp %s: %w", row[timestamp], err)
	}
	entry.Timestamp = t

	o, err := strconv.ParseFloat(row[open], 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing open %s: %w", row[open], err)
	}
	ohlcv.Open = o

	h, err := strconv.ParseFloat(row[high], 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing high %s: %w", row[high], err)
	}
	ohlcv.High = h

	l, err := strconv.ParseFloat(row[low], 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing low %s: %w", row[low], err)
	}
	ohlcv.Low = l

	c, err := strconv.ParseFloat(row[open], 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing close %s: %w", row[close], err)
	}
	ohlcv.Close = c

	v, err := strconv.ParseInt(row[open], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("error parsing volume %s: %w", row[volume], err)
	}
	ohlcv.Volume = v
	entry.PriceData = *ohlcv

	return entry, nil
}

// Parse csv prices timeseries data from a reader
// Returns a map from timestamp to OHLCV
func ParsePricesTimeSeriesCSV(r io.Reader) (map[time.Time]*OHLCV, error) {
	reader := csv.NewReader(r)
	reader.ReuseRecord = true
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true
	reader.FieldsPerRecord = 6

	// drop header
	if _, err := reader.Read(); err != nil {
		if err == io.EOF {
			return nil, nil
		}
		return nil, err
	}

	entries := make(map[time.Time]*OHLCV)

	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		entry, err := parseCSVRecord(row)
		if err != nil {
			return nil, err
		}
		entries[entry.Timestamp] = &entry.PriceData
	}

	return entries, nil
}
