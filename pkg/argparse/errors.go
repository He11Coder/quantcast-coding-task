package argparse

import "errors"

var (
	ErrMissedRequiredArg = errors.New("one of the required command line arguments is missing, run the command with --help or -h for more info on the arguments")
	ErrInvalidDateFormat = errors.New("date must be in YYYY-MM-DD format")
)
