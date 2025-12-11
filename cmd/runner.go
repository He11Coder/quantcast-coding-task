package main

import (
	"fmt"
	"io"
	"time"

	"cookie-cli/pkg/logparse"
	"cookie-cli/pkg/stats"
	"cookie-cli/pkg/utils"
)

func RunCommand(reader logparse.EntryReader, writer io.Writer, date time.Time) error {
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
