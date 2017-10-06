package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var (
	globalChance = 10
	mutateChance = globalChance
	numStrings   = globalChance
	strLength    = globalChance
	generations  = globalChance
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

	crossover := rand.Int() % len(gene.sequence)

	offspring = append(offspring, Genome{gene.sequence[0:crossover] + spouse.sequence[crossover:]})
	offspring = append(offspring, Genome{spouse.sequence[0:crossover] + gene.sequence[crossover:]})
	return offspring
}

// Mutate returns a bitstring with bits flipped at chance 1/n
func (gene Genome) Mutate(n int) Genome {
	mutant := ""
	for _, i := range gene.sequence {
		if rand.Int()%n == 1 {
			if string(i) == "1" {
				mutant += "0"
			} else {
				mutant += "1"
			}
		} else {
			mutant += string(i)
		}
	}
	gene.sequence = mutant
	return gene
}

func (gene Genome) String() string {
	if len(gene.sequence) <= 10 {
		return fmt.Sprintf("{%v, %3v}", gene.sequence, gene.Fitness())
	} else {
		return fmt.Sprintf("%v", gene.Fitness())
	}
}

// ByFitness is a receiver type that implements Sort for Genome []
type ByFitness []Genome

func (a ByFitness) Len() int           { return len(a) }
func (a ByFitness) Swap(x, y int)      { a[x], a[y] = a[y], a[x] }
func (a ByFitness) Less(x, y int) bool { return a[x].Fitness() < a[y].Fitness() }

// AverageFitness returns the average fitness of a [] Genome population
func AverageFitness(population []Genome) int {
	var average int = 0
	for _, i := range population {
		average += i.Fitness()
	}
	return average / int(len(population))
}

// MaxFitness returns the highest fitness found in a [] Genome population
func MaxFitness(population []Genome) int {
	var max int = 0
	for _, i := range population {
		if i.Fitness() > max {
			max = i.Fitness()
		}
	}
	return max
}

// GenerateBitString returns the highest fitness found in a [] Genome population
func GenerateBitString(length int) string {
	if length <= 0 {
		panic(errors.New("strings cannot be zero-length"))
	}
	var bitstring string
	for i := 0; i < length; i++ {
		bitstring += strconv.Itoa(rand.Int() % 2)
	}
	return bitstring
}

// Tournament returns a [] Genome population composed of the best out of randomly selected pairs
func Tournament(population []Genome) []Genome {
	offspring := make([]Genome, 0)

	for i := 0; i < len(population); i++ {
		parent1 := population[rand.Int()%len(population)]
		parent2 := population[rand.Int()%len(population)]

		if parent1.Fitness() > parent2.Fitness() {
			offspring = append(offspring, parent1)
		} else {
			offspring = append(offspring, parent2)
		}
	}

	return offspring
}

func FillRandomPopulation(population []Genome) []Genome {
	for len(population) < numStrings {
		population = append(population, Genome{GenerateBitString(strLength)})
	}
	return population
}

func main() {
	rand.Seed(time.Now().Unix())

	// Init
	population := make([]Genome, 0)
	population = FillRandomPopulation(population)

	// Run breeding cycles
	for y := 1; y <= generations; y++ {
		fmt.Println("Iteration", y)
		fmt.Println("Start Population      :", population, "Average:", AverageFitness(population), "Max:", MaxFitness(population))

		breedingGround := make([]Genome, 0)
		breedingGround = append(breedingGround, Tournament(population)...)
		fmt.Println("Tournament Offspring  :", breedingGround, "Average:", AverageFitness(breedingGround), "Max:", MaxFitness(breedingGround))

		crossoverBreedingGround := make([]Genome, 0)
		for i := 0; i+1 < len(breedingGround); i += 2 {
			crossoverBreedingGround = append(crossoverBreedingGround, breedingGround[i].Crossover(breedingGround[i+1])...)
		}
		breedingGround = crossoverBreedingGround
		fmt.Println("Crossover Offspring   :", breedingGround, "Average:", AverageFitness(breedingGround), "Max:", MaxFitness(breedingGround))

		for index := range breedingGround {
			breedingGround[index] = breedingGround[index].Mutate(mutateChance)
		}
		fmt.Println("Mutation Offspring    :", breedingGround, "Average:", AverageFitness(breedingGround), "Max:", MaxFitness(breedingGround))

		population = make([]Genome, numStrings)
		copy(population, breedingGround)
		fmt.Println()
		fmt.Println()
	}
}
