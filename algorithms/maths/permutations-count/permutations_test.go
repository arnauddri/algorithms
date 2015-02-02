package permutations

import (
	"github.com/arnauddri/algorithms/algorithms/sorting/utils"
	"testing"
)

func TestIterative(t *testing.T) {
	a := []int{1, 3, 5, 2, 4, 6}
	count := iterativeCount(a)

	if count != 3 {
		t.Error("iterative count failed")
	}
}

func TestRecursive(t *testing.T) {
	a := []int{1, 3, 5, 2, 4, 6}
	_, count := recursiveCount(a)

	if count != 3 {
		t.Error("iterative count failed")
	}
}

func BenchmarkRecursive(b *testing.B) {
	array := utils.GetArrayOfSize(10000)

	for i := 0; i < b.N; i++ {
		recursiveCount(array)
	}
}

func BenchmarkIterative(b *testing.B) {
	array := utils.GetArrayOfSize(10000)

	for i := 0; i < b.N; i++ {
		iterativeCount(array)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
