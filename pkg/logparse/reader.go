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

// Implementation of EntryReader interface that supports CSV reading.
type csvEntryReader struct {
	csvReader *csv.Reader
}

// csvEntryReader's constructor.
// Reads and discards the first line from the reader r (i.e., headers). Returns an EntryReader that is ready to do ReadEntry().
func NewCSVEntryReader(r io.Reader) EntryReader {
	csvR := csv.NewReader(r)
	_, _ = csvR.Read()

	return &csvEntryReader{
		csvReader: csvR,
	}
}

// ReadEntry reads another line from csvReader validating the number of fields (must be at least 2)
// and converting each line (i.e., CSV entry) into a structure of Entry type.
// Returns an instance of Entry.
func (r *csvEntryReader) ReadEntry() (Entry, error) {
	record, err := r.csvReader.Read()
	if err != nil {
		return Entry{}, err
	}

	//each line must contain at least 2 fields (cookie and timestamp)
	if len(record) < 2 {
		return Entry{}, fmt.Errorf("got %d columns in a log entry: %w", len(record), ErrInvalidCSVFormat)
	}

	//parsing time (time.RFC3339 format) from string to time.Time representation
	t, err := time.Parse(time.RFC3339, record[1])
	if err != nil {
		return Entry{}, fmt.Errorf("invalid timestamp format: %w", err)
	}

	return Entry{
		Cookie:    record[0],
		Timestamp: t.UTC(),
	}, nil
}
