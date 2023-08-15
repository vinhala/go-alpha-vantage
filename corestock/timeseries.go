package corestock

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/ga42quy/go-alpha-vantage/parsers"
)

type OutputSize string

const (
	COMPACT OutputSize = "compact"
	FULL    OutputSize = "full"
)

// A single entry in a prices time series
type TimeSeriesEntry struct {
	Timestamp time.Time
	PriceData OHLCV
}

// Requirement by csv parser
func (e *TimeSeriesEntry) Key() time.Time {
	return e.Timestamp
}

func (e *TimeSeriesEntry) Value() *OHLCV {
	return &e.PriceData
}

type timeSeriesCSVRecordParser struct {
	dateFormat string
}

func (p *timeSeriesCSVRecordParser) parseTimeSeriesCSVRecord(row []string) (parsers.CSVParsableEntry[time.Time, *OHLCV], error) {
	entry := &TimeSeriesEntry{}
	ohlcv := &OHLCV{}

	t, err := time.Parse(p.dateFormat, row[0])
	if err != nil {
		return nil, fmt.Errorf("error parsing timestamp %s: %w", row[0], err)
	}
	entry.Timestamp = t

	o, err := strconv.ParseFloat(row[1], 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing open %s: %w", row[1], err)
	}
	ohlcv.Open = o

	h, err := strconv.ParseFloat(row[2], 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing high %s: %w", row[2], err)
	}
	ohlcv.High = h

	l, err := strconv.ParseFloat(row[3], 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing low %s: %w", row[3], err)
	}
	ohlcv.Low = l

	c, err := strconv.ParseFloat(row[4], 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing close %s: %w", row[4], err)
	}
	ohlcv.Close = c

	v, err := strconv.ParseFloat(row[5], 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing volume %s: %w", row[5], err)
	}
	ohlcv.Volume = v
	entry.PriceData = *ohlcv

	return entry, nil
}

// Parse csv prices timeseries data from a reader
// Returns a map from timestamp to OHLCV
func parseTimeSeriesCSV(r io.Reader, dateFormat string) (map[time.Time]*OHLCV, error) {
	parser := &timeSeriesCSVRecordParser{dateFormat: dateFormat}
	return parsers.ParseCSV[time.Time, *OHLCV](r, 6, parser.parseTimeSeriesCSVRecord)
}
