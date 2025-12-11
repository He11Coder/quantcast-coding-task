package logparse

import "errors"

var (
	ErrInvalidCSVFormat = errors.New("got wrong input format; the expected format implies entries starting with: 'cookie,timestamp' columns")
)
