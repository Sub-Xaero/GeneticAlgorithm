package ga

import (
	"math/rand"
)

var DefaultMutateFunc func(gene Genome) Genome = func(gene Genome) Genome {
	choice := rand.Int() % len(gene.Sequence)
	if gene.Sequence[choice] == 1 {
		gene.Sequence[choice] = 0
	} else {
		gene.Sequence[choice] = 1
	}
	return gene
}
var mutateFunc = DefaultMutateFunc

// SetMutateFunc changes the mutate function to the function specified
func SetMutateFunc(f func(gene Genome) Genome) {
	mutateFunc = f
}

// Mutate returns a bitstring with bits mutated by a function set by SetMutateFunc
func (gene Genome) Mutate() Genome {
	return mutateFunc(gene)
}
