package ga

import (
	"errors"
	"math/rand"
)

// GenerateBitString returns an encoded string as set by calls SetGenerateBitString. Defaults to binary strings
var GenerateCandidate func(int) []int = func(length int) []int {
	if length <= 0 {
		panic(errors.New("strings cannot be zero-length"))
	}
	var sequence []int
	for i := 0; i < length; i++ {
		sequence = append(sequence, rand.Int()%2)
	}
	return sequence
}

// SetGenerateBitString sets the function that generates the bitstring population
func SetGenerateCandidate(f func(length int) []int) {
	GenerateCandidate = f
}
