package ht

import (
	"math"
)

// Horner's Method to hash string of length L (O(L))
func hashCode(s string) int {
	hash := int32(0)
	for i := 0; i < len(s); i++ {
		hash = int32(hash<<5-hash) + int32(s[i])
		hash &= hash
	}
	return int(math.Abs(float64(hash)))
}
