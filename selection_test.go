package ga

import (
	"math/rand"
	"testing"
	"time"
)

func TestSetSelectionFunc(t *testing.T) {
	rand.Seed(time.Now().Unix())
	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(TournamentSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)

	SetSelectionFunc(func(genomes []Genome) []Genome {
		offspring := make([]Genome, 0)
		for range genomes {
			offspring = append(offspring, genomes[0].Copy())
		}
		return offspring
	})

	population := []Genome{
		{[]int{1, 1, 1, 1}},
		{[]int{0, 1, 1, 1}},
		{[]int{0, 0, 1, 1}},
		{[]int{0, 0, 0, 1}},
	}
	population = Selection(population)

	expectedFitness := 4
	gotFitness := AverageFitness(population)
	if expectedFitness != gotFitness {
		t.Error("Set selection function did not work.", "Expected:", expectedFitness, "Got:", gotFitness)
	} else {
		t.Log("Set selection function worked.", "Expected:", expectedFitness, "Got:", gotFitness)
	}
}

func TestTournament(t *testing.T) {
	rand.Seed(3)
	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(TournamentSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)

	population := []Genome{
		{[]int{1, 1, 1, 1}},
		{[]int{0, 1, 1, 1}},
		{[]int{0, 0, 1, 1}},
		{[]int{0, 0, 0, 1}},
	}
	avgFitnessBefore := AverageFitness(population)
	population = Selection(population)
	avgFitnessAfter := AverageFitness(population)

	if avgFitnessAfter < avgFitnessBefore {
		t.Error("Average Fitness decreased after tournament.", "Was:", avgFitnessBefore, "Now:", avgFitnessAfter)
	} else {
		t.Log("Average Fitness no worse after tournament.", "Was:", avgFitnessBefore, "Now:", avgFitnessAfter)
	}
}

func TestRoulette(t *testing.T) {
	rand.Seed(3)
	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(RouletteSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)

	population := []Genome{
		{[]int{1, 1, 1, 1}},
		{[]int{0, 1, 1, 1}},
		{[]int{0, 0, 1, 1}},
		{[]int{0, 0, 0, 1}},
	}
	avgFitnessBefore := AverageFitness(population)
	population = Selection(population)
	avgFitnessAfter := AverageFitness(population)

	if avgFitnessAfter < avgFitnessBefore {
		t.Error("Average Fitness decreased after tournament.", "Was:", avgFitnessBefore, "Now:", avgFitnessAfter)
	} else {
		t.Log("Average Fitness no worse after tournament.", "Was:", avgFitnessBefore, "Now:", avgFitnessAfter)
	}
}
