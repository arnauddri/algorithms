package bubble

import (
	"fmt"
	"github.com/arnauddri/algorithms/algorithms/sorting/utils"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	list := dummy.GetArrayOfSize(100)

	sort(list)

	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Error()
		}
	}
}
