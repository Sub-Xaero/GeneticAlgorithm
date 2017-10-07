package ga

import "strings"

// AverageFitness returns the average fitness of a [] Genome population
func AverageFitness(population []Genome) int {
	var average int = 0
	for _, i := range population {
		average += i.Fitness()
	}
	return average / int(len(population))
}

// MaxFitness returns the highest fitness found in a [] Genome population
func MaxFitness(population []Genome) int {
	var max int = 0
	for _, i := range population {
		if i.Fitness() > max {
			max = i.Fitness()
		}
	}
	return max
}

var fitnessFunc func(gene Genome) int = func(gene Genome) int {
	return strings.Count(gene.Sequence, "1")
}

// SetFitnessFunc changes the fitness function to the function specified
func SetFitnessFunc(f func(gene Genome) int) {
	fitnessFunc = f
}

// Fitness calculates the suitability of a candidate solution and returns an integral score value
func (gene Genome) Fitness() int {
	return fitnessFunc(gene)
}
