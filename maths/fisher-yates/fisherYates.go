package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(shuffle(arr))
}

func shuffle(a []int) []int {
	rand.Seed(time.Now().UTC().UnixNano())

	for i := len(a) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}

	return a
}
