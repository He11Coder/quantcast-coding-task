package argparse

import (
	"flag"
	"fmt"
	"time"
)

type Options struct {
	Filename string
	Date     time.Time
}

func Parse() (*Options, error) {
	var filename, date string

	flag.StringVar(&filename, "f", "", "Cookie log file name to process")
	flag.StringVar(&date, "d", "", "Date in YYYY-MM-DD format to search")
	flag.Parse()

	if (filename == "") || (date == "") {
		return nil, fmt.Errorf("-f is '%s', -d is '%s': %w", filename, date, ErrMissedRequiredArg)
	}

	dateToSearch, err := time.Parse(time.DateOnly, date)
	if err != nil {
		return nil, fmt.Errorf("error parsing the specified date (you have entered: '%s'): %w", date, ErrInvalidDateFormat)
	}

	return &Options{
		Filename: filename,
		Date:     dateToSearch,
	}, nil
}
