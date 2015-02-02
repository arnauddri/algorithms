package shell

import (
	"fmt"
	"github.com/arnauddri/algorithms/algorithms/sorting/utils"
	"testing"
)

func TestShellSort(t *testing.T) {
	list := utils.GetArrayOfSize(100)

	sort(list)

	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Error()
		}
	}
}
