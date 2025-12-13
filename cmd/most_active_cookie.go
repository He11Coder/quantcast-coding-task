package main

import (
	"fmt"
	"os"

	"cookie-cli/pkg/argparse"
	"cookie-cli/pkg/logparse"
)

func main() {
	//parsing command line arguments (filename and date)
	options, err := argparse.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("Error parsing arguments: %v\n", err)
		return
	}
	if options == nil {
		return
	}

	//open log file
	file, err := os.Open(options.Filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	//create EntryReader to read from CSV file
	reader := logparse.NewCSVEntryReader(file)

	//run the core business logic
	err = RunCommand(reader, os.Stdout, options.Date)
	if err != nil {
		fmt.Printf("Error during command execution: %v\n", err)
	}
}
