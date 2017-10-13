package ga

// AverageFitness returns the average fitness of a [] Genome population
func AverageFitness(population []Genome) int {
	var average int = 0
	for _, i := range population {
		average += i.Fitness()
	}
	return average / int(len(population))
}

// MaxFitness returns the highest fitness found in a [] Genome population
func MaxFitnessCandidate(population []Genome) Genome {
	var (
		max     int = 0
		maxGene Genome
	)
	for _, i := range population {
		if i.Fitness() > max {
			max = i.Fitness()
			maxGene = i
		}
	}
	return maxGene
}

// MaxFitness returns the highest fitness found in a [] Genome population
func MaxFitness(population []Genome) int {
	return MaxFitnessCandidate(population).Fitness()
}

var DefaultFitnessFunc func(gene Genome) int = func(gene Genome) int {
	count := 0
	for _, i := range gene.Sequence {
		if i == 1 {
			count++
		}
	}
	return count
}
var fitnessFunc = DefaultFitnessFunc

// SetFitnessFunc changes the fitness function to the function specified
func SetFitnessFunc(f func(gene Genome) int) {
	fitnessFunc = f
}

// Fitness calculates the suitability of a candidate solution and returns an integral score value
func (gene Genome) Fitness() int {
	return fitnessFunc(gene)
}
