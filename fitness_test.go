package ga

import (
	"strings"
	"testing"
)

func TestGenome_Fitness(t *testing.T) {
	genome := Genome{"1111"}

	SetFitnessFunc(func(gene Genome) int {
		return strings.Count(gene.Sequence, "1")
	})

	if genome.Fitness() != 4 {
		t.Error("String is not correct fitness")
	}

	genome = Genome{"0011"}

	SetFitnessFunc(func(gene Genome) int {
		return strings.Count(gene.Sequence, "0")
	})

	if genome.Fitness() != 2 {
		t.Error("String is not correct fitness")
	}
}

func TestAverageFitness(t *testing.T) {
	SetFitnessFunc(func(gene Genome) int {
		return strings.Count(gene.Sequence, "1")
	})

	population := []Genome{
		{"1111"},
		{"1111"},
		{"0000"},
		{"0000"},
	}

	if AverageFitness(population) != 2 {
		t.Error("Incorrect average fitness")
	}
}

