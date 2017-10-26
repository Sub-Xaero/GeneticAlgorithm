package ga

import (
	"errors"
	"strconv"
	"testing"
	"time"
)

const (
	ROULETTE   = iota
	TOURNAMENT = iota
)

func TestCheckErrors(t *testing.T) {
	t.Parallel()
	if t.Run("NilError", func(t *testing.T) {
		t.Parallel()
		check(nil)
		defer func() {
			if r := recover(); r == nil {
				t.Log("Panic correctly not thrown on nil error")
			} else {
				t.Error("Panic was thrown for a nil error")
			}
		}()
	}) &&
		t.Run("BadError", func(t *testing.T) {
			t.Parallel()
			err := errors.New("horrible error")
			defer func() {
				if r := recover(); r != nil {
					t.Log("Panic successfully thrown on bad error")
				} else {
					t.Error("Panic either not thrown, or could not recover. ")
				}
			}()
			check(err)
		}) {
		t.Log("Check errors successfull")
	} else {
		t.Error("Check errors failed")
	}
}

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

func testCustomFunctions(geneticAlgorithm *GeneticAlgorithm, t *testing.T) {
	err := geneticAlgorithm.Run(1, 1, 1, true, true, true)
	if err == nil {
		t.Error("GA did not error as expected. Got:", err)
	} else {
		t.Log("GA errored as expected. Got:", err)
	}
}

func TestGACustomFunctions(t *testing.T) {
	t.Parallel()
	success := true
	success = success && t.Run("NilGenerateCandidate", func(t *testing.T) {
		t.Parallel()
		var geneticAlgorithm GeneticAlgorithm
		geneticAlgorithm.SetCrossoverFunc(DefaultCrossoverFunc)
		geneticAlgorithm.SetMutateFunc(DefaultMutateFunc)
		geneticAlgorithm.SetFitnessFunc(DefaultFitnessFunc)
		geneticAlgorithm.SetSelectionFunc(TournamentSelection)
		geneticAlgorithm.SetOutputFunc(PrintToConsole)
		geneticAlgorithm.SetSeed(time.Now().Unix())
		testCustomFunctions(&geneticAlgorithm, t)
	})
	success = success && t.Run("NilCrossover", func(t *testing.T) {
		t.Parallel()
		var geneticAlgorithm GeneticAlgorithm
		geneticAlgorithm.SetGenerateCandidate(DefaultGenerateCandidate)
		geneticAlgorithm.SetMutateFunc(DefaultMutateFunc)
		geneticAlgorithm.SetFitnessFunc(DefaultFitnessFunc)
		geneticAlgorithm.SetSelectionFunc(TournamentSelection)
		geneticAlgorithm.SetOutputFunc(PrintToConsole)
		geneticAlgorithm.SetSeed(time.Now().Unix())
		testCustomFunctions(&geneticAlgorithm, t)
	})
	success = success && t.Run("NilMutate", func(t *testing.T) {
		t.Parallel()
		var geneticAlgorithm GeneticAlgorithm
		geneticAlgorithm.SetGenerateCandidate(DefaultGenerateCandidate)
		geneticAlgorithm.SetCrossoverFunc(DefaultCrossoverFunc)
		geneticAlgorithm.SetFitnessFunc(DefaultFitnessFunc)
		geneticAlgorithm.SetSelectionFunc(TournamentSelection)
		geneticAlgorithm.SetOutputFunc(PrintToConsole)
		geneticAlgorithm.SetSeed(time.Now().Unix())
		testCustomFunctions(&geneticAlgorithm, t)
	})
	success = success && t.Run("NilFitness", func(t *testing.T) {
		t.Parallel()
		var geneticAlgorithm GeneticAlgorithm
		geneticAlgorithm.SetGenerateCandidate(DefaultGenerateCandidate)
		geneticAlgorithm.SetCrossoverFunc(DefaultCrossoverFunc)
		geneticAlgorithm.SetMutateFunc(DefaultMutateFunc)
		geneticAlgorithm.SetSelectionFunc(TournamentSelection)
		geneticAlgorithm.SetOutputFunc(PrintToConsole)
		geneticAlgorithm.SetSeed(time.Now().Unix())
		testCustomFunctions(&geneticAlgorithm, t)
	})
	success = success && t.Run("NilSelection", func(t *testing.T) {
		t.Parallel()
		var geneticAlgorithm GeneticAlgorithm
		geneticAlgorithm.SetGenerateCandidate(DefaultGenerateCandidate)
		geneticAlgorithm.SetCrossoverFunc(DefaultCrossoverFunc)
		geneticAlgorithm.SetMutateFunc(DefaultMutateFunc)
		geneticAlgorithm.SetFitnessFunc(DefaultFitnessFunc)
		geneticAlgorithm.SetOutputFunc(PrintToConsole)
		geneticAlgorithm.SetSeed(time.Now().Unix())
		testCustomFunctions(&geneticAlgorithm, t)
	})
	success = success && t.Run("NilOutput", func(t *testing.T) {
		t.Parallel()
		var geneticAlgorithm GeneticAlgorithm
		geneticAlgorithm.SetGenerateCandidate(DefaultGenerateCandidate)
		geneticAlgorithm.SetCrossoverFunc(DefaultCrossoverFunc)
		geneticAlgorithm.SetMutateFunc(DefaultMutateFunc)
		geneticAlgorithm.SetFitnessFunc(DefaultFitnessFunc)
		geneticAlgorithm.SetSelectionFunc(TournamentSelection)
		geneticAlgorithm.SetSeed(time.Now().Unix())
		testCustomFunctions(&geneticAlgorithm, t)
	})
	success = success && t.Run("NilRandom", func(t *testing.T) {
		t.Parallel()
		var geneticAlgorithm GeneticAlgorithm
		geneticAlgorithm.SetGenerateCandidate(DefaultGenerateCandidate)
		geneticAlgorithm.SetCrossoverFunc(DefaultCrossoverFunc)
		geneticAlgorithm.SetMutateFunc(DefaultMutateFunc)
		geneticAlgorithm.SetFitnessFunc(DefaultFitnessFunc)
		geneticAlgorithm.SetSelectionFunc(TournamentSelection)
		geneticAlgorithm.SetOutputFunc(PrintToConsole)
		testCustomFunctions(&geneticAlgorithm, t)
	})
	success = success && t.Run("AllFuncs", func(t *testing.T) {
		t.Parallel()
		var geneticAlgorithm = NewGeneticAlgorithm()
		geneticAlgorithm.SetOutputFunc(func(a ...interface{}) { t.Log(a) })
		err := geneticAlgorithm.Run(10, 10, 1, true, true, true)
		if err != nil {
			t.Error("GA errored unexpectedly. Got:", err)
		} else {
			t.Log("GA returned nil error as expected. Got:", err)
		}
	})
	if success {
		t.Log("Test GA Custom Functions Succeeded")
	} else {
		t.Error("Test GA Custom Functions Failed")
	}
}

