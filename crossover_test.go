package ga

import (
	"fmt"
	"testing"
)

func TestCrossover(t *testing.T) {
	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(TournamentSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)

	population := []Genome{
		{[]int{1, 0, 0, 0}},
		{[]int{0, 0, 0, 1}},
	}
	offspring := population[0].Crossover(population[1])
	found := false
	foundIndex := 0

	expectedString := "{[1 0 0 1],   2}"
	for i, val := range offspring {
		if fmt.Sprint(val) == expectedString {
			found = true
			foundIndex = i
			break
		}
	}
	if !found {
		t.Error("Crossover failed.", "Expected:", expectedString, "Got:", offspring)
	} else {
		t.Log("Crossover succeeded.", "Expected:", expectedString, "Got:", offspring[foundIndex])
	}
}

func TestSetCrossoverFunc(t *testing.T) {
	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(TournamentSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)

	SetCrossoverFunc(func(gene, spouse Genome) []Genome {
		return []Genome{{[]int{1, 2, 3, 4}}}
	})

	expectedString := "[{[1 2 3 4],   1}]"
	gotString := fmt.Sprint(Genome{[]int{}}.Crossover(Genome{[]int{}}))
	if gotString != expectedString {
		t.Error("Crossover function not set.", "Expected:", expectedString, "Got:", gotString)
	} else {
		t.Log("Crossover function set successfully.", "Expected:", expectedString, "Got:", gotString)
	}
}
