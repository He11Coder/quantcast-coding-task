package stats

import (
	"strings"
	"testing"
	"time"

	"cookie-cli/pkg/logparse"

	"github.com/stretchr/testify/require"
)

type testOk struct {
	name             string
	inputLog         string
	dateToSearchFrom time.Time
	dateToSearchTo   time.Time
	expected         []string
}

type testFail struct {
	name              string
	inputLog          string
	dateToSearchFrom  time.Time
	dateToSearchTo    time.Time
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
		dateToSearchFrom: time.Date(2018, 10, 9, 0, 0, 0, 0, time.UTC),
		dateToSearchTo:   time.Date(2018, 10, 9, 23, 59, 59, 0, time.UTC),
		expected:         nil,
	},
	{
		name: "Single output out of one",
		inputLog: `cookie,timestamp
AtY0laUfhglK3lC7,2018-12-09T14:19:00+00:00
SAZuXPGUrfbcn5UA,2018-12-09T10:13:00+00:00
5UAVanZf6UtGyKVS,2018-12-09T07:25:00+00:00
AtY0laUfhglK3lC7,2018-12-09T06:19:00+00:00
SAZuXPGUrfbcn5UA,2018-12-08T22:03:00+00:00
4sMM2LxV07bPJzwf,2018-12-08T21:30:00+00:00
fbcn5UAVanZf6UtG,2018-12-08T09:30:00+00:00
4sMM2LxV07bPJzwf,2018-12-07T23:30:00+00:00`,
		dateToSearchFrom: time.Date(2018, 12, 7, 0, 0, 0, 0, time.UTC),
		dateToSearchTo:   time.Date(2018, 12, 7, 23, 59, 59, 0, time.UTC),
		expected:         []string{"4sMM2LxV07bPJzwf"},
	},
	{
		name: "Single output out of many",
		inputLog: `cookie,timestamp
AtY0laUfhglK3lC7,2018-12-09T14:19:00+00:00
SAZuXPGUrfbcn5UA,2018-12-09T10:13:00+00:00
5UAVanZf6UtGyKVS,2018-12-09T07:25:00+00:00
AtY0laUfhglK3lC7,2018-12-09T06:19:00+00:00
SAZuXPGUrfbcn5UA,2018-12-08T22:03:00+00:00
4sMM2LxV07bPJzwf,2018-12-08T21:30:00+00:00
fbcn5UAVanZf6UtG,2018-12-08T09:30:00+00:00
4sMM2LxV07bPJzwf,2018-12-07T23:30:00+00:00`,
		dateToSearchFrom: time.Date(2018, 12, 9, 0, 0, 0, 0, time.UTC),
		dateToSearchTo:   time.Date(2018, 12, 9, 23, 59, 59, 0, time.UTC),
		expected:         []string{"AtY0laUfhglK3lC7"},
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
		dateToSearchFrom: time.Date(2018, 12, 8, 0, 0, 0, 0, time.UTC),
		dateToSearchTo:   time.Date(2018, 12, 8, 23, 59, 59, 0, time.UTC),
		expected:         []string{"SAZuXPGUrfbcn5UA", "4sMM2LxV07bPJzwf", "fbcn5UAVanZf6UtG"},
	},
	{
		name: "2-day range (single output)",
		inputLog: `cookie,timestamp
AtY0laUfhglK3lC7,2018-12-09T14:19:00+00:00
SAZuXPGUrfbcn5UA,2018-12-09T10:13:00+00:00
5UAVanZf6UtGyKVS,2018-12-09T07:25:00+00:00
AtY0laUfhglK3lC7,2018-12-09T06:19:00+00:00
SAZuXPGUrfbcn5UA,2018-12-08T22:03:00+00:00
4sMM2LxV07bPJzwf,2018-12-08T21:30:00+00:00
fbcn5UAVanZf6UtG,2018-12-08T09:30:00+00:00
4sMM2LxV07bPJzwf,2018-12-07T23:30:00+00:00`,
		dateToSearchFrom: time.Date(2018, 12, 7, 0, 0, 0, 0, time.UTC),
		dateToSearchTo:   time.Date(2018, 12, 8, 23, 59, 59, 0, time.UTC),
		expected:         []string{"4sMM2LxV07bPJzwf"},
	},
	{
		name: "2-day range (multiple outputs)",
		inputLog: `cookie,timestamp
AtY0laUfhglK3lC7,2018-12-09T14:19:00+00:00
SAZuXPGUrfbcn5UA,2018-12-09T10:13:00+00:00
5UAVanZf6UtGyKVS,2018-12-09T07:25:00+00:00
AtY0laUfhglK3lC7,2018-12-09T06:19:00+00:00
SAZuXPGUrfbcn5UA,2018-12-08T22:03:00+00:00
4sMM2LxV07bPJzwf,2018-12-08T21:30:00+00:00
fbcn5UAVanZf6UtG,2018-12-08T09:30:00+00:00
4sMM2LxV07bPJzwf,2018-12-07T23:30:00+00:00`,
		dateToSearchFrom: time.Date(2018, 12, 8, 0, 0, 0, 0, time.UTC),
		dateToSearchTo:   time.Date(2018, 12, 9, 23, 59, 59, 0, time.UTC),
		expected:         []string{"SAZuXPGUrfbcn5UA", "AtY0laUfhglK3lC7"},
	},
	{
		name: "3-hour range (single output)",
		inputLog: `cookie,timestamp
AtY0laUfhglK3lC7,2018-12-10T14:29:00+00:00
SAZuXPGUrfbcn5UA,2018-12-10T12:13:00+00:00
AtY0laUfhglK3lC7,2018-12-10T11:30:00+00:00
SAZuXPGUrfbcn5UA,2018-12-10T11:19:00+00:00
4sMM2LxV07bPJzwf,2018-12-07T23:30:00+00:00`,
		dateToSearchFrom: time.Date(2018, 12, 10, 11, 30, 0, 0, time.UTC),
		dateToSearchTo:   time.Date(2018, 12, 10, 14, 29, 0, 0, time.UTC),
		expected:         []string{"AtY0laUfhglK3lC7"},
	},
	{
		name: "No matching date",
		inputLog: `cookie,timestamp
AtY0laUfhglK3lC7,2018-12-09T14:19:00+00:00
SAZuXPGUrfbcn5UA,2018-12-09T10:13:00+00:00
5UAVanZf6UtGyKVS,2018-12-09T07:25:00+00:00
AtY0laUfhglK3lC7,2018-12-09T06:19:00+00:00
SAZuXPGUrfbcn5UA,2018-12-08T22:03:00+00:00
4sMM2LxV07bPJzwf,2018-12-08T21:30:00+00:00
fbcn5UAVanZf6UtG,2018-12-08T09:30:00+00:00
4sMM2LxV07bPJzwf,2018-12-07T23:30:00+00:00`,
		dateToSearchFrom: time.Date(2018, 12, 20, 0, 0, 0, 0, time.UTC),
		dateToSearchTo:   time.Date(2018, 12, 20, 23, 59, 59, 0, time.UTC),
		expected:         nil,
	},
	{
		name:             "No input (headers only)",
		inputLog:         `cookie,timestamp`,
		dateToSearchFrom: time.Date(2018, 12, 20, 0, 0, 0, 0, time.UTC),
		dateToSearchTo:   time.Date(2018, 12, 20, 23, 59, 59, 0, time.UTC),
		expected:         nil,
	},
	{
		name:             "No input (empty file)",
		inputLog:         ``,
		dateToSearchFrom: time.Date(2018, 12, 20, 0, 0, 0, 0, time.UTC),
		dateToSearchTo:   time.Date(2018, 12, 20, 23, 59, 59, 0, time.UTC),
		expected:         nil,
	},
}

var testSuiteFail = []testFail{
	{
		name: "Error reading entry",
		inputLog: `cookie,timestamp
this_is_not_a_valid_csv_row`,
		dateToSearchFrom:  time.Date(2018, 12, 20, 0, 0, 0, 0, time.UTC),
		dateToSearchTo:    time.Date(2018, 12, 20, 23, 59, 59, 0, time.UTC),
		expectedErrorText: "error reading entry from reader",
	},
}

func TestFindMostFrequentOk(t *testing.T) {
	for _, tt := range testSuiteOk {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.inputLog)
			entryReader := logparse.NewCSVEntryReader(reader)

			result, err := FindMostFrequent(entryReader, tt.dateToSearchFrom, tt.dateToSearchTo)
			require.NoError(t, err)
			require.ElementsMatch(t, result, tt.expected)
		})
	}
}

func TestFindMostFrequentFail(t *testing.T) {
	for _, tt := range testSuiteFail {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.inputLog)
			entryReader := logparse.NewCSVEntryReader(reader)

			_, err := FindMostFrequent(entryReader, tt.dateToSearchFrom, tt.dateToSearchTo)
			require.ErrorContains(t, err, tt.expectedErrorText)
		})
	}
}
