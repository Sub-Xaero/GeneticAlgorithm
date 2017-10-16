package ga

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const (
	ROULETTE   = iota
	TOURNAMENT = iota
)

func testGA(length, generations, expectedFitness, selectionMethod int, terminateEarly bool, t *testing.T) {
	rand.Seed(3)

	switch selectionMethod {
	case TOURNAMENT:
		SetSelectionFunc(TournamentSelection)
	case ROULETTE:
		SetSelectionFunc(RouletteSelection)
	}
	SetMutateFunc(DefaultMutateFunc)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)

	bestCandidate, numIterations, finalPopulation := GeneticAlgorithm(length, length, generations, true, true, terminateEarly)
	fmt.Println(bestCandidate, finalPopulation)

	gotFitness := bestCandidate.Fitness()
	if gotFitness < expectedFitness {
		t.Error("GA did not produce a suitable candidate.", "Expected at least:", expectedFitness, "Got:", gotFitness)
	} else {
		t.Log("GA produced a suitable candidate.", "Expected at least:", expectedFitness, "Got:", gotFitness)
	}

	if terminateEarly && numIterations < generations {
		t.Log("GA successfully detected stagnation. Terminated early. Expected less than", generations, "iterations, GA took", numIterations)
	} else if terminateEarly {
		t.Error("GA did not terminate early. Expected less than", generations, "iterations, GA took", numIterations)
	}
}

func TestGA_Tournament_TerminateEarly_10(t *testing.T) { testGA(10, 100, 8, TOURNAMENT, true, t) }
func TestGA_Tournament_TerminateEarly_20(t *testing.T) { testGA(20, 200, 15, TOURNAMENT, true, t) }
func TestGA_Tournament_TerminateEarly_30(t *testing.T) { testGA(30, 300, 25, TOURNAMENT, true, t) }
func TestGA_Tournament_TerminateEarly_40(t *testing.T) { testGA(40, 400, 35, TOURNAMENT, true, t) }
func TestGA_Tournament_TerminateEarly_50(t *testing.T) { testGA(50, 500, 45, TOURNAMENT, true, t) }
func TestGA_Tournament_Full_10(t *testing.T)           { testGA(10, 100, 8, TOURNAMENT, false, t) }
func TestGA_Tournament_Full_20(t *testing.T)           { testGA(20, 200, 15, TOURNAMENT, false, t) }
func TestGA_Tournament_Full_30(t *testing.T)           { testGA(30, 300, 25, TOURNAMENT, false, t) }
func TestGA_Tournament_Full_40(t *testing.T)           { testGA(40, 400, 35, TOURNAMENT, false, t) }
func TestGA_Tournament_Full_50(t *testing.T)           { testGA(50, 500, 45, TOURNAMENT, false, t) }

func TestGA_Roulette_TerminateEarly_10(t *testing.T) { testGA(10, 100, 8, ROULETTE, true, t) }
func TestGA_Roulette_TerminateEarly_20(t *testing.T) { testGA(20, 200, 15, ROULETTE, true, t) }
func TestGA_Roulette_TerminateEarly_30(t *testing.T) { testGA(30, 300, 25, ROULETTE, true, t) }
func TestGA_Roulette_TerminateEarly_40(t *testing.T) { testGA(40, 400, 35, ROULETTE, true, t) }
func TestGA_Roulette_TerminateEarly_50(t *testing.T) { testGA(50, 500, 45, ROULETTE, true, t) }
func TestGA_Roulette_Full_10(t *testing.T)           { testGA(10, 100, 8, ROULETTE, false, t) }
func TestGA_Roulette_Full_20(t *testing.T)           { testGA(20, 200, 15, ROULETTE, false, t) }
func TestGA_Roulette_Full_30(t *testing.T)           { testGA(30, 300, 25, ROULETTE, false, t) }
func TestGA_Roulette_Full_40(t *testing.T)           { testGA(40, 400, 35, ROULETTE, false, t) }
func TestGA_Roulette_Full_50(t *testing.T)           { testGA(50, 500, 45, ROULETTE, false, t) }

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
