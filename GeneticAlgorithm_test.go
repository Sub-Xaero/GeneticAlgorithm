package ga

import (
	"testing"
)

const (
	ROULETTE   = iota
	TOURNAMENT = iota
)

func TestFillRandomPopulation(t *testing.T) {
	t.Parallel()
	var geneticAlgorithm = NewGeneticAlgorithm()
	geneticAlgorithm.SetSeed(3)
	geneticAlgorithm.SetOutputFunc(func(a ...interface{}) { t.Log(a...) })

	expectedLen := 10
	population := make([]Genome, 0)
	population = geneticAlgorithm.FillRandomPopulation(expectedLen, expectedLen)
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

func TestGA_NilGenerateCandidateFunc(t *testing.T) {
	t.Parallel()
	var geneticAlgorithm GeneticAlgorithm
	geneticAlgorithm.SetCrossoverFunc(DefaultCrossoverFunc)
	geneticAlgorithm.SetMutateFunc(DefaultMutateFunc)
	geneticAlgorithm.SetFitnessFunc(DefaultFitnessFunc)
	geneticAlgorithm.SetSelectionFunc(TournamentSelection)
	geneticAlgorithm.SetOutputFunc(PrintToConsole)
	geneticAlgorithm.SetOutputFunc(PrintToConsole)
	geneticAlgorithm.SetSeed(3)
	err := geneticAlgorithm.Run(1, 1, 1, true, true, true)

	if err == nil {
		t.Error("GA did not error as expected. Got:", err)
	} else {
		t.Log("GA errored as expected. Got:", err)
	}
}

func TestGA_NilCrossoverFunc(t *testing.T) {
	t.Parallel()
	var geneticAlgorithm GeneticAlgorithm
	geneticAlgorithm.SetGenerateCandidate(DefaultGenerateCandidate)
	geneticAlgorithm.SetMutateFunc(DefaultMutateFunc)
	geneticAlgorithm.SetFitnessFunc(DefaultFitnessFunc)
	geneticAlgorithm.SetSelectionFunc(TournamentSelection)
	geneticAlgorithm.SetOutputFunc(PrintToConsole)
	geneticAlgorithm.SetSeed(3)
	err := geneticAlgorithm.Run(1, 1, 1, true, true, true)

	if err == nil {
		t.Error("GA did not error as expected. Got:", err)
	} else {
		t.Log("GA errored as expected. Got:", err)
	}
}

func TestGA_NilMutateFunc(t *testing.T) {
	t.Parallel()
	var geneticAlgorithm GeneticAlgorithm
	geneticAlgorithm.SetGenerateCandidate(DefaultGenerateCandidate)
	geneticAlgorithm.SetCrossoverFunc(DefaultCrossoverFunc)
	geneticAlgorithm.SetFitnessFunc(DefaultFitnessFunc)
	geneticAlgorithm.SetSelectionFunc(TournamentSelection)
	geneticAlgorithm.SetOutputFunc(PrintToConsole)
	geneticAlgorithm.SetSeed(3)
	err := geneticAlgorithm.Run(1, 1, 1, true, true, true)

	if err == nil {
		t.Error("GA did not error as expected. Got:", err)
	} else {
		t.Log("GA errored as expected. Got:", err)
	}
}

func TestGA_NilFitnessFunc(t *testing.T) {
	t.Parallel()
	var geneticAlgorithm GeneticAlgorithm
	geneticAlgorithm.SetGenerateCandidate(DefaultGenerateCandidate)
	geneticAlgorithm.SetCrossoverFunc(DefaultCrossoverFunc)
	geneticAlgorithm.SetMutateFunc(DefaultMutateFunc)
	geneticAlgorithm.SetSelectionFunc(TournamentSelection)
	geneticAlgorithm.SetOutputFunc(PrintToConsole)
	geneticAlgorithm.SetSeed(3)
	err := geneticAlgorithm.Run(1, 1, 1, true, true, true)

	if err == nil {
		t.Error("GA did not error as expected. Got:", err)
	} else {
		t.Log("GA errored as expected. Got:", err)
	}
}

func TestGA_NilSelectionFunc(t *testing.T) {
	t.Parallel()
	var geneticAlgorithm GeneticAlgorithm
	geneticAlgorithm.SetGenerateCandidate(DefaultGenerateCandidate)
	geneticAlgorithm.SetCrossoverFunc(DefaultCrossoverFunc)
	geneticAlgorithm.SetMutateFunc(DefaultMutateFunc)
	geneticAlgorithm.SetFitnessFunc(DefaultFitnessFunc)
	geneticAlgorithm.SetOutputFunc(PrintToConsole)
	geneticAlgorithm.SetSeed(3)
	err := geneticAlgorithm.Run(1, 1, 1, true, true, true)

	if err == nil {
		t.Error("GA did not error as expected. Got:", err)
	} else {
		t.Log("GA errored as expected. Got:", err)
	}
}

func TestGA_NilOutputFunc(t *testing.T) {
	t.Parallel()
	var geneticAlgorithm GeneticAlgorithm
	geneticAlgorithm.SetGenerateCandidate(DefaultGenerateCandidate)
	geneticAlgorithm.SetCrossoverFunc(DefaultCrossoverFunc)
	geneticAlgorithm.SetMutateFunc(DefaultMutateFunc)
	geneticAlgorithm.SetFitnessFunc(DefaultFitnessFunc)
	geneticAlgorithm.SetSelectionFunc(TournamentSelection)
	geneticAlgorithm.SetSeed(3)
	err := geneticAlgorithm.Run(1, 1, 1, true, true, true)

	if err == nil {
		t.Error("GA did not error as expected. Got:", err)
	} else {
		t.Log("GA errored as expected. Got:", err)
	}
}

func TestGA_NilRandom(t *testing.T) {
	t.Parallel()
	var geneticAlgorithm GeneticAlgorithm
	geneticAlgorithm.SetGenerateCandidate(DefaultGenerateCandidate)
	geneticAlgorithm.SetCrossoverFunc(DefaultCrossoverFunc)
	geneticAlgorithm.SetMutateFunc(DefaultMutateFunc)
	geneticAlgorithm.SetFitnessFunc(DefaultFitnessFunc)
	geneticAlgorithm.SetSelectionFunc(TournamentSelection)
	geneticAlgorithm.SetOutputFunc(PrintToConsole)
	err := geneticAlgorithm.Run(1, 1, 1, true, true, true)

	if err == nil {
		t.Error("GA did not error as expected. Got:", err)
	} else {
		t.Log("GA errored as expected. Got:", err)
	}
}

func testGA(length, generations, expectedFitness, selectionMethod int, terminateEarly bool, t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip()
	}
	var geneticAlgorithm = NewGeneticAlgorithm()

	switch selectionMethod {
	case TOURNAMENT:
		geneticAlgorithm.SetSelectionFunc(TournamentSelection)
	case ROULETTE:
		geneticAlgorithm.SetSelectionFunc(RouletteSelection)
	}
	geneticAlgorithm.SetOutputFunc(func(a ...interface{}) { t.Log(a...) })

	geneticAlgorithm.Run(length, length, generations, true, true, terminateEarly)
	geneticAlgorithm.Output(geneticAlgorithm.BestCandidate, geneticAlgorithm.Population)

	gotFitness := geneticAlgorithm.Fitness(geneticAlgorithm.BestCandidate)
	if gotFitness < expectedFitness {
		t.Error("GA did not produce a suitable candidate.", "Expected at least:", expectedFitness, "Got:", gotFitness)
	} else {
		t.Log("GA produced a suitable candidate.", "Expected at least:", expectedFitness, "Got:", gotFitness)
	}

	if terminateEarly && geneticAlgorithm.Generations < generations {
		t.Log("GA successfully detected stagnation. Terminated early. Expected less than", generations, "iterations, GA took", geneticAlgorithm.Generations)
	} else if terminateEarly {
		t.Error("GA did not terminate early. Expected less than", generations, "iterations, GA took", geneticAlgorithm.Generations)
	}
}

