package nearearth_test

import (
	"testing"

	"github.com/smccutcheon-whs/solidinpractice/nearearth"
)

var fileCountTestCases = []struct {
	name     string
	filepath string
	expected int
	fail     bool
}{
	{
		name:     "no impacts",
		filepath: "testdata/nearearth_empty.json",
		expected: 0,
		fail:     false,
	},
	{
		name:     "one impact",
		filepath: "testdata/nearearth_one.json",
		expected: 1,
		fail:     false,
	},
	{
		name:     "two impacts",
		filepath: "testdata/nearearth_two.json",
		expected: 2,
		fail:     false,
	},
	{
		name:     "error",
		filepath: "testdata/nearearth_error.json",
		expected: 0,
		fail:     true,
	},
}

func TestCount(t *testing.T) {
	for _, tc := range fileCountTestCases {
		t.Run(tc.name, func(t *testing.T) {
			fs, err := nearearth.NewImpactFileService(nearearth.WithTestFilePath(tc.filepath))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			count, err := nearearth.Count(fs)
			if tc.fail {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if count != tc.expected {
				t.Fatalf("expected %d, got %d", tc.expected, count)
			}
		})
	}
}
