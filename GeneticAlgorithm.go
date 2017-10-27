package ga

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type GeneticAlgorithm struct {
	Candidates    Population
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

func (genA *GeneticAlgorithm) FillRandomPopulation(populationSize, candidateLength int) Population {
	candidatePool := make(Population, 0)
	for len(candidatePool) < populationSize {
		bitstring, err := genA.GenerateCandidate(candidateLength, genA.RandomEngine)
		check(err)
		candidatePool = append(candidatePool, Genome{bitstring})
	}
	return candidatePool
}

func (genA *GeneticAlgorithm) Summarise(title string, candidatePool Population) {
	output := ""
	output += title
	output += "{"
	for _, val := range candidatePool {
		output += "["
		if len(val.Sequence) <= 10 {
			output += val.Sequence.String()
		} else {
			output += fmt.Sprintf("%3v", genA.Fitness(val))
		}
		output += "]"
	}
	output += "}"
	output += " Max : " + strconv.Itoa(genA.MaxFitness(candidatePool))
	output += " Average : " + strconv.Itoa(genA.AverageFitness(candidatePool))
	output += " Best : " + genA.MaxFitnessCandidate(candidatePool).String()
	genA.Output(output)
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

	// Init
	genA.Candidates = make(Population, 0)
	genA.Candidates = genA.FillRandomPopulation(populationSize, bitstringLength)
	genA.UpdateBestCandidate(genA.MaxFitnessCandidate(genA.Candidates))

	// Run breeding cycles
	for y := 1; y <= generations; y++ {
		var bestCandidateOfGeneration Genome

		bestCandidateOfGeneration = genA.MaxFitnessCandidate(genA.Candidates)
		genA.UpdateBestCandidate(bestCandidateOfGeneration)
		genA.Output("Iteration", y)
		genA.Summarise("Start Population      :", genA.Candidates)

		// Tournament
		breedingGround := make(Population, 0)
		breedingGround = append(breedingGround, genA.Selection(genA.Fitness, genA.Candidates, genA.RandomEngine)...)
		bestCandidateOfGeneration = genA.MaxFitnessCandidate(genA.Candidates)
		genA.UpdateBestCandidate(bestCandidateOfGeneration)
		genA.Summarise("Tournament Offspring  :", breedingGround)

		// Crossover
		if crossover {
			crossoverBreedingGround := make(Population, 0)
			for i := 0; i+1 < len(breedingGround); i += 2 {
				newOffspring, err := genA.Crossover(breedingGround[i], breedingGround[i+1], genA.RandomEngine)
				check(err)
				crossoverBreedingGround = append(crossoverBreedingGround, newOffspring...)
			}
			breedingGround = crossoverBreedingGround
			bestCandidateOfGeneration = genA.MaxFitnessCandidate(genA.Candidates)
			genA.UpdateBestCandidate(bestCandidateOfGeneration)
			genA.Summarise("Crossover Offspring   :", breedingGround)
		}

		// Mutation
		if mutate {
			for index := range breedingGround {
				breedingGround[index] = genA.Mutate(breedingGround[index], genA.RandomEngine)
			}
			bestCandidateOfGeneration = genA.MaxFitnessCandidate(genA.Candidates)
			genA.UpdateBestCandidate(bestCandidateOfGeneration)
			genA.Summarise("Mutation Offspring    :", breedingGround)
		}

		genA.Generations++
		genA.IterationsSinceChange++
		genA.Candidates = make(Population, populationSize)
		copy(genA.Candidates, breedingGround)
		genA.Summarise("Final Population      :", breedingGround)
		genA.Output()
		genA.Output()

		if terminateEarly && float32(genA.IterationsSinceChange) > float32(generations)*0.25 {
			genA.Output("Termination : Stagnating change")
			genA.Output("Best Candidate Found:", genA.BestCandidate.Sequence, "Fitness:", genA.Fitness(genA.BestCandidate))
			break
		}
	}

	genA.Output("Best Candidate Found:", genA.BestCandidate.Sequence, "Fitness:", genA.Fitness(genA.BestCandidate))
	return nil
}