func TestGA_Tournament_TerminateEarly_10(t *testing.T) { testGA(10, 100, 8, TOURNAMENT, true, t) }
func TestGA_Tournament_TerminateEarly_20(t *testing.T) { testGA(20, 200, 15, TOURNAMENT, true, t) }
func TestGA_Tournament_TerminateEarly_30(t *testing.T) { testGA(30, 300, 25, TOURNAMENT, true, t) }
func TestGA_Tournament_TerminateEarly_40(t *testing.T) { testGA(40, 400, 35, TOURNAMENT, true, t) }
func TestGA_Tournament_TerminateEarly_50(t *testing.T) { testGA(50, 500, 43, TOURNAMENT, true, t) }
func TestGA_Tournament_Full_10(t *testing.T)           { testGA(10, 100, 8, TOURNAMENT, false, t) }
func TestGA_Tournament_Full_20(t *testing.T)           { testGA(20, 200, 15, TOURNAMENT, false, t) }
func TestGA_Tournament_Full_30(t *testing.T)           { testGA(30, 300, 25, TOURNAMENT, false, t) }
func TestGA_Tournament_Full_40(t *testing.T)           { testGA(40, 400, 35, TOURNAMENT, false, t) }
func TestGA_Tournament_Full_50(t *testing.T)           { testGA(50, 500, 40, TOURNAMENT, false, t) }

func TestGA_Roulette_TerminateEarly_10(t *testing.T) { testGA(10, 100, 8, ROULETTE, true, t) }
func TestGA_Roulette_TerminateEarly_20(t *testing.T) { testGA(20, 200, 15, ROULETTE, true, t) }
func TestGA_Roulette_TerminateEarly_30(t *testing.T) { testGA(30, 300, 25, ROULETTE, true, t) }
func TestGA_Roulette_TerminateEarly_40(t *testing.T) { testGA(40, 400, 35, ROULETTE, true, t) }
func TestGA_Roulette_TerminateEarly_50(t *testing.T) { testGA(50, 500, 40, ROULETTE, true, t) }
func TestGA_Roulette_Full_10(t *testing.T)           { testGA(10, 100, 8, ROULETTE, false, t) }
func TestGA_Roulette_Full_20(t *testing.T)           { testGA(20, 200, 15, ROULETTE, false, t) }
func TestGA_Roulette_Full_30(t *testing.T)           { testGA(30, 300, 25, ROULETTE, false, t) }
func TestGA_Roulette_Full_40(t *testing.T)           { testGA(40, 400, 35, ROULETTE, false, t) }
func TestGA_Roulette_Full_50(t *testing.T)           { testGA(50, 500, 40, ROULETTE, false, t) }
