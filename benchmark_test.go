package ga

import (
	"strings"
	"testing"
)

func BenchmarkGeneticAlgorithm_10Length(b *testing.B) {
	length := 10
	SetFitnessFunc(func(gene Genome) int {
		return strings.Count(gene.Sequence, "1")
	})
	GeneticAlgorithm(length, length, b.N, length)
}

func BenchmarkGeneticAlgorithm_20Length(b *testing.B) {
	length := 20
	SetFitnessFunc(func(gene Genome) int {
		return strings.Count(gene.Sequence, "1")
	})
	GeneticAlgorithm(length, length, b.N, length)
}

func BenchmarkGeneticAlgorithm_50Length(b *testing.B) {
	length := 50
	SetFitnessFunc(func(gene Genome) int {
		return strings.Count(gene.Sequence, "1")
	})
	GeneticAlgorithm(length, length, b.N, length)
}