package argparse

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type testOk struct {
	name      string
	inputArgs []string
	expected  *Options
}

type testFail struct {
	name      string
	inputArgs []string
	expected  error
}

var testSuiteOk = []testOk{
	{
		name:      "Valid arguments",
		inputArgs: []string{"-f", "cookie_log.txt", "-d", "2018-10-09"},
		expected: &Options{
			Filename: "cookie_log.txt",
			Date:     time.Date(2018, 10, 9, 0, 0, 0, 0, time.UTC),
		},
	},
	{
		name:      "Help flag (-h)",
		inputArgs: []string{"-h"},
		expected:  nil,
	},
	{
		name:      "Help flag (--help)",
		inputArgs: []string{"--help"},
		expected:  nil,
	},
}

var testSuiteFail = []testFail{
	{
		name:      "No filename argument",
		inputArgs: []string{"-d", "2018-10-09"},
		expected:  ErrMissedRequiredArg,
	},
	{
		name:      "No date argument",
		inputArgs: []string{"-f", "cookie_log.txt"},
		expected:  ErrMissedRequiredArg,
	},
	{
		name:      "No arguments",
		inputArgs: []string{},
		expected:  ErrMissedRequiredArg,
	},
	{
		name:      "Empty filename argument",
		inputArgs: []string{"-f", "", "-d", "2018-10-09"},
		expected:  ErrMissedRequiredArg,
	},
	{
		name:      "Invalid date format (YYYY/MM/DD)",
		inputArgs: []string{"-f", "cookie_log.txt", "-d", "2018/10/09"},
		expected:  ErrInvalidDateFormat,
	},
	{
		name:      "Invalid date format (DD-MM-YYYY)",
		inputArgs: []string{"-f", "cookie_log.txt", "-d", "09-10-2018"},
		expected:  ErrInvalidDateFormat,
	},
	{
		name:      "Invalid date format (YYYY-MM-DD with time)",
		inputArgs: []string{"-f", "cookie_log.txt", "-d", "2018-10-09T23:30:00+00:00"},
		expected:  ErrInvalidDateFormat,
	},
}

func TestParseOk(t *testing.T) {
	for _, tt := range testSuiteOk {
		t.Run(tt.name, func(t *testing.T) {
			options, err := Parse(tt.inputArgs)
			require.NoError(t, err)
			require.Equal(t, options, tt.expected)
		})
	}
}

func TestParseFail(t *testing.T) {
	for _, tt := range testSuiteFail {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Parse(tt.inputArgs)
			require.ErrorIs(t, err, tt.expected)
		})
	}
}
