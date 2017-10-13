package ga

import (
	"fmt"
	"testing"
)


func TestTournament(t *testing.T) {
	population := []Genome{
		{[]int{1, 1, 1, 1}},
		{[]int{0, 1, 1, 1}},
		{[]int{0, 0, 1, 1}},
		{[]int{0, 0, 0, 1}},
	}
	avgFitnessBefore := AverageFitness(population)
	population = Tournament(population)
	avgFitnessAfter := AverageFitness(population)

	if avgFitnessAfter < avgFitnessBefore {
		t.Error("Fitness decreased after tournament")
	}
}

func TestCrossover(t *testing.T) {
	population := []Genome{
		{[]int{1, 0, 0, 0}},
		{[]int{0, 0, 0, 1}},
	}
	offspring := population[0].Crossover(population[1])
	found := false
	for _, val := range offspring {
		if fmt.Sprint(val) == "{[1 0 0 1],   2}" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Crossover failed")
	}
}
