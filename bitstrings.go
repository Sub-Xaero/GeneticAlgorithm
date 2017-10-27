package ga

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

type Bitstring []string

type GenerateCandidateFunction func(int, *rand.Rand) (Bitstring, error)

// GenerateBitString returns an encoded string as set by calls SetGenerateBitString. Defaults to binary strings
var DefaultGenerateCandidate GenerateCandidateFunction = func(length int, random *rand.Rand) (Bitstring, error) {
	if length <= 0 {
		return nil, errors.New("strings cannot be zero-length")
	}
	var sequence Bitstring
	for i := 0; i < length; i++ {
		sequence = append(sequence, strconv.Itoa(random.Int()%2))
	}
	return sequence, nil
}

// SetGenerateBitString sets the function that generates the Bitstring candidatePool
func (genA *GeneticAlgorithm) SetGenerateCandidate(f GenerateCandidateFunction) {
	genA.GenerateCandidate = f
}

// SetGenerateBitString sets the function that generates the Bitstring candidatePool
func (b Bitstring) String() string {
	output := ""
	for _, val := range b {
		output += fmt.Sprintf("%v", val) + " "
	}
	return "[" + output + "]"
}
