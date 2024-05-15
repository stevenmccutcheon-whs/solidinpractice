package nearearth_test

import (
	"fmt"
	"testing"

	"github.com/smccutcheon-whs/solidinpractice/nearearth"
)

var (
	countTestCases = []struct {
		name     string
		impacts  nearearth.Impacts
		expected int
		fail     bool
		err      error
	}{
		{
			name:     "no impacts",
			impacts:  nearearth.Impacts{},
			expected: 0,
			fail:     false,
			err:      nil,
		},
		{
			name: "one impact",
			impacts: nearearth.Impacts{
				{
					Designation:   "2021 AB",
					DiscoveryDate: "2021-01-01",
				},
			},
			expected: 1,
			fail:     false,
			err:      nil,
		},
		{
			name: "two impacts",
			impacts: nearearth.Impacts{
				{
					Designation:   "2021 AB",
					DiscoveryDate: "2021-01-01",
				},
				{
					Designation:   "2021 AC",
					DiscoveryDate: "2021-01-02",
				},
			},
			expected: 2,
			fail:     false,
			err:      nil,
		},
		{
			name:     "error",
			impacts:  nearearth.Impacts{},
			expected: 0,
			fail:     true,
			err:      fmt.Errorf("error"),
		},
	}
)

func Test_Count(t *testing.T) {
	for _, tc := range countTestCases {
		fake := nearearth.NewFakeImpactService(t, nearearth.OnGetAll(func() (nearearth.Impacts, error) {
			return tc.impacts, tc.err
		}))
		t.Run(tc.name, func(t *testing.T) {
			count, err := nearearth.Count(fake)
			if tc.fail {
				if err == nil {
					t.Fatalf("expected error but got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if count != tc.expected {
				t.Fatalf("expected count %d but got %d", tc.expected, count)
			}
		})
	}
}

var getAllTestCases = []struct {
	name    string
	impacts nearearth.Impacts
	fail    bool
	err     error
}{
	{
		name:    "no impacts",
		impacts: nearearth.Impacts{},
		fail:    false,
		err:     nil,
	},
	{
		name: "one impact",
		impacts: nearearth.Impacts{
			{
				Designation:   "2021 AB",
				DiscoveryDate: "2021-01-01",
			},
		},
		fail: false,
		err:  nil,
	},
	{
		name: "two impacts",
		impacts: nearearth.Impacts{
			{
				Designation:   "2021 AB",
				DiscoveryDate: "2021-01-01",
			},
			{
				Designation:   "2021 AC",
				DiscoveryDate: "2021-01-02",
			},
		},
		fail: false,
		err:  nil,
	},
	{
		name:    "error",
		impacts: nearearth.Impacts{},
		fail:    true,
		err:     fmt.Errorf("error"),
	},
}

func Test_GetAll(t *testing.T) {
	for _, tc := range getAllTestCases {
		fake := nearearth.NewFakeImpactService(t, nearearth.OnGetAll(func() (nearearth.Impacts, error) {
			return tc.impacts, tc.err
		}))
		t.Run(tc.name, func(t *testing.T) {
			impacts, err := nearearth.GetAllImpacts(fake)
			if tc.fail {
				if err == nil {
					t.Fatalf("expected error but got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(impacts) != len(tc.impacts) {
				t.Fatalf("expected %d impacts but got %d", len(tc.impacts), len(impacts))
			}
			if impacts.String() != tc.impacts.String() {
				t.Fatalf("expected %s but got %s", tc.impacts.String(), impacts.String())
			}
			for i := range impacts {
				if impacts[i].String() != tc.impacts[i].String() {
					t.Fatalf("expected %s but got %s", tc.impacts[i].Designation, impacts[i].Designation)
				}
			}
		})
	}
}
