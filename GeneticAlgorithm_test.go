package ga

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGeneticAlgorithmTerminateEarly10(t *testing.T) {
	rand.Seed(3)

	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(TournamentSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)

	globalVar := 10
	multiplier := 10
	bestCandidate, numIterations, finalPopulation := GeneticAlgorithm(globalVar, globalVar, globalVar*multiplier, true, true, true)
	fmt.Println(bestCandidate, finalPopulation)

	expectedFitness := 8
	gotFitness := bestCandidate.Fitness()
	if gotFitness < expectedFitness {
		t.Error("GA did not produce a suitable candidate.", "Expected at least:", expectedFitness, "Got:", gotFitness)
	} else {
		t.Log("GA produced a suitable candidate.", "Expected at least:", expectedFitness, "Got:", gotFitness)
	}

	if numIterations < globalVar*multiplier {
		t.Log("GA successfully detected stagnation. Terminated early. Expected less than", globalVar*multiplier, "iterations, GA took", numIterations)
	} else {
		t.Error("GA did not terminate early. Expected less than", globalVar*multiplier, "iterations, GA took", numIterations)
	}
}

func TestGeneticAlgorithmTerminateEarly20(t *testing.T) {
	rand.Seed(3)

	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(TournamentSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)

	globalVar := 20
	multiplier := 10
	bestCandidate, numIterations, finalPopulation := GeneticAlgorithm(globalVar, globalVar, globalVar*multiplier, true, true, true)
	fmt.Println(bestCandidate, finalPopulation)

	expectedFitness := 18
	gotFitness := bestCandidate.Fitness()
	if gotFitness < expectedFitness {
		t.Error("GA did not produce a suitable candidate.", "Expected at least:", expectedFitness, "Got:", gotFitness)
	} else {
		t.Log("GA produced a suitable candidate.", "Expected at least:", expectedFitness, "Got:", gotFitness)
	}

	if numIterations < globalVar*multiplier {
		t.Log("GA successfully detected stagnation. Terminated early. Expected less than", globalVar*multiplier, "iterations, GA took", numIterations)
	} else {
		t.Error("GA did not terminate early. Expected less than", globalVar*multiplier, "iterations, GA took", numIterations)
	}
}

func TestGeneticAlgorithmTerminateEarly30(t *testing.T) {
	rand.Seed(3)

	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(TournamentSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)

	globalVar := 30
	multiplier := 10
	bestCandidate, numIterations, finalPopulation := GeneticAlgorithm(globalVar, globalVar, globalVar*10, true, true, true)
	fmt.Println(bestCandidate, finalPopulation)

	expectedFitness := 28
	gotFitness := bestCandidate.Fitness()
	if gotFitness < expectedFitness {
		t.Error("GA did not produce a suitable candidate.", "Expected at least:", expectedFitness, "Got:", gotFitness)
	} else {
		t.Log("GA produced a suitable candidate.", "Expected at least:", expectedFitness, "Got:", gotFitness)
	}

	if numIterations < globalVar*multiplier {
		t.Log("GA successfully detected stagnation. Terminated early. Expected less than", globalVar*multiplier, "iterations, GA took", numIterations)
	} else {
		t.Error("GA did not terminate early. Expected less than", globalVar*multiplier, "iterations, GA took", numIterations)
	}
}

func TestGeneticAlgorithmTerminateEarly40(t *testing.T) {
	rand.Seed(3)

	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(TournamentSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)

	globalVar := 40
	multiplier := 10
	bestCandidate, numIterations, finalPopulation := GeneticAlgorithm(globalVar, globalVar, globalVar*multiplier, true, true, true)
	fmt.Println(bestCandidate, finalPopulation)

	expectedFitness := 38
	gotFitness := bestCandidate.Fitness()
	if gotFitness < expectedFitness {
		t.Error("GA did not produce a suitable candidate.", "Expected at least:", expectedFitness, "Got:", gotFitness)
	} else {
		t.Log("GA produced a suitable candidate.", "Expected at least:", expectedFitness, "Got:", gotFitness)
	}

	if numIterations < globalVar*multiplier {
		t.Log("GA successfully detected stagnation. Terminated early. Expected less than", globalVar*multiplier, "iterations, GA took", numIterations)
	} else {
		t.Error("GA did not terminate early. Expected less than", globalVar*multiplier, "iterations, GA took", numIterations)
	}
}

