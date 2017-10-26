package ga

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSetMutateFunc(t *testing.T) {
	t.Parallel()
	var genA = NewGeneticAlgorithm()
	genA.SetSeed(3)
	genA.SetOutputFunc(func(a ...interface{}) { t.Log(a...) })

	genA.SetMutateFunc(func(gene Genome, random *rand.Rand) Genome {
		return Genome{Bitstring{"1", "2", "3", "4"}}
	})

	output := fmt.Sprint(genA.Mutate(Genome{}, genA.RandomEngine))
	expectedOutput := "{[1234]}"
	if output != expectedOutput {
		t.Error("Mutate function not set. Expected:", expectedOutput, "Got:", output)
	} else {
		t.Log("Mutate function was set successfully. Expected:", expectedOutput, "Got:", output)
	}
}

func TestDefaultMutateFunc(t *testing.T) {
	t.Parallel()
	genA := NewGeneticAlgorithm()
	gene := Genome{Bitstring{"1", "0", "1", "0", "1"}}
	geneOutput := genA.Mutate(gene, genA.RandomEngine)

	if gene.String() == geneOutput.String() {
		t.Error("Mutate did not change bitstrings. At least one mutation should occur. Was:", gene, "Mutated:", geneOutput)
	} else {
		t.Log("Mutate successfully changed bitstrings. At least one mutation should occur. Was:", gene, "Mutated:", geneOutput)
	}
}
