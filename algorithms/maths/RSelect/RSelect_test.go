package rselect

import (
	"math/rand"
	"sort"
	"testing"
)

func TestRSelect(t *testing.T) {
	expected := []int{1, 2, 3, 5, 6, 8}
	for i, want := range expected {
		arr := []int{5, 2, 6, 8, 3, 1}
		if got := RSelect(arr, len(arr), i); got != want {
			t.Errorf("RSelect(_, 6, %d) = %d, want %d", i, got, want)
		}
	}
}

func TestRSelectRandom(t *testing.T) {
	for trial := 0; trial < 500; trial++ {
		n := rand.Intn(50) + 1
		arr := make([]int, n)
		for j := range arr {
			arr[j] = rand.Intn(20)
		}
		sorted := append([]int(nil), arr...)
		sort.Ints(sorted)
		i := rand.Intn(n)
		if got := RSelect(arr, n, i); got != sorted[i] {
			t.Fatalf("RSelect(%v, %d, %d) = %d, want %d", arr, n, i, got, sorted[i])
		}
	}
}
