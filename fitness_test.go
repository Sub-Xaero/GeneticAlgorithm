package ga

import (
	"testing"
)

func TestGenome_DefaultFitness(t *testing.T) {
	genome := Genome{[]int{1, 1, 1, 1}}

	if genome.Fitness() != 4 {
		t.Error("String is not correct fitness")
	}

	genome = Genome{[]int{0, 0, 1, 1}}

	SetFitnessFunc(func(gene Genome) int {
		count := 0
		for _, i := range gene.Sequence {
			if i == 0 {
				count++
			}
		}
		return count
	})

	if genome.Fitness() != 2 {
		t.Error("String is not correct fitness")
	}
}

func TestGenome_CustomFitness(t *testing.T) {
	genome := Genome{[]int{1, 1, 1, 1}}

	SetFitnessFunc(func(gene Genome) int {
		count := 0
		for _, i := range gene.Sequence {
			if i == 1 {
				count++
			}
		}
		return count
	})

	if genome.Fitness() != 4 {
		t.Error("String is not correct fitness")
	}

	genome = Genome{[]int{0, 0, 1, 1}}

	SetFitnessFunc(func(gene Genome) int {
		count := 0
		for _, i := range gene.Sequence {
			if i == 0 {
				count++
			}
		}
		return count
	})

	if genome.Fitness() != 2 {
		t.Error("String is not correct fitness")
	}
}

func TestAverageFitness(t *testing.T) {
	SetFitnessFunc(func(gene Genome) int {
		count := 0
		for _, i := range gene.Sequence {
			if i == 1 {
				count++
			}
		}
		return count
	})
	population := []Genome{
		{[]int{1, 1, 1, 1}},
		{[]int{1, 1, 1, 1}},
		{[]int{0, 0, 0, 0}},
		{[]int{0, 0, 0, 0}},
	}

	if AverageFitness(population) != 2 {
		t.Error("Incorrect average fitness")
	}
}

func TestMaxFitness(t *testing.T) {
	SetFitnessFunc(func(gene Genome) int {
		count := 0
		for _, i := range gene.Sequence {
			if i == 1 {
				count++
			}
		}
		return count
	})

	population := []Genome{
		{[]int{1, 1, 1, 1, 1, 1, 1, 1}},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1}},
		{[]int{0, 0, 0, 0, 0, 0, 0, 0}},
		{[]int{0, 0, 0, 0, 0, 0, 0, 0}},
	}

	if MaxFitness(population) != 8 {
		t.Error("Incorrect max fitness")
	}

	SetFitnessFunc(func(gene Genome) int {
		count := 0
		for _, i := range gene.Sequence {
			if i == 0 {
				count++
			}
		}
		return count
	})

	population = []Genome{
		{[]int{1, 1, 1, 1, 1, 1, 1, 1}},
		{[]int{0, 0, 0, 0, 1, 1, 1, 1}},
		{[]int{0, 0, 0, 0, 0, 0, 0, 0}},
		{[]int{0, 0, 0, 0, 0, 0, 0, 0}},
	}

	if MaxFitness(population) != 8 {
		t.Error("Incorrect max fitness")
	}
}
