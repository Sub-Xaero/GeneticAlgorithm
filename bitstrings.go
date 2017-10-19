package ga

import (
	"errors"
	"math/rand"
)

type GenerateCandidateFunction func(length int) ([]int, error)

// GenerateBitString returns an encoded string as set by calls SetGenerateBitString. Defaults to binary strings
var DefaultGenerateCandidate GenerateCandidateFunction = func(length int) ([]int, error) {
	if length <= 0 {
		return nil, errors.New("strings cannot be zero-length")
	}
	var sequence []int
	for i := 0; i < length; i++ {
		sequence = append(sequence, rand.Int()%2)
	}
	return sequence, nil
}

// SetGenerateBitString sets the function that generates the bitstring population
func (genA *GeneticAlgorithm) SetGenerateCandidate(f GenerateCandidateFunction) {
	genA.GenerateCandidate = f
}
