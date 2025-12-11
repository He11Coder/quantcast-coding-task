package logparse

import (
	"encoding/csv"
	"fmt"
	"io"
	"time"
)

type EntryReader interface {
	ReadEntry() (Entry, error)
}

type csvEntryReader struct {
	csvReader *csv.Reader
}

func NewCSVEntryReader(r io.Reader) EntryReader {
	return &csvEntryReader{
		csvReader: csv.NewReader(r),
	}
}

func (r *csvEntryReader) ReadEntry() (Entry, error) {
	record, err := r.csvReader.Read()
	if err != nil {
		return Entry{}, err
	}

	if len(record) < 2 {
		return Entry{}, fmt.Errorf("got %d columns in a log entry: %w", len(record), ErrInvalidCSVFormat)
	}

	t, err := time.Parse(time.RFC3339, record[1])
	if err != nil {
		return Entry{}, fmt.Errorf("invalid timestamp format: %w", err)
	}

	return Entry{
		Cookie:    record[0],
		Timestamp: t,
	}, nil
}
