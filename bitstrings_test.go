package ga

import (
	"sort"
	"strings"
	"testing"
)

func TestSort(t *testing.T) {
	SetFitnessFunc(func(gene Genome) int {
		return strings.Count(gene.Sequence, "1")
	})

	population := []Genome{
		{"111111"},
		{"011111"},
		{"001111"},
		{"000111"},
		{"000011"},
		{"000001"},
	}

	after := []Genome{
		{"111111"},
		{"011111"},
		{"001111"},
		{"000111"},
		{"000011"},
		{"000001"},
	}
	sort.Sort(ByFitness(population))

	if len(population) != len(after) {
		t.Error("Populations do not match length")
	}

	for i := range population {
		if population[i] != after[len(after)-i-1] {
			t.Error("Not sorted")
			break
		}
	}

}

