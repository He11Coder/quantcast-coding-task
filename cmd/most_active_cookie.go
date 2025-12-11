package main

import (
	"fmt"
	"os"

	"cookie-cli/pkg/argparse"
	"cookie-cli/pkg/logparse"
)

func main() {
	options, err := argparse.Parse()
	if err != nil {
		fmt.Printf("Error parsing arguments: %v\n", err)
		return
	}

	file, err := os.Open(options.Filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	reader := logparse.NewCSVEntryReader(file)

	err = RunCommand(reader, os.Stdout, options.Date)
	if err != nil {
		fmt.Printf("Error during command execution: %v\n", err)
	}
}