func testGA(length, generations, expectedFitness, selectionMethod int, terminateEarly bool, t *testing.T) {
	t.Parallel()
	var geneticAlgorithm = NewGeneticAlgorithm()

	switch selectionMethod {
	case TOURNAMENT:
		geneticAlgorithm.SetSelectionFunc(TournamentSelection)
	case ROULETTE:
		geneticAlgorithm.SetSelectionFunc(RouletteSelection)
	}
	geneticAlgorithm.SetOutputFunc(func(a ...interface{}) { t.Log(a...) })
	geneticAlgorithm.SetSeed(3)

	geneticAlgorithm.Run(length, length, generations, true, true, terminateEarly)
	geneticAlgorithm.Output(geneticAlgorithm.BestCandidate, geneticAlgorithm.Population)

	gotBestCandidateLength := len(geneticAlgorithm.BestCandidate.Sequence)
	if  gotBestCandidateLength != length {
		t.Error("GA Best Candidate is not correct length", "Expected at least:", length, "Got:", gotBestCandidateLength)
	} else {
		t.Log("GA produced a suitable best candidate of a correct length.", "Expected at least:", length, "Got:", gotBestCandidateLength)
	}

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

func TestGA(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip()
	}
	t.Run("Tournament", func(t *testing.T) {
		t.Parallel()
		for i := 1; i <= 5; i++ {
			t.Run("TerminateEarly"+strconv.Itoa(i*10), func(t *testing.T) { testGA(i*10, i*100, int(float32(i*10.0)*0.89), TOURNAMENT, true, t) })
			t.Run("Full"+strconv.Itoa(i*10), func(t *testing.T) { testGA(i*10, i*100, int(float32(i*10.0)*0.90), TOURNAMENT, false, t) })
		}
	})
	t.Run("Roulette", func(t *testing.T) {
		t.Parallel()
		for i := 1; i <= 5; i++ {
			t.Run("TerminateEarly"+strconv.Itoa(i*10), func(t *testing.T) { testGA(i*10, i*100, int(float32(i*10.0)*0.89), ROULETTE, true, t) })
			t.Run("Full"+strconv.Itoa(i*10), func(t *testing.T) { testGA(i*10, i*100, int(float32(i*10.0)*0.90), ROULETTE, false, t) })
		}
	})
}
