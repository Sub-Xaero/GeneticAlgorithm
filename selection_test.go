package ga

import (
	"math/rand"
	"testing"
)

func TestSetSelectionFunc(t *testing.T) {
	t.Parallel()
	rand.Seed(3)
	var genA = NewGeneticAlgorithm()
	genA.SetOutputFunc(func(a ...interface{}) { t.Log(a...) })

	genA.SetSelectionFunc(func(Fitness FitnessFunction, genomes []Genome) []Genome {
		offspring := make([]Genome, 0)
		for range genomes {
			offspring = append(offspring, genomes[0].Copy())
		}
		return offspring
	})

	genA.Population = []Genome{
		{[]int{1, 1, 1, 1}},
		{[]int{0, 1, 1, 1}},
		{[]int{0, 0, 1, 1}},
		{[]int{0, 0, 0, 1}},
	}
	genA.Population = genA.Selection(genA.Fitness, genA.Population)

	expectedFitness := 4
	gotFitness := genA.AverageFitness(genA.Population)
	if expectedFitness != gotFitness {
		t.Error("Set selection function did not work.", "Expected:", expectedFitness, "Got:", gotFitness)
	} else {
		t.Log("Set selection function worked.", "Expected:", expectedFitness, "Got:", gotFitness)
	}
}

func TestTournament(t *testing.T) {
	t.Parallel()
	rand.Seed(3)
	var genA = NewGeneticAlgorithm()
	genA.SetOutputFunc(func(a ...interface{}) { t.Log(a...) })

	genA.Population = []Genome{
		{[]int{1, 1, 1, 1}},
		{[]int{0, 1, 1, 1}},
		{[]int{0, 0, 1, 1}},
		{[]int{0, 0, 0, 1}},
	}
	avgFitnessBefore := genA.AverageFitness(genA.Population)
	genA.Population = genA.Selection(genA.Fitness, genA.Population)
	avgFitnessAfter := genA.AverageFitness(genA.Population)

	if avgFitnessAfter < avgFitnessBefore {
		t.Error("Average Fitness decreased after tournament.", "Was:", avgFitnessBefore, "Now:", avgFitnessAfter)
	} else {
		t.Log("Average Fitness no worse after tournament.", "Was:", avgFitnessBefore, "Now:", avgFitnessAfter)
	}
}

func TestRoulette(t *testing.T) {
	t.Parallel()
	rand.Seed(3)
	var genA = NewGeneticAlgorithm()
	genA.SetOutputFunc(func(a ...interface{}) { t.Log(a...) })

	genA.Population = []Genome{
		{[]int{1, 1, 1, 1}},
		{[]int{0, 1, 1, 1}},
		{[]int{0, 0, 1, 1}},
		{[]int{0, 0, 0, 1}},
	}
	avgFitnessBefore := genA.AverageFitness(genA.Population)
	genA.Population = genA.Selection(genA.Fitness, genA.Population)
	avgFitnessAfter := genA.AverageFitness(genA.Population)

	if avgFitnessAfter < avgFitnessBefore {
		t.Error("Average Fitness decreased after tournament.", "Was:", avgFitnessBefore, "Now:", avgFitnessAfter)
	} else {
		t.Log("Average Fitness no worse after tournament.", "Was:", avgFitnessBefore, "Now:", avgFitnessAfter)
	}
}
