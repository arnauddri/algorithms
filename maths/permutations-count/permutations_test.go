package permutations

import (
	"bufio"
	"os"
	"strconv"
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
	array := _getArray()

	for i := 0; i < b.N; i++ {
		recursiveCount(array)
	}
}

func BenchmarkIterative(b *testing.B) {
	array := _getArray()

	for i := 0; i < b.N; i++ {
		iterativeCount(array)
	}
}

func _getArray() []int {
	f, err := os.Open("./IntegerArray.txt")
	defer f.Close()
	check(err)

	numbers := make([]int, 0)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		s, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, s)
	}

	return numbers
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
