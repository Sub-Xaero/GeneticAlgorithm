package ga

import (
	"math/rand"
)

var TournamentSelection func([]Genome) []Genome = func(population []Genome) []Genome {
	offspring := make([]Genome, 0)

	for i := 0; i < len(population); i++ {
		parent1 := population[rand.Int()%len(population)]
		parent2 := population[rand.Int()%len(population)]

		if parent1.Fitness() > parent2.Fitness() {
			offspring = append(offspring, parent1)
		} else {
			offspring = append(offspring, parent2)
		}
	}

	return offspring
}

var RouletteSelection func([]Genome) []Genome = func(population []Genome) []Genome {
	offspring := make([]Genome, 0)
	for range population {
		weightSum := 0
		for _, val := range population {
			weightSum += val.Fitness()
		}
		choice := rand.Float32() * float32(weightSum)
		for _, val := range population {
			choice -= float32(val.Fitness())
			if choice <= 0 {
				offspring = append(offspring, val.Copy())
				break
			}
		}
	}
	return offspring
}

var selectionFunc = TournamentSelection

// SetSelectionFunc changes the selection function to the function specified
func SetSelectionFunc(f func([]Genome) []Genome) {
	selectionFunc = f
}

// Selection processes a population according to the function defined in SetSelectionFunc and returns an array of offspring
func Selection(population []Genome) []Genome {
	return selectionFunc(population)
}
