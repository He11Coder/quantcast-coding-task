package argparse

import (
	"flag"
	"fmt"
)

type Options struct {
	Filename string
	Date     string
}

func Parse() (*Options, error) {
	options := &Options{}

	flag.StringVar(&options.Filename, "f", "", "Cookie log file name to process")
	flag.StringVar(&options.Date, "d", "", "Date in YYYY-MM-DD format to search")
	flag.Parse()

	if (options.Filename == "") || (options.Date == "") {
		return nil, fmt.Errorf("-f is '%s', -d is '%s': %w", options.Filename, options.Date, ErrMissedRequiredArg)
	}

	return options, nil
}
