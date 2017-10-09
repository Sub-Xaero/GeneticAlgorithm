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

var crossoverFunc func(gene, spouse Genome) []Genome = func(gene, spouse Genome) []Genome {
	offspring := make([]Genome, 0)

	if len(gene.Sequence) != len(spouse.Sequence) {
		panic(errors.New("strings are not current length"))
	}

	crossover := rand.Int() % len(gene.Sequence)

	offspring = append(offspring, Genome{gene.Sequence[0:crossover] + spouse.Sequence[crossover:]})
	offspring = append(offspring, Genome{spouse.Sequence[0:crossover] + gene.Sequence[crossover:]})
	return offspring
}

// SetCrossoverFunc changes the crossover function to the function specified
func SetCrossoverFunc(f func(gene, spouse Genome) []Genome) {
	crossoverFunc = f
}

// Crossover applies a function, set by SetCrossoverFunc to the receiver gene and a specified pair
func (gene Genome) Crossover(spouse Genome) []Genome {
	return crossoverFunc(gene, spouse)
}

var mutateFunc func(gene Genome, chance int) Genome = func(gene Genome, chance int) Genome {
	mutant := ""
	for _, i := range gene.Sequence {
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
	gene.Sequence = mutant
	return gene
}

// SetMutateFunc changes the mutate function to the function specified
func SetMutateFunc(f func(gene Genome, chance int) Genome) {
	mutateFunc = f
}

// Mutate returns a bitstring with bits mutated by a function set by SetMutateFunc
func (gene Genome) Mutate(chance int) Genome {
	return mutateFunc(gene, chance)
}

func (gene Genome) String() string {
	if len(gene.Sequence) <= 10 {
		return fmt.Sprintf("{%v, %3v}", gene.Sequence, gene.Fitness())
	} else {
		return fmt.Sprintf("%v", gene.Fitness())
	}
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
	bestCandidate := population[0]

	// Run breeding cycles
	for y := 1; y <= generations; y++ {
		var bestCandidateOfGeneration Genome

		bestCandidateOfGeneration = MaxFitnessCandidate(population)
		if bestCandidateOfGeneration.Fitness() > bestCandidate.Fitness() {
			bestCandidate = bestCandidateOfGeneration
		}
		fmt.Println("Iteration", y)
		fmt.Println("Start Population      :", population, "Average:", AverageFitness(population), "Max:", MaxFitness(population), "Best:", bestCandidateOfGeneration.Sequence)

		// Tournament
		breedingGround := make([]Genome, 0)
		breedingGround = append(breedingGround, Tournament(population)...)
		bestCandidateOfGeneration = MaxFitnessCandidate(breedingGround)
		if bestCandidateOfGeneration.Fitness() > bestCandidate.Fitness() {
			bestCandidate = bestCandidateOfGeneration
		}
		fmt.Println("Tournament Offspring  :", breedingGround, "Average:", AverageFitness(breedingGround), "Max:", MaxFitness(breedingGround), "Best:", bestCandidateOfGeneration.Sequence)

		// Crossover
		crossoverBreedingGround := make([]Genome, 0)
		for i := 0; i+1 < len(breedingGround); i += 2 {
			crossoverBreedingGround = append(crossoverBreedingGround, breedingGround[i].Crossover(breedingGround[i+1])...)
		}
		breedingGround = crossoverBreedingGround
		bestCandidateOfGeneration = MaxFitnessCandidate(breedingGround)
		if bestCandidateOfGeneration.Fitness() > bestCandidate.Fitness() {
			bestCandidate = bestCandidateOfGeneration
		}
		fmt.Println("Crossover Offspring   :", breedingGround, "Average:", AverageFitness(breedingGround), "Max:", MaxFitness(breedingGround), "Best:", bestCandidateOfGeneration.Sequence)

		// Mutation
		for index := range breedingGround {
			breedingGround[index] = breedingGround[index].Mutate(mutateChance)
		}
		bestCandidateOfGeneration = MaxFitnessCandidate(breedingGround)
		if bestCandidateOfGeneration.Fitness() > bestCandidate.Fitness() {
			bestCandidate = bestCandidateOfGeneration
		}
		fmt.Println("Mutation Offspring    :", breedingGround, "Average:", AverageFitness(breedingGround), "Max:", MaxFitness(breedingGround), "Best:", bestCandidateOfGeneration.Sequence)

		population = make([]Genome, populationSize)
		copy(population, breedingGround)
		fmt.Println()
		fmt.Println()

		outputString := strings.Join([]string{strconv.Itoa(y), strconv.Itoa(AverageFitness(population)), strconv.Itoa(MaxFitness(population)), "\n"}, ",")
		f.WriteString(outputString)
	}

	fmt.Println("Best Candidate Found:", bestCandidate.Sequence, "Fitness:", bestCandidate.Fitness())
	return population
}
