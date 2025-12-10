package logparse

import "errors"

var (
	ErrInvalidFormat = errors.New("got wrong input format; the expected format implies 2-value entries: 'cookie,timestamp'")
)
