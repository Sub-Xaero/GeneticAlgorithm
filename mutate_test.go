package ga

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSetMutateFunc(t *testing.T) {
	t.Parallel()
	rand.Seed(3)
	var genA = NewGeneticAlgorithm()
	genA.SetOutputFunc(func(a ...interface{}) { t.Log(a...) })

	genA.SetMutateFunc(func(gene Genome) Genome {
		return Genome{bitstring{"1", "2", "3", "4"}}
	})

	output := fmt.Sprint(genA.Mutate(Genome{bitstring{}}))
	expectedOutput := "{[1 2 3 4]}"
	if output != expectedOutput {
		t.Error("Mutate function not set. Expected:", expectedOutput, "Got:", output)
	} else {
		t.Log("Mutate function was set successfully. Expected:", expectedOutput, "Got:", output)
	}
}
