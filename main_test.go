package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGenerateRandomElements проверяет корректность работы функции generateRandomElements.
func TestGenerateRandomElements(t *testing.T) {
	t.Run("positive_size_returns_correct_length", func(t *testing.T) {
		size := 5
		result := generateRandomElements(size)
		require.Len(t, result, size, "expected length %d, got %d", size, len(result))
	})

	t.Run("zero_size_returns_empty_slice", func(t *testing.T) {
		result := generateRandomElements(0)
		require.Empty(t, result, "expected empty slice when size=0, got length %d", len(result))
	})

	t.Run("negative_size_returns_empty_slice", func(t *testing.T) {
		result := generateRandomElements(-10)
		require.Empty(t, result, "expected empty slice for negative size, got length %d", len(result))
	})

	t.Run("large_size_does_not_panic", func(t *testing.T) {
		size := 10000
		result := generateRandomElements(size)
		require.Len(t, result, size, 
			"for large size %d, expected length %d, got %d",
			size, size, len(result))
	})

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
	t.Run("multiple_positive_numbers", func(t *testing.T) {
		data := []int{1, 5, 3, 9, 2}
		max, err := maximum(data)
		require.NoError(t, err, "unexpected error: %v", err)
		assert.Equal(t, 9, max, "expected maximum value 9, got %d", max)
	})

	t.Run("negative_numbers", func(t *testing.T) {
		data := []int{-10, -5, -20, -1}
		max, err := maximum(data)
		require.NoError(t, err, "unexpected error: %v", err)
		assert.Equal(t, -1, max, "expected maximum value -1, got %d", max)
	})

	t.Run("mixed_numbers", func(t *testing.T) {
		data := []int{-5, 0, 5, -10, 3}
		max, err := maximum(data)
		require.NoError(t, err, "unexpected error: %v", err)
		assert.Equal(t, 5, max, "expected maximum value 5, got %d", max)
	})

	t.Run("single_element", func(t *testing.T) {
		data := []int{42}
		max, err := maximum(data)
		require.NoError(t, err, "unexpected error: %v", err)
		assert.Equal(t, 42, max, "expected single element 42, got %d", max)
	})

	t.Run("empty_slice_returns_error", func(t *testing.T) {
		data := []int{}
		max, err := maximum(data)
		require.Error(t, err, "expected error for empty slice, but no error returned")
		assert.Zero(t, max, "expected zero value for max on empty slice, got %d", max)
	})

	t.Run("nil_slice_returns_error", func(t *testing.T) {
		var data []int = nil
		max, err := maximum(data)
		require.Error(t, err, "expected error for nil slice, but no error returned")
		assert.Zero(t, max, "expected zero value for max on nil slice, got %d", max)
	})

	t.Run("max_at_beginning", func(t *testing.T) {
		data := []int{100, 1, 2, 3}
		max, err := maximum(data)
		require.NoError(t, err, "unexpected error: %v", err)
		assert.Equal(t, 100, max, "expected max 100 at beginning, got %d", max)
	})

	t.Run("max_at_end", func(t *testing.T) {
		data := []int{1, 2, 3, 100}
		max, err := maximum(data)
		require.NoError(t, err, "unexpected error: %v", err)
		assert.Equal(t, 100, max, "expected max 100 at end, got %d", max)
	})

	t.Run("all_elements_same", func(t *testing.T) {
		data := []int{5, 5, 5, 5}
		max, err := maximum(data)
		require.NoError(t, err, "unexpected error: %v", err)
		assert.Equal(t, 5, max, "expected max 5 for identical elements, got %d", max)
	})
}


func TestMaxChunks(t *testing.T) {
	t.Run("large_slice_splits_into_chunks", func(t *testing.T) {
		data := []int{1, 5, 3, 9, 2, 8, 4, 7, 6} // len=9 > CHUNKS=8
		max, err := maxChunks(data)
		require.NoError(t, err, "unexpected error: %v", err)
		assert.Equal(t, 9, max, "expected maximum value 9, got %d", max)
	})

	t.Run("small_slice_uses_single_chunk", func(t *testing.T) {
		data := []int{1, 3, 2} // len=3 < CHUNKS=8
		max, err := maxChunks(data)
		require.NoError(t, err, "unexpected error: %v", err)
		assert.Equal(t, 3, max, "expected maximum value 3, got %d", max)
	})

	t.Run("empty_slice_returns_error", func(t *testing.T) {
		data := []int{}
		max, err := maxChunks(data)
		require.Error(t, err, "expected error for empty slice, but no error returned")
		assert.Zero(t, max, "expected zero value for max on empty slice, got %d", max)
	})

	t.Run("nil_slice_returns_error", func(t *testing.T) {
		var data []int = nil
		max, err := maxChunks(data)
		require.Error(t, err, "expected error for nil slice, but no error returned")
		assert.Zero(t, max, "expected zero value for max on nil slice, got %d", max)
	})

	t.Run("all_elements_same_value", func(t *testing.T) {
		data := make([]int, 10)
		for i := range data {
			data[i] = 5
		}
		max, err := maxChunks(data)
		require.NoError(t, err, "unexpected error: %v", err)
		assert.Equal(t, 5, max, "expected max 5 for identical elements, got %d", max)
	})

	t.Run("slice_length_equals_chunks", func(t *testing.T) {
		data := make([]int, CHUNKS)
		for i := 0; i < CHUNKS; i++ {
			data[i] = i * 2 // 0, 2, 4, ..., 14 (при CHUNKS=8)
		}
		max, err := maxChunks(data)
		require.NoError(t, err, "unexpected error: %v", err)
		assert.Equal(t, 14, max, "expected maximum value 14, got %d", max)
	})
}
