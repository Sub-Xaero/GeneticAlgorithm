package ga

import (
	"testing"
)

func TestGenome_DefaultFitness(t *testing.T) {
	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(TournamentSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)

	genome := Genome{[]int{1, 1, 1, 1}}

	t.Log("Genome:", genome)
	t.Log("Setting fitness func to default...")
	SetFitnessFunc(DefaultFitnessFunc)

	expectedFitness := 4
	gotFitness := genome.Fitness()

	if gotFitness != expectedFitness {
		t.Error("String is not correct fitness.", "Expected:", expectedFitness, "Got:", gotFitness)
	} else {
		t.Log("String is correct fitness.", "Expected:", expectedFitness, "Got:", gotFitness)
	}
}

func TestGenome_CustomFitness(t *testing.T) {
	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(TournamentSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)

	genome := Genome{[]int{0, 0, 0, 1}}
	t.Log(genome)
	t.Log("Setting fitness func to custom...")
	SetFitnessFunc(func(gene Genome) int {
		count := 0
		for _, i := range gene.Sequence {
			if i == 0 {
				count++
			}
		}
		return count
	})

	expectedFitness := 3
	gotFitness := genome.Fitness()

	if gotFitness != expectedFitness {
		t.Error("String is not correct fitness.", "Expected:", expectedFitness, "Got:", gotFitness)
	} else {
		t.Log("String is correct fitness.", "Expected:", expectedFitness, "Got:", gotFitness)
	}
}

func TestAverageFitness(t *testing.T) {
	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(TournamentSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)

	t.Log("Setting fitness func to default...")
	SetFitnessFunc(DefaultFitnessFunc)

	population := []Genome{
		{[]int{1, 1, 1, 1}},
		{[]int{1, 1, 1, 1}},
		{[]int{0, 0, 0, 0}},
		{[]int{0, 0, 0, 0}},
	}
	t.Log("Created population:", population)

	expectedFitness := 2
	gotFitness := AverageFitness(population)
	if gotFitness != expectedFitness {
		t.Error("Incorrect average fitness.", "Expected:", expectedFitness, "Got:", gotFitness)
	} else {
		t.Log("Correct average fitness.", "Expected:", expectedFitness, "Got:", gotFitness)
	}
}

func TestMaxFitness(t *testing.T) {
	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(TournamentSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)

	t.Log("Setting fitness func to default...")
	SetFitnessFunc(DefaultFitnessFunc)

	population := []Genome{
		{[]int{1, 1, 1, 1, 1, 1, 1, 1}},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1}},
		{[]int{0, 0, 0, 0, 0, 0, 0, 0}},
		{[]int{0, 0, 0, 0, 0, 0, 0, 0}},
	}
	t.Log("Created population:", population)

	expectedFitness := 8
	gotFitness := MaxFitness(population)
	if gotFitness != expectedFitness {
		t.Error("Incorrect max fitness.", "Expected:", expectedFitness, "Got:", gotFitness)
	} else {
		t.Log("Correct max fitness.", "Expected:", expectedFitness, "Got:", gotFitness)
	}

	t.Log("Setting fitness func to custom...")
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
	t.Log("Created population:", population)

	expectedFitness = 8
	gotFitness = MaxFitness(population)
	if gotFitness != expectedFitness {
		t.Error("Incorrect max fitness.", "Expected:", expectedFitness, "Got:", gotFitness)
	} else {
		t.Log("Correct max fitness.", "Expected:", expectedFitness, "Got:", gotFitness)
	}
}