func TestGeneticAlgorithmTerminateEarly50(t *testing.T) {
	rand.Seed(3)

	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(TournamentSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)

	globalVar := 50
	multiplier := 10
	bestCandidate, numIterations, finalPopulation := GeneticAlgorithm(globalVar, globalVar, globalVar*multiplier, true, true, true)
	fmt.Println(bestCandidate, finalPopulation)

	expectedFitness := 48
	gotFitness := bestCandidate.Fitness()
	if gotFitness < expectedFitness {
		t.Error("GA did not produce a suitable candidate.", "Expected at least:", expectedFitness, "Got:", gotFitness)
	} else {
		t.Log("GA produced a suitable candidate.", "Expected at least:", expectedFitness, "Got:", gotFitness)
	}

	if numIterations < globalVar*multiplier {
		t.Log("GA successfully detected stagnation. Terminated early. Expected less than", globalVar*multiplier, "iterations, GA took", numIterations)
	} else {
		t.Error("GA did not terminate early. Expected less than", globalVar*multiplier, "iterations, GA took", numIterations)
	}
}

func TestGeneticAlgorithmFull10(t *testing.T) {
	rand.Seed(3)

	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(TournamentSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)

	globalVar := 10
	bestCandidate, _, finalPopulation := GeneticAlgorithm(globalVar, globalVar, globalVar, true, true, false)
	fmt.Println(bestCandidate, finalPopulation)

	expectedFitness := 8
	gotFitness := bestCandidate.Fitness()
	if gotFitness < expectedFitness {
		t.Error("GA did not produce a suitable candidate.", "Expected at least:", expectedFitness, "Got:", gotFitness)
	} else {
		t.Log("GA produced a suitable candidate.", "Expected at least:", expectedFitness, "Got:", gotFitness)
	}
}

func TestGeneticAlgorithmFull20(t *testing.T) {
	rand.Seed(3)

	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(TournamentSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)

	globalVar := 20
	bestCandidate, _, finalPopulation := GeneticAlgorithm(globalVar, globalVar, globalVar, true, true, false)
	fmt.Println(bestCandidate, finalPopulation)

	expectedFitness := 18
	gotFitness := bestCandidate.Fitness()
	if gotFitness < expectedFitness {
		t.Error("GA did not produce a suitable candidate.", "Expected at least:", expectedFitness, "Got:", gotFitness)
	} else {
		t.Log("GA produced a suitable candidate.", "Expected at least:", expectedFitness, "Got:", gotFitness)
	}
}

func TestGeneticAlgorithmFull30(t *testing.T) {
	rand.Seed(3)

	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(TournamentSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)

	globalVar := 30
	bestCandidate, _, finalPopulation := GeneticAlgorithm(globalVar, globalVar, globalVar, true, true, false)
	fmt.Println(bestCandidate, finalPopulation)

	expectedFitness := 28
	gotFitness := bestCandidate.Fitness()
	if gotFitness < expectedFitness {
		t.Error("GA did not produce a suitable candidate.", "Expected at least:", expectedFitness, "Got:", gotFitness)
	} else {
		t.Log("GA produced a suitable candidate.", "Expected at least:", expectedFitness, "Got:", gotFitness)
	}
}

func TestGeneticAlgorithmFull40(t *testing.T) {
	rand.Seed(3)

	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(TournamentSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)

	globalVar := 40
	bestCandidate, _, finalPopulation := GeneticAlgorithm(globalVar, globalVar, globalVar, true, true, false)
	fmt.Println(bestCandidate, finalPopulation)

	expectedFitness := 38
	gotFitness := bestCandidate.Fitness()
	if gotFitness < expectedFitness {
		t.Error("GA did not produce a suitable candidate.", "Expected at least:", expectedFitness, "Got:", gotFitness)
	} else {
		t.Log("GA produced a suitable candidate.", "Expected at least:", expectedFitness, "Got:", gotFitness)
	}
}

func TestGeneticAlgorithmFull50(t *testing.T) {
	rand.Seed(3)

	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(TournamentSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)

	globalVar := 50
	bestCandidate, _, finalPopulation := GeneticAlgorithm(globalVar, globalVar, globalVar, true, true, false)
	fmt.Println(bestCandidate, finalPopulation)

	expectedFitness := 48
	gotFitness := bestCandidate.Fitness()
	if gotFitness < expectedFitness {
		t.Error("GA did not produce a suitable candidate.", "Expected at least:", expectedFitness, "Got:", gotFitness)
	} else {
		t.Log("GA produced a suitable candidate.", "Expected at least:", expectedFitness, "Got:", gotFitness)
	}
}

func TestFillRandomPopulation(t *testing.T) {
	rand.Seed(time.Now().Unix())
	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(TournamentSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)

	expectedLen := 10
	population := make([]Genome, 0)
	population = FillRandomPopulation(population, expectedLen, expectedLen)
	gotLen := len(population)

	if gotLen != expectedLen {
		t.Error("Population not filled to specified size. Expected:", expectedLen, "Got:", gotLen)
	} else {
		t.Log("Population successfuly filled to specified size. Expected:", expectedLen, "Got:", gotLen)
	}

	for i, val := range population {
		if len(val.Sequence) == 0 {
			t.Error("Bitstring:", i, " in population is empty")
		} else {
			t.Log("Bitstring:", i, " in population successfully filled")
		}
	}
}
