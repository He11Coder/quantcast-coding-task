package logparse

import (
	"time"
)

type Entry struct {
	Cookie    string
	Timestamp time.Time
}
