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
		return fmt.Sprintf("{%v, %3v}", gene.Sequence, gene.Fitness())
	} else {
		return fmt.Sprintf("%v", gene.Fitness())
	}
}

func FillRandomPopulation(population []Genome, populationSize, candidateLength int) []Genome {
	for len(population) < populationSize {
		population = append(population, Genome{GenerateCandidate(candidateLength)})
	}
	return population
}

func GeneticAlgorithm(populationSize, bitstringLength, generations int, crossover, mutate, terminateEarly bool) (Genome, []Genome) {
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
	population := make([]Genome, 0)
	population = FillRandomPopulation(population, populationSize, bitstringLength)

	bestCandidate := population[0]

	UpdateBestCandidate := func(bestOverall *Genome, bestGeneration Genome, iterationsSinceChange *int) {
		if bestGeneration.Fitness() > bestOverall.Fitness() {
			*bestOverall = bestGeneration
			*iterationsSinceChange = 0
		}
	}

	// Run breeding cycles
	for y := 1; y <= generations; y++ {
		var bestCandidateOfGeneration Genome

		bestCandidateOfGeneration = MaxFitnessCandidate(population)
		if bestCandidateOfGeneration.Fitness() > bestCandidate.Fitness() {
			bestCandidate = bestCandidateOfGeneration
			iterationsSinceChange = 0
		}
		fmt.Println("Iteration", y)
		fmt.Println("Start Population      :", population, "Average:", AverageFitness(population), "Max:", MaxFitness(population), "Best:", bestCandidateOfGeneration.Sequence)

		// Tournament
		breedingGround := make([]Genome, 0)
		breedingGround = append(breedingGround, Selection(population)...)
		bestCandidateOfGeneration = MaxFitnessCandidate(breedingGround)
		UpdateBestCandidate(&bestCandidate, bestCandidateOfGeneration, &iterationsSinceChange)
		fmt.Println("Tournament Offspring  :", breedingGround, "Average:", AverageFitness(breedingGround), "Max:", MaxFitness(breedingGround), "Best:", bestCandidateOfGeneration.Sequence)

		// Crossover
		if crossover {
			crossoverBreedingGround := make([]Genome, 0)
			for i := 0; i+1 < len(breedingGround); i += 2 {
				crossoverBreedingGround = append(crossoverBreedingGround, breedingGround[i].Crossover(breedingGround[i+1])...)
			}
			breedingGround = crossoverBreedingGround
			bestCandidateOfGeneration = MaxFitnessCandidate(breedingGround)
			UpdateBestCandidate(&bestCandidate, bestCandidateOfGeneration, &iterationsSinceChange)
			fmt.Println("Crossover Offspring   :", breedingGround, "Average:", AverageFitness(breedingGround), "Max:", MaxFitness(breedingGround), "Best:", bestCandidateOfGeneration.Sequence)
		}

		// Mutation
		if mutate {
			for index := range breedingGround {
				breedingGround[index] = breedingGround[index].Mutate()
			}
			bestCandidateOfGeneration = MaxFitnessCandidate(breedingGround)
			UpdateBestCandidate(&bestCandidate, bestCandidateOfGeneration, &iterationsSinceChange)
			fmt.Println("Mutation Offspring    :", breedingGround, "Average:", AverageFitness(breedingGround), "Max:", MaxFitness(breedingGround), "Best:", bestCandidateOfGeneration.Sequence)
		}

		iterationsSinceChange++
		population = make([]Genome, populationSize)
		copy(population, breedingGround)
		fmt.Println("Final Population      :", breedingGround, "Average:", AverageFitness(breedingGround), "Max:", MaxFitness(breedingGround), "Best:", bestCandidateOfGeneration.Sequence)
		fmt.Println()
		fmt.Println()

		outputString := strings.Join([]string{strconv.Itoa(y), strconv.Itoa(AverageFitness(population)), strconv.Itoa(MaxFitness(population)), "\n"}, ",")
		f.WriteString(outputString)

		if terminateEarly && float32(iterationsSinceChange) > float32(generations)*0.25 {
			fmt.Println("Termination : Stagnating change")
			fmt.Println("Best Candidate Found:", bestCandidate.Sequence, "Fitness:", bestCandidate.Fitness())
			return bestCandidate, population
		}
	}

	fmt.Println("Best Candidate Found:", bestCandidate.Sequence, "Fitness:", bestCandidate.Fitness())
	return bestCandidate, population
}
