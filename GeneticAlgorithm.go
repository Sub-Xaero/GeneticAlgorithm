package ga

import (
	"fmt"
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
	Sequence []int
}

func (gene Genome) Copy() Genome {
	sequence := make([]int, len(gene.Sequence))
	copy(sequence, gene.Sequence)
	return Genome{sequence}
}

func (gene Genome) String() string {
	if len(gene.Sequence) <= 10 {
		return fmt.Sprintf("{%v, %3v}", gene.Sequence, Fitness(gene))
	} else {
		return fmt.Sprintf("%v", Fitness(gene))
	}
}

func FillRandomPopulation(population []Genome, populationSize, candidateLength int) []Genome {
	for len(population) < populationSize {
		bitstring, err := GenerateCandidate(candidateLength)
		check(err)

		population = append(population, Genome{bitstring})
	}
	return population
}

func GeneticAlgorithm(populationSize, bitstringLength, generations int, crossover, mutate, terminateEarly bool) (bestCandidate Genome, numGenerations int, population []Genome) {
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

	iterationsSinceChange := 0

	// Init
	population = make([]Genome, 0)
	population = FillRandomPopulation(population, populationSize, bitstringLength)

	bestCandidate = population[0]

	UpdateBestCandidate := func(bestOverall *Genome, bestGeneration Genome, iterationsSinceChange *int) {
		if Fitness(bestGeneration) > Fitness(*bestOverall) {
			*bestOverall = bestGeneration
			*iterationsSinceChange = 0
		}
	}

	// Run breeding cycles
	for y := 1; y <= generations; y++ {
		var bestCandidateOfGeneration Genome

		bestCandidateOfGeneration = MaxFitnessCandidate(population)
		UpdateBestCandidate(&bestCandidate, bestCandidateOfGeneration, &iterationsSinceChange)
		Output("Iteration", y)
		Output("Start Population      :", population, "Average:", AverageFitness(population), "Max:", MaxFitness(population), "Best:", bestCandidateOfGeneration.Sequence)

		// Tournament
		breedingGround := make([]Genome, 0)
		breedingGround = append(breedingGround, Selection(population)...)
		bestCandidateOfGeneration = MaxFitnessCandidate(breedingGround)
		UpdateBestCandidate(&bestCandidate, bestCandidateOfGeneration, &iterationsSinceChange)
		Output("Tournament Offspring  :", breedingGround, "Average:", AverageFitness(breedingGround), "Max:", MaxFitness(breedingGround), "Best:", bestCandidateOfGeneration.Sequence)

		// Crossover
		if crossover {
			crossoverBreedingGround := make([]Genome, 0)
			for i := 0; i+1 < len(breedingGround); i += 2 {
				newOffspring, err := Crossover(breedingGround[i], breedingGround[i+1])
				check(err)
				crossoverBreedingGround = append(crossoverBreedingGround, newOffspring...)
			}
			breedingGround = crossoverBreedingGround
			bestCandidateOfGeneration = MaxFitnessCandidate(breedingGround)
			UpdateBestCandidate(&bestCandidate, bestCandidateOfGeneration, &iterationsSinceChange)
			Output("Crossover Offspring   :", breedingGround, "Average:", AverageFitness(breedingGround), "Max:", MaxFitness(breedingGround), "Best:", bestCandidateOfGeneration.Sequence)
		}

		// Mutation
		if mutate {
			for index := range breedingGround {
				breedingGround[index] = Mutate(breedingGround[index])
			}
			bestCandidateOfGeneration = MaxFitnessCandidate(breedingGround)
			UpdateBestCandidate(&bestCandidate, bestCandidateOfGeneration, &iterationsSinceChange)
			Output("Mutation Offspring    :", breedingGround, "Average:", AverageFitness(breedingGround), "Max:", MaxFitness(breedingGround), "Best:", bestCandidateOfGeneration.Sequence)
		}

		numGenerations++
		iterationsSinceChange++
		population = make([]Genome, populationSize)
		copy(population, breedingGround)
		Output("Final Population      :", breedingGround, "Average:", AverageFitness(breedingGround), "Max:", MaxFitness(breedingGround), "Best:", bestCandidateOfGeneration.Sequence)
		Output()
		Output()

		outputString := strings.Join([]string{strconv.Itoa(y), strconv.Itoa(AverageFitness(population)), strconv.Itoa(MaxFitness(population)), "\n"}, ",")
		f.WriteString(outputString)

		if terminateEarly && float32(iterationsSinceChange) > float32(generations)*0.25 {
			Output("Termination : Stagnating change")
			Output("Best Candidate Found:", bestCandidate.Sequence, "Fitness:", Fitness(bestCandidate))
			return bestCandidate, numGenerations, population
		}
	}

	Output("Best Candidate Found:", bestCandidate.Sequence, "Fitness:", Fitness(bestCandidate))
	return bestCandidate, numGenerations, population
}
