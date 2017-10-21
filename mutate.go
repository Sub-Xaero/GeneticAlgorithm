package ga

import (
	"math/rand"
)

type MutateFunction func(Genome, *rand.Rand) Genome

var DefaultMutateFunc MutateFunction = func(gene Genome, random *rand.Rand) Genome {
	gene = gene.Copy()
	choice := random.Int() % len(gene.Sequence)
	if gene.Sequence[choice] == "1" {
		gene.Sequence[choice] = "0"
	} else {
		gene.Sequence[choice] = "1"
	}
	return gene
}

// SetMutateFunc changes the mutate function to the function specified
func (genA *GeneticAlgorithm) SetMutateFunc(f MutateFunction) {
	genA.Mutate = f
}
