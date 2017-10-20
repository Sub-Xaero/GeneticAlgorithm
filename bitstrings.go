package ga

import (
	"errors"
	"math/rand"
	"strconv"
)

type bitstring []string

type GenerateCandidateFunction func(length int) (bitstring, error)

// GenerateBitString returns an encoded string as set by calls SetGenerateBitString. Defaults to binary strings
var DefaultGenerateCandidate GenerateCandidateFunction = func(length int) (bitstring, error) {
	if length <= 0 {
		return nil, errors.New("strings cannot be zero-length")
	}
	var sequence bitstring
	for i := 0; i < length; i++ {
		sequence = append(sequence, strconv.Itoa(rand.Int()%2))
	}
	return sequence, nil
}

// SetGenerateBitString sets the function that generates the bitstring population
func (genA *GeneticAlgorithm) SetGenerateCandidate(f GenerateCandidateFunction) {
	genA.GenerateCandidate = f
}
