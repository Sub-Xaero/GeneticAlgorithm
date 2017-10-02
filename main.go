package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Genome struct {
	sequence string
}

func (g Genome) fitness() int64 {
	var score int64
	score, err := strconv.ParseInt(g.sequence, 2, 32)
	if err == nil {
		return score
	} else {
		panic(errors.New("could not parse bitstring"))
	}
}

func generateBitString(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("strings cannot be zero-length")
	}
	var bitstring string
	for i := 0; i < length; i++ {
		bitstring += strconv.Itoa(rand.Int() % 2)
	}
	return bitstring, nil
}

func main() {
	rand.Seed(time.Now().Unix())

}
