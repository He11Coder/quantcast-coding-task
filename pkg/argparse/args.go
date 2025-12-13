package argparse

import (
	"errors"
	"flag"
	"fmt"
	"time"
)

type Options struct {
	Filename string
	Date     time.Time
}

// Parse parses command line arguments passed in args slice,
// ensures that both -f and -d flags are present.
// Returns *Options containing the specified filename and date as a time.Time value.
// If -h or --help flag was used, returns nil, nil
func Parse(args []string) (*Options, error) {
	var filename, date string

	flagSet := flag.NewFlagSet("cookie-tool", flag.ContinueOnError)

	flagSet.StringVar(&filename, "f", "", "Cookie log file name to process")
	flagSet.StringVar(&date, "d", "", "Date in YYYY-MM-DD format to search")

	err := flagSet.Parse(args)
	if err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return nil, nil
		}
		return nil, fmt.Errorf("error parsing command line arguments: %w", err)
	}

	//omitted arguments are not allowed
	if (filename == "") || (date == "") {
		return nil, fmt.Errorf("-f is '%s', -d is '%s': %w", filename, date, ErrMissedRequiredArg)
	}

	//parsing string date to time.Time representation
	dateToSearch, err := time.Parse(time.DateOnly, date)
	if err != nil {
		return nil, fmt.Errorf("error parsing the specified date (you have entered: '%s'): %w", date, ErrInvalidDateFormat)
	}

	return &Options{
		Filename: filename,
		Date:     dateToSearch,
	}, nil
}
