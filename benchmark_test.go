package ga

import (
	"math/rand"
	"testing"
)

var seed int64 = 3
var (
	globalBestCandidate Genome
	globalNumIterations int
	globalPopulation    []Genome
)

func benchmarkGATournament(length, generations int, terminateEarly bool, b *testing.B) {
	rand.Seed(seed)
	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(TournamentSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)
	var (
		bestCandidate Genome
		numIterations int
		population    []Genome
	)
	for n := 0; n < b.N; n++ {
		bestCandidate, numIterations, population = GeneticAlgorithm(length, length, generations, true, true, terminateEarly)
		b.Log("Best Candidate", bestCandidate)
		b.Log("Num Iterations:", numIterations)
		b.Log("Population:", population)
	}
	globalBestCandidate = bestCandidate
	globalNumIterations = numIterations
	globalPopulation = population
}

func BenchmarkGATournamentFull_10(b *testing.B)           { benchmarkGATournament(10, 100, false, b) }
func BenchmarkGATournamentFull_20(b *testing.B)           { benchmarkGATournament(20, 200, false, b) }
func BenchmarkGATournamentFull_50(b *testing.B)           { benchmarkGATournament(50, 500, false, b) }
func BenchmarkGATournamentTerminateEarly_10(b *testing.B) { benchmarkGATournament(10, 100, true, b) }
func BenchmarkGATournamentTerminateEarly_20(b *testing.B) { benchmarkGATournament(20, 200, true, b) }
func BenchmarkGATournamentTerminateEarly_50(b *testing.B) { benchmarkGATournament(50, 500, true, b) }

func benchmarkGARoulette(length, generations int, terminateEarly bool, b *testing.B) {
	rand.Seed(seed)
	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(RouletteSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)
	var (
		bestCandidate Genome
		numIterations int
		population    []Genome
	)
	for n := 0; n < b.N; n++ {
		bestCandidate, numIterations, population = GeneticAlgorithm(length, length, generations, true, true, terminateEarly)
		b.Log("Best Candidate", bestCandidate)
		b.Log("Num Iterations:", numIterations)
		b.Log("Population:", population)
	}
	globalBestCandidate = bestCandidate
	globalNumIterations = numIterations
	globalPopulation = population
}

func BenchmarkGARouletteFull_10(b *testing.B)           { benchmarkGARoulette(10, 100, false, b) }
func BenchmarkGARouletteFull_20(b *testing.B)           { benchmarkGARoulette(20, 200, false, b) }
func BenchmarkGARouletteFull_50(b *testing.B)           { benchmarkGARoulette(50, 500, false, b) }
func BenchmarkGARouletteTerminateEarly_10(b *testing.B) { benchmarkGARoulette(10, 100, true, b) }
func BenchmarkGARouletteTerminateEarly_20(b *testing.B) { benchmarkGARoulette(20, 200, true, b) }
func BenchmarkGARouletteTerminateEarly_50(b *testing.B) { benchmarkGARoulette(50, 500, true, b) }
