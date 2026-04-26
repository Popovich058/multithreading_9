package main

import (
	"testing"


	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	testCases := []struct {
		name      string
		size      int
		expectedLen int
	}{
		{"positive_size", 5, 5},
		{"zero_size", 0, 0},
		{"negative_size", -10, 0},
		{"large_size", 10000, 10000},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := generateRandomElements(tc.size)
			require.Len(t, result, tc.expectedLen,
				"expected length %d, got %d", tc.expectedLen, len(result))
		})
	}

	t.Run("repeated_calls_likely_produce_different_results", func(t *testing.T) {
		size := 10
		result1 := generateRandomElements(size)
		result2 := generateRandomElements(size)

		different := false
		for i := range result1 {
			if result1[i] != result2[i] {
				different = true
				break
			}
		}

		if !different {
			t.Log("possible seed issue: identical results on repeated calls")
		}
	})
}

func TestMaximum(t *testing.T) {
	testCases := []struct {
		name     string
		data     []int
		expected int
	}{
		{"multiple_positive", []int{1, 5, 3, 9, 2}, 9},
		{"negative_numbers", []int{-10, -5, -20, -1}, -1},
		{"mixed_numbers", []int{-5, 0, 5, -10, 3}, 5},
		{"single_element", []int{42}, 42},
		{"empty_slice", []int{}, 0},
		{"nil_slice", nil, 0},
		{"max_at_beginning", []int{100, 1, 2, 3}, 100},
		{"max_at_end", []int{1, 2, 3, 100}, 100},
		{"all_same", []int{5, 5, 5, 5}, 5},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			max := maximum(tc.data)
			assert.Equal(t, tc.expected, max,
				"expected maximum %d, got %d for case %s", tc.expected, max, tc.name)
		})
	}
}

func TestMaxChunks(t *testing.T) {
	testCases := []struct {
		name       string
		data       []int
		expected   int
		expectErr  bool
	}{
		{"large_slice", []int{1, 5, 3, 9, 2, 8, 4, 7, 6}, 9, false},
		{"small_slice", []int{1, 3, 2}, 3, false},
		{"empty_slice", []int{}, 0, true},
		{"nil_slice", nil, 0, true},
		{"all_same_values", nil, 5, false}, 
		{"slice_equals_chunks", nil, 14, false}, 
	}

	for i, tc := range testCases {
		switch tc.name {
		case "all_same_values":
			testCases[i].data = make([]int, 10)
			for j := range testCases[i].data {
				testCases[i].data[j] = 5
			}
		case "slice_equals_chunks":
			testCases[i].data = make([]int, CHUNKS)
			for j := 0; j < CHUNKS; j++ {
				testCases[i].data[j] = j * 2
			}
		}
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			max, err := maxChunks(tc.data)
			if tc.expectErr {
				require.Error(t, err, "expected error for case %s", tc.name)
			} else {
				require.NoError(t, err, "unexpected error for case %s: %v", tc.name, err)
			}
			assert.Equal(t, tc.expected, max,
				"expected maximum %d, got %d for case %s", tc.expected, max, tc.name)
		})
	}
}
