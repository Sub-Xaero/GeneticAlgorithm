package ga

import (
	"fmt"
	"testing"
)

func TestCrossover(t *testing.T) {
	SetCrossoverFunc(DefaultCrossoverFunc)

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
		t.Error("Crossover failed", offspring)
	}
}

func TestSetCrossoverFunc(t *testing.T) {
	SetCrossoverFunc(func(gene, spouse Genome) []Genome {
		return []Genome{{[]int{1, 2, 3, 4}}}
	})

	if fmt.Sprint(Genome{[]int{}}.Crossover(Genome{[]int{}})) != "[{[1 2 3 4],   1}]" {
		t.Error("Crossover function not set")
	}
}
