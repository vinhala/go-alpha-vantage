package corestock

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/ga42quy/go-alpha-vantage/parsers"
)

type OHLCVAdjusted struct {
	Open             float64 `json:"1. open"`
	High             float64 `json:"2. high"`
	Low              float64 `json:"3. low"`
	Close            float64 `json:"4. close"`
	AdjustedClose    float64 `json:"5. adjusted close"`
	Volume           float64 `json:"6. volume"`
	DividendAmount   float64 `json:"7. dividend amount"`
	SplitCoefficient float64 `json:"8. split coefficient"`
}

type TimeSeriesEntryAdjusted struct {
	Timestamp time.Time
	PriceData OHLCVAdjusted
}

// Requirement by csv parser
func (e *TimeSeriesEntryAdjusted) Key() time.Time {
	return e.Timestamp
}

func (e *TimeSeriesEntryAdjusted) Value() *OHLCVAdjusted {
	return &e.PriceData
}

type timeSeriesAdjustedCSVRecordParser struct {
	dateFormat    string
	skipSplitCoef bool
}

func (p *timeSeriesAdjustedCSVRecordParser) parseCSVRecord(row []string) (parsers.CSVParsableEntry[time.Time, *OHLCVAdjusted], error) {
	entry := &TimeSeriesEntryAdjusted{}
	ohlcvAdj := &OHLCVAdjusted{}

	t, err := time.Parse(p.dateFormat, row[0])
	if err != nil {
		return nil, fmt.Errorf("error parsing timestamp %s: %w", row[0], err)
	}
	entry.Timestamp = t

	o, err := strconv.ParseFloat(row[1], 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing open %s: %w", row[1], err)
	}
	ohlcvAdj.Open = o

	h, err := strconv.ParseFloat(row[2], 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing high %s: %w", row[2], err)
	}
	ohlcvAdj.High = h

	l, err := strconv.ParseFloat(row[3], 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing low %s: %w", row[3], err)
	}
	ohlcvAdj.Low = l

	c, err := strconv.ParseFloat(row[4], 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing close %s: %w", row[4], err)
	}
	ohlcvAdj.Close = c

	cAdj, err := strconv.ParseFloat(row[5], 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing adjusted close %s: %w", row[5], err)
	}
	ohlcvAdj.AdjustedClose = cAdj

	v, err := strconv.ParseFloat(row[6], 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing volume %s: %w", row[6], err)
	}
	ohlcvAdj.Volume = v

	da, err := strconv.ParseFloat(row[7], 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing dividend amount %s: %w", row[7], err)
	}
	ohlcvAdj.DividendAmount = da

	if !p.skipSplitCoef {
		sc, err := strconv.ParseFloat(row[8], 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing split coefficient %s: %w", row[8], err)
		}
		ohlcvAdj.SplitCoefficient = sc
	}
	entry.PriceData = *ohlcvAdj

	return entry, nil
}

// Parse adjusted prices timeseries data csv from a reader
// Returns a map from timestamp to OHLCVAdjusted
func parseTimeSeriesAdjustedCSV(r io.Reader, dateFormat string, skipSplitCoef bool) (map[time.Time]*OHLCVAdjusted, error) {
	parser := &timeSeriesAdjustedCSVRecordParser{dateFormat: dateFormat, skipSplitCoef: skipSplitCoef}
	columns := 9
	if skipSplitCoef {
		columns = 8
	}
	return parsers.ParseCSV[time.Time, *OHLCVAdjusted](r, columns, parser.parseCSVRecord)
}
