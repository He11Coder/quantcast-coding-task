package logparse

import (
	"io"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type testOk struct {
	name     string
	inputLog string
	expected []Entry
}

type testFail struct {
	name     string
	inputLog string
	expected error
}

var testSuiteOk = []testOk{
	{
		name: "Single entry",
		inputLog: `cookie,timestamp
AtY0laUfhglK3lC7,2018-12-09T14:19:00+00:00`,
		expected: []Entry{
			{
				Cookie:    "AtY0laUfhglK3lC7",
				Timestamp: time.Date(2018, 12, 9, 14, 19, 0, 0, time.UTC),
			},
		},
	},
	{
		name: "Multiple entries",
		inputLog: `cookie,timestamp
4sMM2LxV07bPJzwf,2018-12-08T21:30:00+00:00
5UAVanZf6UtGyKVS,2018-12-09T07:25:00+00:00
SAZuXPGUrfbcn5UA,2018-12-08T22:03:00+00:00
AtY0laUfhglK3lC7,2018-12-09T06:19:00+00:00
SAZuXPGUrfbcn5UA,2018-12-09T10:13:00+00:00
AtY0laUfhglK3lC7,2018-12-09T14:19:00+00:00
4sMM2LxV07bPJzwf,2018-12-07T23:30:00+00:00
fbcn5UAVanZf6UtG,2018-12-08T09:30:00+00:00`,
		expected: []Entry{
			{
				Cookie:    "4sMM2LxV07bPJzwf",
				Timestamp: time.Date(2018, 12, 8, 21, 30, 0, 0, time.UTC),
			},
			{
				Cookie:    "5UAVanZf6UtGyKVS",
				Timestamp: time.Date(2018, 12, 9, 7, 25, 0, 0, time.UTC),
			},
			{
				Cookie:    "SAZuXPGUrfbcn5UA",
				Timestamp: time.Date(2018, 12, 8, 22, 03, 0, 0, time.UTC),
			},
			{
				Cookie:    "AtY0laUfhglK3lC7",
				Timestamp: time.Date(2018, 12, 9, 06, 19, 0, 0, time.UTC),
			},
			{
				Cookie:    "SAZuXPGUrfbcn5UA",
				Timestamp: time.Date(2018, 12, 9, 10, 13, 0, 0, time.UTC),
			},
			{
				Cookie:    "AtY0laUfhglK3lC7",
				Timestamp: time.Date(2018, 12, 9, 14, 19, 0, 0, time.UTC),
			},
			{
				Cookie:    "4sMM2LxV07bPJzwf",
				Timestamp: time.Date(2018, 12, 7, 23, 30, 0, 0, time.UTC),
			},
			{
				Cookie:    "fbcn5UAVanZf6UtG",
				Timestamp: time.Date(2018, 12, 8, 9, 30, 0, 0, time.UTC),
			},
		},
	},
	{
		name: "Extra fields",
		inputLog: `cookie,timestamp,some_extra_field_header
AtY0laUfhglK3lC7,2018-12-09T14:19:00+00:00,some_extra_field_value
fbcn5UAVanZf6UtG,2018-12-08T09:30:00+00:00,some_extra_field_value`,
		expected: []Entry{
			{
				Cookie:    "AtY0laUfhglK3lC7",
				Timestamp: time.Date(2018, 12, 9, 14, 19, 0, 0, time.UTC),
			},
			{
				Cookie:    "fbcn5UAVanZf6UtG",
				Timestamp: time.Date(2018, 12, 8, 9, 30, 0, 0, time.UTC),
			},
		},
	},
}

var testSuiteFail = []testFail{
	{
		name:     "No input (headers only)",
		inputLog: `cookie,timestamp`,
		expected: io.EOF,
	},
	{
		name:     "No input (empty file)",
		inputLog: ``,
		expected: io.EOF,
	},
	{
		name: "Insufficient number of fields",
		inputLog: `cookie
SAZuXPGUrfbcn5UA`,
		expected: ErrInvalidCSVFormat,
	},
}

func TestReadEntryOk(t *testing.T) {
	for _, tt := range testSuiteOk {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.inputLog)
			entryReader := NewCSVEntryReader(reader)

			i := 0
			for ; ; i++ {
				entry, err := entryReader.ReadEntry()
				if err == io.EOF {
					break
				}

				require.NoError(t, err)
				require.Equal(t, entry, tt.expected[i])
			}

			require.Equal(t, len(tt.expected), i)
		})
	}
}

func TestReadEntryFail(t *testing.T) {
	for _, tt := range testSuiteFail {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.inputLog)
			entryReader := NewCSVEntryReader(reader)

			_, err := entryReader.ReadEntry()
			require.ErrorIs(t, err, tt.expected)
		})
	}
}
