package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{"Zero size", 0},
		{"One", 1},
		{"Ten", 10},
		{"Large", 1_000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)

			if tt.size == 0 {
				assert.Nil(t, data, "nil slice for size=0")
				return
			}

			require.NotNil(t, data)
			assert.Len(t, data, tt.size)

			for i, v := range data {
				assert.GreaterOrEqualf(t, v, 0, "value at %d must be >= 0", i)
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want int
	}{
		{"Nil slice", nil, 0},
		{"Empty slice", []int{}, 0},
		{"Single", []int{42}, 42},
		{"Ascending", []int{1, 2, 3, 4, 5}, 5},
		{"Mixed", []int{-5, -2, -10, 0, 7, 3, -1}, 7},
		{"All equal", []int{7, 7, 7}, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.in)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMaxChunks_EqualsMaximum(t *testing.T) {
	tests := []struct {
		name string
		in   []int
	}{
		{"Empty", []int{}},
		{"One", []int{9}},
		{"LessThanChunks", []int{5, 4, 3}},
		{"Arbitrary", []int{1, 9, 3, 7, 2, 6}},
		{"SkewSmall", func() []int {
			n := 17
			a := make([]int, n)
			for i := range a {
				a[i] = i
			}
			a[n-1] = 10_000_000
			return a
		}()},
		{"SkewBigger", func() []int {
			n := 10_000
			a := make([]int, n)
			for i := range a {
				a[i] = i % 1234
			}
			a[n-1] = 99_999_999
			return a
		}()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := maximum(tt.in)
			got := maxChunks(tt.in)
			require.Equal(t, want, got, "maxChunks = maximum")
		})
	}
}
