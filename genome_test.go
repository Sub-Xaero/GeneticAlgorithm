package ga

import (
	"testing"
)

func TestGenome_ToString(t *testing.T) {
	t.Parallel()
	var genA = NewGeneticAlgorithm()
	genA.SetSeed(3)
	genA.SetOutputFunc(func(a ...interface{}) { t.Log(a...) })

	genA.SetFitnessFunc(func(gene Genome) int {
		count := 0
		for _, i := range gene.Sequence {
			if i == "1" {
				count++
			}
		}
		return count
	})

	outputString := Genome{Bitstring{"1", "1", "1", "1"}}.String()
	expected := "{[1 1 1 1 ]}"
	if outputString != expected {
		t.Error("Incorrect string:", outputString, "Expected:", expected)
	}

	outputString = Genome{Bitstring{"1", "0", "1", "0", "1", "0", "1", "0", "1", "0"}}.String()
	expected = "{[1 0 1 0 1 0 1 0 1 0 ]}"
	if outputString != expected {
		t.Error("Incorrect string:", outputString, "Expected:", expected)
	}

	outputString = Genome{Bitstring{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1"}}.String()
	expected = "{[1 1 1 1 1 1 1 1 1 1 ]}"
	if outputString != expected {
		t.Error("Incorrect string:", outputString, "Expected:", expected)
	}

	outputString = Genome{Bitstring{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1"}}.String()
	expected = "{[1 1 1 1 1 1 1 1 1 1 1 1 ]}"
	if outputString != expected {
		t.Error("Incorrect string:", outputString, "Expected:", expected)
	}
}
