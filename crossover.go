package ga

import (
	"errors"
	"math/rand"
)

var DefaultCrossoverFunc = func(gene, spouse Genome) ([]Genome, error) {
	if len(gene.Sequence) != len(spouse.Sequence) {
		return nil, errors.New("strings are not current length")
	}
	crossover := rand.Int() % len(gene.Sequence)
	return []Genome{
		{append(append(make([]int, 0), gene.Sequence[:crossover]...), spouse.Sequence[crossover:]...)},
		{append(append(make([]int, 0), spouse.Sequence[:crossover]...), gene.Sequence[crossover:]...)},
	}, nil
}
var crossoverFunc = DefaultCrossoverFunc

// SetCrossoverFunc changes the crossover function to the function specified
func SetCrossoverFunc(f func(Genome, Genome) ([]Genome, error)) {
	crossoverFunc = f
}

// Crossover applies a function, set by SetCrossoverFunc to the receiver gene and a specified pair
func (gene Genome) Crossover(spouse Genome) ([]Genome, error) {
	return crossoverFunc(gene, spouse)
}
