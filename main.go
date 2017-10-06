package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var (
	mutateChance = 100
	numStrings   = 10
	strLength    = 8
)

// Genome represents a bitstring and associated fitness value
type Genome struct {
	sequence string
}

// Fitness calculates the suitability of a candidate solution and returns an integral score value
func (gene Genome) Fitness() int {
	return strings.Count(gene.sequence, "1")
}

// Crossover returns bitstring pair which is product of two bitstrings with their tails swapped at a random index
func (gene Genome) Crossover(spouse Genome) []Genome {
	offspring := make([]Genome, 0)

	if len(gene.sequence) != len(spouse.sequence) {
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

// GenerateBitString returns the highest fitness found in a [] Genome population
func GenerateBitString(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("strings cannot be zero-length")
	}
	var bitstring string
	for i := 0; i < length; i++ {
		bitstring += strconv.Itoa(rand.Int() % 2)
	}
	return bitstring, nil
}

func rouletteSelect(populace []Genome, weightSum float64) Genome {
	var value float64 = rand.Float64() * weightSum

	for _, x := range populace {
		value -= float64(x.fitness())
		if value <= 0 {
			return x
		}
	}
	return populace[len(populace)-1]
}

func roulette(populace []Genome, numParents int) []Genome {
	breedingParents := make([]Genome, numParents)

	var weightSum float64 = 0
	for _, x := range populace {
		weightSum += float64(x.fitness())
	}

	for i := 0; i < numParents; i++ {
		breedingParents[i] = rouletteSelect(populace, weightSum)
	}

	return breedingParents
}

func fillRandomPopulation(populace []Genome) []Genome {
	for len(populace) < numStrings {
		str, err := generateBitString(strLength)
		if err == nil {
			populace = append(populace, Genome{str})
		} else {
			panic(errors.New("failed to initialise population"))
		}
	}
	fmt.Println("Initial population:", populace)
	return populace
}

func main() {
	rand.Seed(time.Now().Unix())

	// Init
	population := make([]Genome, 0)
	population = fillRandomPopulation(population)

	for y := 0; y < 100; y++ {
		fmt.Println("Interation", y)
		fmt.Println("Start Population:", population)

		breedingGround := make([]Genome, 0)
		breedingGround = append(breedingGround, roulette(population, numStrings/2)...)

		fmt.Println("Roulette:", breedingGround)

		crossoverBreedingGround := make([]Genome, 0)
		for i := 0; i+1 < len(breedingGround); i += 2 {
			crossoverBreedingGround = append(crossoverBreedingGround, breedingGround[i].crossover(breedingGround[i+1])...)
		}
		breedingGround = crossoverBreedingGround
		fmt.Println("Crossover:", breedingGround)

		for _, i := range breedingGround {
			i.mutate(mutateChance)
		}
		fmt.Println("Mutation:", breedingGround)

		population = make([]Genome, 0)
		copy(population, breedingGround)
		population = fillRandomPopulation(population)
		fmt.Println("Fill Population:", population)
		fmt.Println()
	}
}
