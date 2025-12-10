package main

import (
	"fmt"
	"os"

	"cookie-cli/argparse"
	"cookie-cli/logparse"
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
	//TODO: implement the main processing logic
}
