package snippets

import (
	"fmt"
	"testing"
)

func Test_mergeIntervals(t *testing.T) {
	type test struct {
		input     IntervalArray
		cutoffMax float64
		expect    IntervalArray
	}
	tests := []test{
		{
			[]Interval{},
			0,
			[]Interval{},
		},
		{
			[]Interval{{1, 2}},
			0,
			[]Interval{},
		},
		{
			[]Interval{{1, 2}},
			1.5,
			[]Interval{{1, 1.5}},
		},
		{
			[]Interval{{1, 2}},
			2,
			[]Interval{{1, 2}},
		},
		{
			[]Interval{{1, 2}, {2, 3}},
			4,
			[]Interval{{1, 3}},
		},
		{
			[]Interval{{1, 2}, {2, 4}},
			2,
			[]Interval{{1, 2}},
		},
		{
			[]Interval{{1, 2}, {5, 6}},
			2,
			[]Interval{{1, 2}},
		},
		{
			[]Interval{{1, 2}, {5, 6}},
			6,
			[]Interval{{1, 2}, {5, 6}},
		},
		{
			[]Interval{{3, 4}, {1, 2}},
			4,
			[]Interval{{1, 2}, {3, 4}},
		},
		{
			[]Interval{{2, 3}, {1, 2}},
			4,
			[]Interval{{1, 3}},
		},
		{
			[]Interval{{1, 2}, {0, 1}, {-1, 2}},
			2,
			[]Interval{{-1, 2}},
		},
		{
			[]Interval{{1, 2}, {0, 3}},
			3,
			[]Interval{{0, 3}},
		},
		{
			[]Interval{{1, 2}, {1.5, 2.5}},
			1.5,
			[]Interval{{1, 1.5}},
		},
		{
			[]Interval{{1, 2}, {1.5, 2.5}},
			1,
			[]Interval{{1, 1}},
		},
		{
			[]Interval{{1, 2}, {1.5, 2.5}},
			10,
			[]Interval{{1, 2.5}},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			itsCopy := make(IntervalArray, len(tt.input))
			copy(itsCopy, tt.input)
			mergeIntervals(&itsCopy, tt.cutoffMax)
			if !compareIntervalArray(itsCopy, tt.expect) {
				t.Errorf("mergeIntervals(%v, maxMin=%v):\n  Expected %v, got %v", tt.input, tt.cutoffMax, tt.expect, itsCopy)
			}
		})
	}
}

func compareIntervalArray(a, b IntervalArray) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if abs(a[i].min-b[i].min) > eplison {
			return false
		}
		if abs(a[i].max-b[i].max) > eplison {
			return false
		}
	}
	return true
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
