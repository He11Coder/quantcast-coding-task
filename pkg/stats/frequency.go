package stats

import (
	"fmt"
	"io"
	"time"

	"cookie-cli/pkg/logparse"
)

// FindMostFrequent searches for a cookie or a set of cookies that occurred the most number of times in reader (the most frequent cookies).
// It only searches for such cookies within a time range of [dateToSearchFrom; dateToSearchTo].
// Returns a slice containing the most frequent cookie(s). If no cookies were found, returns nil, nil.
func FindMostFrequent(reader logparse.EntryReader, dateToSearchFrom, dateToSearchTo time.Time) ([]string, error) {
	//map to count cookies
	cookieCounts := make(map[string]int)

	for {
		entry, err := reader.ReadEntry()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading entry from reader: %w", err)
		}

		//since the data is sorted by time, skip lines until we reach the required time range
		if entry.Timestamp.After(dateToSearchTo) {
			continue
		}
		//break the loop when we get out of the required time range
		if entry.Timestamp.Before(dateToSearchFrom) {
			break
		}

		//incrementing counter for a cookie within the required time range
		cookieCounts[entry.Cookie]++
	}

	//find how many times the most frequent cookie(s) occurred
	maxCount := 0
	for _, count := range cookieCounts {
		if count > maxCount {
			maxCount = count
		}
	}

	//find the most frequent cookie(s)
	var result []string
	for cookie, count := range cookieCounts {
		if count == maxCount {
			result = append(result, cookie)
		}
	}

	return result, nil
}
