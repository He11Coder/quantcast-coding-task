package stats

import (
	"io"
	"time"

	"cookie-cli/logparse"
)

func FindMostFrequent(reader logparse.EntryReader, dateToSearchFrom, dateToSearchTo time.Time) ([]string, error) {
	cookieCounts := make(map[string]int)

	for {
		entry, err := reader.ReadEntry()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if entry.Timestamp.After(dateToSearchTo) {
			continue
		}
		if entry.Timestamp.Before(dateToSearchFrom) {
			break
		}

		cookieCounts[entry.Cookie]++
	}

	maxCount := 0
	for _, count := range cookieCounts {
		if count > maxCount {
			maxCount = count
		}
	}

	var result []string
	for cookie, count := range cookieCounts {
		if count == maxCount {
			result = append(result, cookie)
		}
	}

	return result, nil
}
