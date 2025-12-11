package main

import (
	"fmt"
	"os"
	"time"

	"cookie-cli/argparse"
	"cookie-cli/logparse"
	"cookie-cli/stats"
	"cookie-cli/utils"
)

func main() {
	options, err := argparse.Parse()
	if err != nil {
		fmt.Printf("Error parsing arguments: %v", err)
		return
	}

	file, err := os.Open(options.Filename)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return
	}
	defer file.Close()

	reader := logparse.NewCSVEntryReader(file)

	dateToSearchFrom := options.Date
	dateToSearchTo := options.Date.Add(24 * time.Hour).Add(-1 * time.Nanosecond)

	results, err := stats.FindMostFrequent(reader, dateToSearchFrom, dateToSearchTo)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return
	}

	err = utils.PrintResults(os.Stdout, results)
	if err != nil {
		fmt.Printf("Error writing output: %v", err)
		return
	}
}
