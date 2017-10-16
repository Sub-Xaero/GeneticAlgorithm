package ga

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestDefaultGenerateCandidate(t *testing.T) {
	rand.Seed(time.Now().Unix())
	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(TournamentSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)

	expectedLength := 10
	bitstring, err := GenerateCandidate(expectedLength)

	if err == nil {
		t.Log("Successfully generated candidate")
	} else {
		t.Error("Unexpected error:", err)
	}

	length := len(bitstring)
	if length != expectedLength {
		t.Error("String is not correct length.", "Expected:", expectedLength, "Got:", length)
	} else {
		t.Log("String is correct length", "Expected:", expectedLength, "Got:", length)
	}

	expectedSum := 20
	sum := 0
	for _, i := range bitstring {
		sum += i
	}
	if sum >= 20 {
		t.Error("String is not correct value.", "Expected:", expectedSum, "Got:", sum)
	} else {
		t.Log("String is correct length.", "Expected:", expectedSum, "Got:", sum)
	}
}

func TestBadDefaultGenerateCandidate(t *testing.T) {
	rand.Seed(time.Now().Unix())
	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(TournamentSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)

	expectedLength := 0
	_, err := GenerateCandidate(expectedLength)

	if err == nil {
		t.Error("Bad candidate length did not throw error as it should. Err:", err)
	} else {
		t.Log("Successfully threw and caught error:", err)
	}
}

func TestCustomGenerateCandidate(t *testing.T) {
	rand.Seed(time.Now().Unix())
	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(TournamentSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)

	SetGenerateCandidate(func(length int) ([]int, error) {
		var sequence []int
		for i := 1; i <= length; i++ {
			sequence = append(sequence, i)
		}
		return sequence, nil
	})

	expectedLength := 9
	bitstring, err := GenerateCandidate(expectedLength)

	if err == nil {
		t.Log("Successfully generated candidate")
	} else {
		t.Error("Unexpected error:", err)
	}

	length := len(bitstring)
	if length != expectedLength {
		t.Error("String is not correct length.", "Expected:", expectedLength, "Got:", length)
	} else {
		t.Log("String is correct length. Expected:", expectedLength, "Got:", length)
	}

	expected := "[1 2 3 4 5 6 7 8 9]"
	if fmt.Sprint(bitstring) != expected {
		t.Error("String is not correct string.", "Expected:", expected, "Got:", bitstring)
	} else {
		t.Log("String is correct string.", "Expected:", expected, "Got:", bitstring)
	}
}

func TestGenome_ToString(t *testing.T) {
	rand.Seed(time.Now().Unix())
	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(TournamentSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)

	SetFitnessFunc(func(gene Genome) int {
		count := 0
		for _, i := range gene.Sequence {
			if i == 1 {
				count++
			}
		}
		return count
	})

	outputString := Genome{[]int{1, 1, 1, 1}}.String()
	expected := "{[1 1 1 1],   4}"
	if outputString != expected {
		t.Error("Incorrect string:", outputString, "Expected:", expected)
	}

	outputString = Genome{[]int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0}}.String()
	expected = "{[1 0 1 0 1 0 1 0 1 0],   5}"
	if outputString != expected {
		t.Error("Incorrect string:", outputString, "Expected:", expected)
	}

	outputString = Genome{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}.String()
	expected = "{[1 1 1 1 1 1 1 1 1 1],  10}"
	if outputString != expected {
		t.Error("Incorrect string:", outputString, "Expected:", expected)
	}

	outputString = Genome{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}.String()
	expected = "12"
	if outputString != expected {
		t.Error("Incorrect string:", outputString, "Expected:", expected)
	}
}
