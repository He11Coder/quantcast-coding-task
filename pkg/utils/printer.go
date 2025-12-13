package utils

import (
	"fmt"
	"io"
)

// PrintResults prints cookies iteratively (one by one) to writer w appending \n after each item.
func PrintResults(w io.Writer, cookies []string) error {
	for _, cookie := range cookies {
		_, err := fmt.Fprintln(w, cookie)
		if err != nil {
			return fmt.Errorf("failed to write output: %w", err)
		}
	}

	return nil
}
