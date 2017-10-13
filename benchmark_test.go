package ga

import (
	"testing"
)

func BenchmarkGeneticAlgorithm_10Length(b *testing.B) {
	length := 10
	SetFitnessFunc(func(gene Genome) int {
		count := 0
		for _, i := range gene.Sequence {
			if i == 1 {
				count++
			}
		}
		return count
	})
	GeneticAlgorithm(length, length, b.N, true, true, false)
}

func BenchmarkGeneticAlgorithm_20Length(b *testing.B) {
	length := 20
	SetFitnessFunc(func(gene Genome) int {
		count := 0
		for _, i := range gene.Sequence {
			if i == 1 {
				count++
			}
		}
		return count
	})
	GeneticAlgorithm(length, length, b.N, true, true, false)
}

func BenchmarkGeneticAlgorithm_50Length(b *testing.B) {
	length := 50
	SetFitnessFunc(func(gene Genome) int {
		count := 0
		for _, i := range gene.Sequence {
			if i == 1 {
				count++
			}
		}
		return count
	})
	GeneticAlgorithm(length, length, b.N, true, true, false)
}
