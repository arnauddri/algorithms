package heap

import (
	"fmt"
	"github.com/arnauddri/algorithms/algorithms/sorting/utils"
	"testing"
)

func TestHeapSort(t *testing.T) {
	list := utils.GetArrayOfSize(10)

	sort(list)

	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Error()
		}
	}
}

func benchmarkHeapSort(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)
	for i := 0; i < b.N; i++ {
		sort(list)
	}
}

func BenchmarkHeapSort100(b *testing.B)    { benchmarkHeapSort(100, b) }
func BenchmarkHeapSort1000(b *testing.B)   { benchmarkHeapSort(1000, b) }
func BenchmarkHeapSort10000(b *testing.B)  { benchmarkHeapSort(10000, b) }
func BenchmarkHeapSort100000(b *testing.B) { benchmarkHeapSort(100000, b) }
