package ga

import (
	"math/rand"
)

type SelectFunction func(FitnessFunction, []Genome, *rand.Rand) []Genome

var TournamentSelection SelectFunction = func(Fitness FitnessFunction, population []Genome, random *rand.Rand) []Genome {
	offspring := make([]Genome, 0)

	for i := 0; i < len(population); i++ {
		parent1 := population[random.Int()%len(population)]
		parent2 := population[random.Int()%len(population)]

		if Fitness(parent1) > Fitness(parent2) {
			offspring = append(offspring, parent1)
		} else {
			offspring = append(offspring, parent2)
		}
	}

	return offspring
}

var RouletteSelection SelectFunction = func(Fitness FitnessFunction, population []Genome, random *rand.Rand) []Genome {
	offspring := make([]Genome, 0)
	for range population {
		weightSum := 0
		for _, val := range population {
			weightSum += Fitness(val)
		}
		choice := random.Float32() * float32(weightSum)
		for _, val := range population {
			choice -= float32(Fitness(val))
			if choice <= 0 {
				offspring = append(offspring, val.Copy())
				break
			}
		}
	}
	return offspring
}

// SetSelectionFunc changes the selection function to the function specified
func (genA *GeneticAlgorithm) SetSelectionFunc(f SelectFunction) {
	genA.Selection = f
}
