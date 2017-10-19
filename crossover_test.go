package ga

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestDefaultCrossover(t *testing.T) {
	t.Parallel()
	rand.Seed(3)
	var genA = NewGeneticAlgorithm()
	genA.SetOutputFunc(func(a ...interface{}) { t.Log(a...) })
	population := []Genome{
		{[]int{1, 0, 0, 0}},
		{[]int{0, 0, 0, 1}},
	}
	offspring, err := genA.Crossover(population[0], population[1])

	if err != nil {
		t.Error("Unexpected error:", err)
	} else {
		t.Log("Crossover threw no errors")
	}

	found := false
	foundIndex := 0

	expectedString := "{[1 0 0 1]}"
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

func TestBadDefaultCrossover(t *testing.T) {
	t.Parallel()
	rand.Seed(3)
	var genA = NewGeneticAlgorithm()
	genA.SetOutputFunc(func(a ...interface{}) { t.Log(a...) })

	population := []Genome{
		{[]int{1, 0, 0, 0}},
		{[]int{0, 0, 0}},
	}
	_, err := genA.Crossover(population[0], population[1])
	if err == nil {
		t.Error("Expected error but got:", err)
	} else {
		t.Log("Successfuly threw and caught err:", err)
	}
}

func TestSetCrossoverFunc(t *testing.T) {
	t.Parallel()
	rand.Seed(3)
	var genA = NewGeneticAlgorithm()
	genA.SetOutputFunc(func(a ...interface{}) { t.Log(a...) })

	genA.SetCrossoverFunc(func(gene, spouse Genome) ([]Genome, error) {
		return []Genome{{[]int{1, 2, 3, 4}}}, nil
	})

	expectedString := "[{[1 2 3 4]}]"
	crossoverGene, err := genA.Crossover(Genome{[]int{}}, Genome{[]int{}})

	if err != nil {
		t.Error("Unexpected error:", err)
	} else {
		t.Log("Crossover threw no errors")
	}

	gotString := fmt.Sprint(crossoverGene)
	if gotString != expectedString {
		t.Error("Crossover function not set.", "Expected:", expectedString, "Got:", gotString)
	} else {
		t.Log("Crossover function set successfully.", "Expected:", expectedString, "Got:", gotString)
	}
}
