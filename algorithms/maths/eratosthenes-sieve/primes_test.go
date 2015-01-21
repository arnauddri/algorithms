package primes

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetAllPrimesTo(t *testing.T) {
	primes := getAllPrimesTo(60)
	expected := []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59}
	if !reflect.DeepEqual(primes, expected) {
		fmt.Println(primes)
		t.Error()
	}
}
