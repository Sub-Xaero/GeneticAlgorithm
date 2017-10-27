package ga

import (
	"math/rand"
	"testing"
)

func TestSetSelectionFunc(t *testing.T) {
	t.Parallel()
	var genA = NewGeneticAlgorithm()
	genA.SetSeed(3)
	genA.SetOutputFunc(func(a ...interface{}) { t.Log(a...) })

	genA.SetSelectionFunc(func(Fitness FitnessFunction, genomes Population, random *rand.Rand) Population {
		offspring := make(Population, 0)
		for range genomes {
			offspring = append(offspring, genomes[0].Copy())
		}
		return offspring
	})

	genA.Candidates = Population{
		{Bitstring{"1", "1", "1", "1"}},
		{Bitstring{"0", "1", "1", "1"}},
		{Bitstring{"0", "0", "1", "1"}},
		{Bitstring{"0", "0", "0", "1"}},
	}
	genA.Candidates = genA.Selection(genA.Fitness, genA.Candidates, genA.RandomEngine)

	expectedFitness := 4
	gotFitness := genA.AverageFitness(genA.Candidates)
	if expectedFitness != gotFitness {
		t.Error("Set selection function did not work.", "Expected:", expectedFitness, "Got:", gotFitness)
	} else {
		t.Log("Set selection function worked.", "Expected:", expectedFitness, "Got:", gotFitness)
	}
}

func TestTournament(t *testing.T) {
	t.Parallel()
	var genA = NewGeneticAlgorithm()
	genA.SetSeed(3)
	genA.SetOutputFunc(func(a ...interface{}) { t.Log(a...) })

	genA.Candidates = Population{
		{Bitstring{"1", "1", "1", "1"}},
		{Bitstring{"0", "1", "1", "1"}},
		{Bitstring{"0", "0", "1", "1"}},
		{Bitstring{"0", "0", "0", "1"}},
	}
	avgFitnessBefore := genA.AverageFitness(genA.Candidates)
	genA.Candidates = genA.Selection(genA.Fitness, genA.Candidates, genA.RandomEngine)
	avgFitnessAfter := genA.AverageFitness(genA.Candidates)

	if avgFitnessAfter < avgFitnessBefore {
		t.Error("Average Fitness decreased after tournament.", "Was:", avgFitnessBefore, "Now:", avgFitnessAfter)
	} else {
		t.Log("Average Fitness no worse after tournament.", "Was:", avgFitnessBefore, "Now:", avgFitnessAfter)
	}
}

func TestRoulette(t *testing.T) {
	t.Parallel()
	var genA = NewGeneticAlgorithm()
	genA.SetSeed(3)
	genA.SetOutputFunc(func(a ...interface{}) { t.Log(a...) })

	genA.Candidates = Population{
		{Bitstring{"1", "1", "1", "1"}},
		{Bitstring{"0", "1", "1", "1"}},
		{Bitstring{"0", "0", "1", "1"}},
		{Bitstring{"0", "0", "0", "1"}},
	}
	avgFitnessBefore := genA.AverageFitness(genA.Candidates)
	genA.Candidates = genA.Selection(genA.Fitness, genA.Candidates, genA.RandomEngine)
	avgFitnessAfter := genA.AverageFitness(genA.Candidates)

	if avgFitnessAfter < avgFitnessBefore {
		t.Error("Average Fitness decreased after tournament.", "Was:", avgFitnessBefore, "Now:", avgFitnessAfter)
	} else {
		t.Log("Average Fitness no worse after tournament.", "Was:", avgFitnessBefore, "Now:", avgFitnessAfter)
	}
}
