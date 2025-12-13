package main

import (
	"fmt"
	"io"
	"time"

	"cookie-cli/pkg/logparse"
	"cookie-cli/pkg/stats"
	"cookie-cli/pkg/utils"
)

// RunCommand runs the tool's pipeline:
// creates a 1-day time window ([date; date+24hrs)) and invokes FindMostFrequent function
// to search for the most active cookie(s) within the time window among all of the cookies from reader.
// The result is written into writer.
func RunCommand(reader logparse.EntryReader, writer io.Writer, date time.Time) error {
	//create a 1-day time window
	dateToSearchFrom := date
	dateToSearchTo := date.Add(24 * time.Hour).Add(-1 * time.Nanosecond)

	results, err := stats.FindMostFrequent(reader, dateToSearchFrom, dateToSearchTo)
	if err != nil {
		return fmt.Errorf("error finding the most active cookie: %w", err)
	}

	err = utils.PrintResults(writer, results)
	if err != nil {
		return fmt.Errorf("error writing output: %w", err)
	}

	return nil
}
