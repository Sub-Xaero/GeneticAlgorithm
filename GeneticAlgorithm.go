package ga

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Genome represents a bitstring and associated fitness value
type Genome struct {
	Sequence string
}

// Crossover returns bitstring pair which is product of two bitstrings with their tails swapped at a random index
func (gene Genome) Crossover(spouse Genome) []Genome {
	offspring := make([]Genome, 0)

	if len(gene.Sequence) != len(spouse.Sequence) {
		panic(errors.New("strings are not current length"))
	}

	crossover := rand.Int() % len(gene.Sequence)

	offspring = append(offspring, Genome{gene.Sequence[0:crossover] + spouse.Sequence[crossover:]})
	offspring = append(offspring, Genome{spouse.Sequence[0:crossover] + gene.Sequence[crossover:]})
	return offspring
}

// Mutate returns a bitstring with bits flipped at chance 1/n
func (gene Genome) Mutate(n int) Genome {
	mutant := ""
	for _, i := range gene.Sequence {
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
	gene.Sequence = mutant
	return gene
}

func (gene Genome) String() string {
	if len(gene.Sequence) <= 10 {
		return fmt.Sprintf("{%v, %3v}", gene.Sequence, gene.Fitness())
	} else {
		return fmt.Sprintf("%v", gene.Fitness())
	}
}

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

func FillRandomPopulation(population []Genome, populationSize, bitstringLength int) []Genome {
	for len(population) < populationSize {
		population = append(population, Genome{GenerateBitString(bitstringLength)})
	}
	return population
}

func GeneticAlgorithm(populationSize, bitstringLength, generations, mutateChance int) []Genome {
	// Open output file, to save results to
	outputFile := "output.txt"
	_, err := os.Stat(outputFile)
	if !os.IsNotExist(err) {
		os.Remove(outputFile)
	}
	f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	check(err)
	defer f.Close()
	defer f.Sync()

	outputString := strings.Join([]string{"Iteration", "AverageFitness", "MaxFitness", "\n"}, ",")
	f.WriteString(outputString)

	// Init
	population := make([]Genome, 0)
	population = FillRandomPopulation(population, populationSize, bitstringLength)

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

		population = make([]Genome, populationSize)
		copy(population, breedingGround)
		fmt.Println()
		fmt.Println()

		outputString := strings.Join([]string{strconv.Itoa(y), strconv.Itoa(AverageFitness(population)), strconv.Itoa(MaxFitness(population)), "\n"}, ",")
		f.WriteString(outputString)
	}

	return population
}
