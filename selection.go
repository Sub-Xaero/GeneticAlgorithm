package ga

import (
	"math/rand"
)

type SelectFunction func(FitnessFunction, Population, *rand.Rand) Population

var TournamentSelection SelectFunction = func(Fitness FitnessFunction, candidatePool Population, random *rand.Rand) Population {
	offspring := make(Population, 0)

	for i := 0; i < len(candidatePool); i++ {
		parent1 := candidatePool[random.Int()%len(candidatePool)]
		parent2 := candidatePool[random.Int()%len(candidatePool)]

		if Fitness(parent1) > Fitness(parent2) {
			offspring = append(offspring, parent1)
		} else {
			offspring = append(offspring, parent2)
		}
	}

	return offspring
}

var RouletteSelection SelectFunction = func(Fitness FitnessFunction, candidatePool Population, random *rand.Rand) Population {
	offspring := make(Population, 0)
	for range candidatePool {
		weightSum := 0
		for _, val := range candidatePool {
			weightSum += Fitness(val)
		}
		choice := random.Float32() * float32(weightSum)
		for _, val := range candidatePool {
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
