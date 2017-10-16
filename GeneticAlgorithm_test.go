package ga

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGeneticAlgorithm(t *testing.T) {
	rand.Seed(3)

	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(TournamentSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)

	fmt.Println(GeneticAlgorithm(10, 10, 50, true, true, false))
}

func TestFillRandomPopulation(t *testing.T) {
	rand.Seed(time.Now().Unix())
	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(TournamentSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)

	population := make([]Genome, 0)
	population = FillRandomPopulation(population, 10, 10)

	if len(population) != 10 {
		t.Error("Population not filled to specified size")
	}

	for _, val := range population {
		if len(val.Sequence) == 0 {
			t.Error("Bitstrings in population are empty")
		}
	}
}
