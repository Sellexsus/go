package main

import (
	"testing"
)

// generateRandomElements

func TestGenerateRandomElements_SizeZero(t *testing.T) {
	got := generateRandomElements(0)
	if len(got) != 0 {
		t.Fatalf("len(generateRandomElements(0)) = %d; want 0", len(got))
	}
}

func TestGenerateRandomElements_BasicProps(t *testing.T) {
	const n = 1_000
	got := generateRandomElements(n)

	if len(got) != n {
		t.Fatalf("len = %d; want %d", len(got), n)
	}

	for i, v := range got {
		if v < 0 || v >= 1_000_000_000 {
			t.Fatalf("value out of range at %d: %d", i, v)
		}
	}
}

//  maximum

func TestMaximum_EmptySlice(t *testing.T) {
	if max := maximum(nil); max != 0 {
		t.Fatalf("maximum(nil) = %d; want 0", max)
	}
	if max := maximum([]int{}); max != 0 {
		t.Fatalf("maximum([]) = %d; want 0", max)
	}
}

func TestMaximum_SingleElement(t *testing.T) {
	if max := maximum([]int{42}); max != 42 {
		t.Fatalf("maximum([42]) = %d; want 42", max)
	}
}

func TestMaximum_MixedValues(t *testing.T) {
	data := []int{-5, -2, -10, 0, 7, 3, -1}
	if max := maximum(data); max != 7 {
		t.Fatalf("maximum(%v) = %d; want 7", data, max)
	}
}

func TestMaximum_Ascending(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	if max := maximum(data); max != 5 {
		t.Fatalf("maximum(%v) = %d; want 5", data, max)
	}
}

// maxChunks vs maximum

func TestMaxChunks_EqualsMaximum(t *testing.T) {
	cases := [][]int{
		{},
		{9},
		{5, 4, 3},
		{1, 9, 3, 7, 2, 6},
	}

	{
		n := 17
		a := make([]int, n)
		for i := range a {
			a[i] = i
		}
		a[n-1] = 10_000_000
		cases = append(cases, a)
	}
	{
		n := 10_000
		a := make([]int, n)
		for i := range a {
			a[i] = i % 1234
		}
		a[n-1] = 99_999_999
		cases = append(cases, a)
	}

	for _, data := range cases {
		want := maximum(data)
		got := maxChunks(data)
		if got != want {
			t.Fatalf("maxChunks(%d elems) = %d; want %d; data(head)=%v",
				len(data), got, want, head(data, 10))
		}
	}
}

func head(a []int, k int) []int {
	if len(a) <= k {
		return a
	}
	cp := make([]int, k)
	copy(cp, a[:k])
	return cp
}
