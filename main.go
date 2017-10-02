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

func (self Genome) fitness() int64 {
	var score int64
	score, err := strconv.ParseInt(self.sequence, 2, 32)
	if err == nil {
		return score
	} else {
		panic(errors.New("could not parse bitstring"))
	}
}

func (self Genome) crossover(spouse Genome) []Genome {
	offspring := make([]Genome, 0)

	if len(self.sequence) != len(spouse.sequence) {
		panic(errors.New("strings are not current length"))
	}

	crossover := rand.Int() % len(self.sequence)

	offspring = append(offspring, Genome{self.sequence[0:crossover] + spouse.sequence[crossover:]})
	offspring = append(offspring, Genome{spouse.sequence[0:crossover] + self.sequence[crossover:]})
	return offspring
}

func (self *Genome) mutate(chance int) Genome {
	mutant := ""
	for _, i := range self.sequence {
		if rand.Int()%chance == 1 {
			if string(i) == "1" {
				mutant += "0"
			} else {
				mutant += "1"
			}
		} else {
			mutant += string(i)
		}
	}
	self.sequence = mutant
	return *self
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
