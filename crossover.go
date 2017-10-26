package ga

import (
	"errors"
	"math/rand"
)

type CrossoverFunction func(Genome, Genome, *rand.Rand) (Population, error)

var DefaultCrossoverFunc CrossoverFunction = func(gene, spouse Genome, random *rand.Rand) (Population, error) {
	gene = gene.Copy()
	spouse = spouse.Copy()
	if len(gene.Sequence) != len(spouse.Sequence) {
		return nil, errors.New("strings are not same length")
	}
	crossover := random.Int() % len(gene.Sequence)
	return Population{
		{append(append(make(Bitstring, 0), gene.Sequence[:crossover]...), spouse.Sequence[crossover:]...)},
		{append(append(make(Bitstring, 0), spouse.Sequence[:crossover]...), gene.Sequence[crossover:]...)},
	}, nil
}

// SetCrossoverFunc changes the crossover function to the function specified
func (genA *GeneticAlgorithm) SetCrossoverFunc(f CrossoverFunction) {
	genA.Crossover = f
}
