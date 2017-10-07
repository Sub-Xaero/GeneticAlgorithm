package ga

import (
	"errors"
	"math/rand"
	"strconv"
)


// GenerateBitString returns an encoded string as set by calls SetGenerateBitString. Defaults to binary strings
var GenerateBitString func(int) string = func(length int) string {
	if length <= 0 {
		panic(errors.New("strings cannot be zero-length"))
	}
	var bitstring string
	for i := 0; i < length; i++ {
		bitstring += strconv.Itoa(rand.Int() % 2)
	}
	return bitstring
}

// SetGenerateBitString sets the function that generates the bitstring population
func SetGenerateBitString(f func(length int) string) {
	GenerateBitString = f
}
