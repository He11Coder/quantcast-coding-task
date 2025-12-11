package utils

import (
	"fmt"
	"io"
)

func PrintResults(w io.Writer, cookies []string) error {
	for _, cookie := range cookies {
		_, err := fmt.Fprintln(w, cookie)
		if err != nil {
			return fmt.Errorf("failed to write output: %w", err)
		}
	}

	return nil
}
