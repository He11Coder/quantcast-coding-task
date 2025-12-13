package logparse

import (
	"time"
)

// Structure representing a single entry from a cookie log.
type Entry struct {
	Cookie    string
	Timestamp time.Time
}
