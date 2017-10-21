package ga

import (
	"testing"
)

func TestGenome_DefaultFitness(t *testing.T) {
	t.Parallel()
	var genA = NewGeneticAlgorithm()
	genA.SetSeed(3)
	genA.SetOutputFunc(func(a ...interface{}) { t.Log(a...) })

	genome := Genome{bitstring{"1", "1", "1", "1"}}

	t.Log("Genome:", genome)
	t.Log("Setting fitness func to default...")
	genA.SetFitnessFunc(DefaultFitnessFunc)

	expectedFitness := 4
	gotFitness := genA.Fitness(genome)

	if gotFitness != expectedFitness {
		t.Error("String is not correct fitness.", "Expected:", expectedFitness, "Got:", gotFitness)
	} else {
		t.Log("String is correct fitness.", "Expected:", expectedFitness, "Got:", gotFitness)
	}
}

func TestGenome_CustomFitness(t *testing.T) {
	t.Parallel()
	var genA = NewGeneticAlgorithm()
	genA.SetSeed(3)
	genA.SetOutputFunc(func(a ...interface{}) { t.Log(a...) })

	genome := Genome{bitstring{"0", "0", "0", "1"}}
	t.Log(genome)
	t.Log("Setting fitness func to custom...")
	genA.SetFitnessFunc(func(gene Genome) int {
		count := 0
		for _, i := range gene.Sequence {
			if i == "0" {
				count++
			}
		}
		return count
	})

	expectedFitness := 3
	gotFitness := genA.Fitness(genome)

	if gotFitness != expectedFitness {
		t.Error("String is not correct fitness.", "Expected:", expectedFitness, "Got:", gotFitness)
	} else {
		t.Log("String is correct fitness.", "Expected:", expectedFitness, "Got:", gotFitness)
	}
}

func TestAverageFitness(t *testing.T) {
	t.Parallel()
	var genA = NewGeneticAlgorithm()
	genA.SetSeed(3)
	genA.SetOutputFunc(func(a ...interface{}) { t.Log(a...) })

	t.Log("Setting fitness func to default...")
	genA.SetFitnessFunc(DefaultFitnessFunc)

	population := []Genome{
		{bitstring{"1", "1", "1", "1"}},
		{bitstring{"1", "1", "1", "1"}},
		{bitstring{"0", "0", "0", "0"}},
		{bitstring{"0", "0", "0", "0"}},
	}
	t.Log("Created population:", population)

	expectedFitness := 2
	gotFitness := genA.AverageFitness(population)
	if gotFitness != expectedFitness {
		t.Error("Incorrect average fitness.", "Expected:", expectedFitness, "Got:", gotFitness)
	} else {
		t.Log("Correct average fitness.", "Expected:", expectedFitness, "Got:", gotFitness)
	}
}

func TestMaxFitness(t *testing.T) {
	t.Parallel()
	var genA = NewGeneticAlgorithm()
	genA.SetSeed(3)
	genA.SetOutputFunc(func(a ...interface{}) { t.Log(a...) })

	t.Log("Setting fitness func to default...")
	genA.SetFitnessFunc(DefaultFitnessFunc)

	population := []Genome{
		{bitstring{"1", "1", "1", "1", "1", "1", "1", "1"}},
		{bitstring{"1", "1", "1", "1", "1", "1", "1", "1"}},
		{bitstring{"0", "0", "0", "0", "0", "0", "0", "0"}},
		{bitstring{"0", "0", "0", "0", "0", "0", "0", "0"}},
	}
	t.Log("Created population:", population)

	expectedFitness := 8
	gotFitness := genA.MaxFitness(population)
	if gotFitness != expectedFitness {
		t.Error("Incorrect max fitness.", "Expected:", expectedFitness, "Got:", gotFitness)
	} else {
		t.Log("Correct max fitness.", "Expected:", expectedFitness, "Got:", gotFitness)
	}

	t.Log("Setting fitness func to custom...")
	genA.SetFitnessFunc(func(gene Genome) int {
		count := 0
		for _, i := range gene.Sequence {
			if i == "0" {
				count++
			}
		}
		return count
	})

	population = []Genome{
		{bitstring{"1", "1", "1", "1", "1", "1", "1", "1"}},
		{bitstring{"0", "0", "0", "0", "1", "1", "1", "1"}},
		{bitstring{"0", "0", "0", "0", "0", "0", "0", "0"}},
		{bitstring{"0", "0", "0", "0", "0", "0", "0", "0"}},
	}
	t.Log("Created population:", population)

	expectedFitness = 8
	gotFitness = genA.MaxFitness(population)
	if gotFitness != expectedFitness {
		t.Error("Incorrect max fitness.", "Expected:", expectedFitness, "Got:", gotFitness)
	} else {
		t.Log("Correct max fitness.", "Expected:", expectedFitness, "Got:", gotFitness)
	}
}
