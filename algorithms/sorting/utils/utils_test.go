package utils

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	for i := 0; i < 100000; i++ {
		if i%300 == 0 {
			array := GetArrayOfSize(i)

			if len(array) != i {
				fmt.Println(array)
				t.Error()
			}
		}
	}
}
