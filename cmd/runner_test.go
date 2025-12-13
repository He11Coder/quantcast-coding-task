package main

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"cookie-cli/pkg/logparse"

	"github.com/stretchr/testify/require"
)

type testOk struct {
	name     string
	inputLog string
	date     time.Time
	expected string
}

type testFail struct {
	name              string
	inputLog          string
	date              time.Time
	expectedErrorText string
}

var testSuiteOk = []testOk{
	{
		name: "No output",
		inputLog: `cookie,timestamp
AtY0laUfhglK3lC7,2018-12-09T14:19:00+00:00
SAZuXPGUrfbcn5UA,2018-12-09T10:13:00+00:00
5UAVanZf6UtGyKVS,2018-12-09T07:25:00+00:00
AtY0laUfhglK3lC7,2018-12-09T06:19:00+00:00
SAZuXPGUrfbcn5UA,2018-12-08T22:03:00+00:00
4sMM2LxV07bPJzwf,2018-12-08T21:30:00+00:00
fbcn5UAVanZf6UtG,2018-12-08T09:30:00+00:00
4sMM2LxV07bPJzwf,2018-12-07T23:30:00+00:00`,
		date:     time.Date(2018, 12, 1, 0, 0, 0, 0, time.UTC),
		expected: "",
	},
	{
		name: "Single output",
		inputLog: `cookie,timestamp
AtY0laUfhglK3lC7,2018-12-09T14:19:00+00:00
SAZuXPGUrfbcn5UA,2018-12-09T10:13:00+00:00
5UAVanZf6UtGyKVS,2018-12-09T07:25:00+00:00
AtY0laUfhglK3lC7,2018-12-09T06:19:00+00:00
SAZuXPGUrfbcn5UA,2018-12-08T22:03:00+00:00
4sMM2LxV07bPJzwf,2018-12-08T21:30:00+00:00
fbcn5UAVanZf6UtG,2018-12-08T09:30:00+00:00
4sMM2LxV07bPJzwf,2018-12-07T23:30:00+00:00`,
		date:     time.Date(2018, 12, 9, 0, 0, 0, 0, time.UTC),
		expected: "AtY0laUfhglK3lC7\n",
	},
	{
		name: "Multiple outputs",
		inputLog: `cookie,timestamp
AtY0laUfhglK3lC7,2018-12-09T14:19:00+00:00
SAZuXPGUrfbcn5UA,2018-12-09T10:13:00+00:00
5UAVanZf6UtGyKVS,2018-12-09T07:25:00+00:00
AtY0laUfhglK3lC7,2018-12-09T06:19:00+00:00
SAZuXPGUrfbcn5UA,2018-12-08T22:03:00+00:00
4sMM2LxV07bPJzwf,2018-12-08T21:30:00+00:00
fbcn5UAVanZf6UtG,2018-12-08T09:30:00+00:00
4sMM2LxV07bPJzwf,2018-12-07T23:30:00+00:00`,
		date:     time.Date(2018, 12, 8, 0, 0, 0, 0, time.UTC),
		expected: "SAZuXPGUrfbcn5UA\n4sMM2LxV07bPJzwf\nfbcn5UAVanZf6UtG\n",
	},
}

var testSuiteFail = []testFail{
	{
		name: "Error reading entry",
		inputLog: `cookie,timestamp
this_is_not_a_valid_csv_row`,
		date:              time.Date(2018, 12, 20, 0, 0, 0, 0, time.UTC),
		expectedErrorText: "error reading entry from reader",
	},
}

func TestRunCommandOk(t *testing.T) {
	for _, tt := range testSuiteOk {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.inputLog)
			entryReader := logparse.NewCSVEntryReader(reader)

			var outputWriter bytes.Buffer

			err := RunCommand(entryReader, &outputWriter, tt.date)
			require.NoError(t, err)

			outputString := outputWriter.String()
			if len(tt.expected) == 0 {
				require.Equal(t, outputString, tt.expected)
			} else {
				outputLines := strings.Split(outputString, "\n")
				expectedLines := strings.Split(tt.expected, "\n")

				require.ElementsMatch(t, outputLines, expectedLines)
			}
		})
	}
}

func TestRunCommandFail(t *testing.T) {
	for _, tt := range testSuiteFail {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.inputLog)
			entryReader := logparse.NewCSVEntryReader(reader)

			var outputWriter bytes.Buffer

			err := RunCommand(entryReader, &outputWriter, tt.date)
			require.ErrorContains(t, err, tt.expectedErrorText)
		})
	}
}
