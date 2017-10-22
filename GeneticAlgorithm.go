package ga

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"math/rand"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type GeneticAlgorithm struct {
	Population    []Genome
	BestCandidate Genome
	Generations   int

	IterationsSinceChange int

	GenerateCandidate GenerateCandidateFunction
	Crossover         CrossoverFunction
	Mutate            MutateFunction
	Fitness           FitnessFunction
	Selection         SelectFunction
	Output            func(a ...interface{})

	RandomEngine *rand.Rand
}

func NewGeneticAlgorithm() GeneticAlgorithm {
	var geneticAlgorithm GeneticAlgorithm
	geneticAlgorithm.SetGenerateCandidate(DefaultGenerateCandidate)
	geneticAlgorithm.SetCrossoverFunc(DefaultCrossoverFunc)
	geneticAlgorithm.SetMutateFunc(DefaultMutateFunc)
	geneticAlgorithm.SetFitnessFunc(DefaultFitnessFunc)
	geneticAlgorithm.SetSelectionFunc(TournamentSelection)
	geneticAlgorithm.SetOutputFunc(PrintToConsole)
	geneticAlgorithm.SetSeed(time.Now().Unix())
	return geneticAlgorithm
}

func (genA *GeneticAlgorithm) SetSeed(seed int64) {
	genA.RandomEngine = rand.New(rand.NewSource(seed))
}

func (genA *GeneticAlgorithm) UpdateBestCandidate(bestGeneration Genome) {
	if genA.Fitness(bestGeneration) > genA.Fitness(genA.BestCandidate) {
		genA.BestCandidate = bestGeneration.Copy()
		genA.IterationsSinceChange = 0
	}
}

func (genA *GeneticAlgorithm) FillRandomPopulation(populationSize, candidateLength int) []Genome {
	population := make([]Genome, 0)
	for len(population) < populationSize {
		bitstring, err := genA.GenerateCandidate(candidateLength, genA.RandomEngine)
		check(err)
		population = append(population, Genome{bitstring})
	}
	return population
}

func (genA *GeneticAlgorithm) Run(populationSize, bitstringLength, generations int, crossover, mutate, terminateEarly bool) error {

	if genA.GenerateCandidate == nil {
		return errors.New("generate func candidate is nil")
	}
	if genA.Crossover == nil {
		return errors.New("crossover func is nil")
	}
	if genA.Mutate == nil {
		return errors.New("mutate func is nil")
	}
	if genA.Fitness == nil {
		return errors.New("fitness func is nil")
	}
	if genA.Selection == nil {
		return errors.New("selection func is nil")
	}
	if genA.Output == nil {
		return errors.New("output func is nil")
	}
	if genA.RandomEngine == nil {
		return errors.New("random generator is not initialised")
	}

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
	genA.Population = make([]Genome, 0)
	genA.Population = genA.FillRandomPopulation(populationSize, bitstringLength)
	genA.UpdateBestCandidate(genA.MaxFitnessCandidate(genA.Population))

	// Run breeding cycles
	for y := 1; y <= generations; y++ {
		var bestCandidateOfGeneration Genome

		bestCandidateOfGeneration = genA.MaxFitnessCandidate(genA.Population)
		genA.UpdateBestCandidate(bestCandidateOfGeneration)
		genA.Output("Iteration", y)
		genA.Output("Start Population      :", genA.Population, "Average:", genA.AverageFitness(genA.Population), "Max:", genA.MaxFitness(genA.Population), "Best:", bestCandidateOfGeneration.Sequence)

		// Tournament
		breedingGround := make([]Genome, 0)
		breedingGround = append(breedingGround, genA.Selection(genA.Fitness, genA.Population, genA.RandomEngine)...)
		bestCandidateOfGeneration = genA.MaxFitnessCandidate(genA.Population)
		genA.UpdateBestCandidate(bestCandidateOfGeneration)
		genA.Output("Tournament Offspring  :", breedingGround, "Average:", genA.AverageFitness(breedingGround), "Max:", genA.MaxFitness(breedingGround), "Best:", bestCandidateOfGeneration.Sequence)

		// Crossover
		if crossover {
			crossoverBreedingGround := make([]Genome, 0)
			for i := 0; i+1 < len(breedingGround); i += 2 {
				newOffspring, err := genA.Crossover(breedingGround[i], breedingGround[i+1], genA.RandomEngine)
				check(err)
				crossoverBreedingGround = append(crossoverBreedingGround, newOffspring...)
			}
			breedingGround = crossoverBreedingGround
			bestCandidateOfGeneration = genA.MaxFitnessCandidate(genA.Population)
			genA.UpdateBestCandidate(bestCandidateOfGeneration)
			genA.Output("Crossover Offspring   :", breedingGround, "Average:", genA.AverageFitness(breedingGround), "Max:", genA.MaxFitness(breedingGround), "Best:", bestCandidateOfGeneration.Sequence)
		}

		// Mutation
		if mutate {
			for index := range breedingGround {
				breedingGround[index] = genA.Mutate(breedingGround[index], genA.RandomEngine)
			}
			bestCandidateOfGeneration = genA.MaxFitnessCandidate(genA.Population)
			genA.UpdateBestCandidate(bestCandidateOfGeneration)
			genA.Output("Mutation Offspring    :", breedingGround, "Average:", genA.AverageFitness(breedingGround), "Max:", genA.MaxFitness(breedingGround), "Best:", bestCandidateOfGeneration.Sequence)
		}

		genA.Generations++
		genA.IterationsSinceChange++
		genA.Population = make([]Genome, populationSize)
		copy(genA.Population, breedingGround)
		genA.Output("Final Population      :", breedingGround, "Average:", genA.AverageFitness(breedingGround), "Max:", genA.MaxFitness(breedingGround), "Best:", bestCandidateOfGeneration.Sequence)
		genA.Output()
		genA.Output()

		outputString := strings.Join([]string{strconv.Itoa(y), strconv.Itoa(genA.AverageFitness(genA.Population)), strconv.Itoa(genA.MaxFitness(genA.Population)), "\n"}, ",")
		f.WriteString(outputString)

		if terminateEarly && float32(genA.IterationsSinceChange) > float32(generations)*0.25 {
			genA.Output("Termination : Stagnating change")
			genA.Output("Best Candidate Found:", genA.BestCandidate.Sequence, "Fitness:", genA.Fitness(genA.BestCandidate))
			break
		}
	}

	genA.Output("Best Candidate Found:", genA.BestCandidate.Sequence, "Fitness:", genA.Fitness(genA.BestCandidate))
	return nil
}
