package ga

type FitnessFunction func(gene Genome) int

var DefaultFitnessFunc FitnessFunction = func(gene Genome) int {
	count := 0
	for _, i := range gene.Sequence {
		if i == "1" {
			count++
		}
	}
	return count
}

// SetFitnessFunc changes the fitness function to the function specified
func (genA *GeneticAlgorithm) SetFitnessFunc(f FitnessFunction) {
	genA.Fitness = f
}

// AverageFitness returns the average fitness of a [] Genome candidatePool
func (genA *GeneticAlgorithm) AverageFitness(candidatePool Population) int {
	var average int = 0
	for _, i := range candidatePool {
		average += genA.Fitness(i)
	}
	return average / int(len(candidatePool))
}

// MaxFitness returns the highest fitness found in a [] Genome candidatePool
func (genA *GeneticAlgorithm) MaxFitnessCandidate(candidatePool Population) Genome {
	var (
		max     int = 0
		maxGene Genome
	)
	for _, i := range candidatePool {
		if genA.Fitness(i) > max {
			max = genA.Fitness(i)
			maxGene = i
		}
	}
	return maxGene
}

// MaxFitness returns the highest fitness found in a [] Genome candidatePool
func (genA *GeneticAlgorithm) MaxFitness(candidatePool Population) int {
	return genA.Fitness(genA.MaxFitnessCandidate(candidatePool))
}
