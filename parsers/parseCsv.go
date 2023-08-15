package parsers

import (
	"encoding/csv"
	"io"
)

// Each CSVParsableEntry corresponds to a row in a csv file
type CSVParsableEntry[Key comparable, Value any] interface {
	Key() Key
	Value() Value
}

// Parse csv data provided by a reader
func ParseCSV[Key comparable, Value any](r io.Reader, fieldsPerRecord int, rowParser func(row []string) (CSVParsableEntry[Key, Value], error)) (map[Key]Value, error) {
	reader := csv.NewReader(r)
	reader.ReuseRecord = true
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true
	reader.FieldsPerRecord = fieldsPerRecord

	// drop header
	if _, err := reader.Read(); err != nil {
		if err == io.EOF {
			return nil, nil
		}
		return nil, err
	}

	entries := make(map[Key]Value)

	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		entry, err := rowParser(row)
		if err != nil {
			return nil, err
		}
		entries[entry.Key()] = entry.Value()
	}

	return entries, nil
}
