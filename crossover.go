package ga

import (
	"errors"
	"math/rand"
)

type CrossoverFunction func(Genome, Genome, *rand.Rand) ([]Genome, error)

var DefaultCrossoverFunc CrossoverFunction = func(gene, spouse Genome, random *rand.Rand) ([]Genome, error) {
	gene = gene.Copy()
	spouse = spouse.Copy()
	if len(gene.Sequence) != len(spouse.Sequence) {
		return nil, errors.New("strings are not same length")
	}
	crossover := random.Int() % len(gene.Sequence)
	return []Genome{
		{append(append(make(bitstring, 0), gene.Sequence[:crossover]...), spouse.Sequence[crossover:]...)},
		{append(append(make(bitstring, 0), spouse.Sequence[:crossover]...), gene.Sequence[crossover:]...)},
	}, nil
}

// SetCrossoverFunc changes the crossover function to the function specified
func (genA *GeneticAlgorithm) SetCrossoverFunc(f CrossoverFunction) {
	genA.Crossover = f
}
