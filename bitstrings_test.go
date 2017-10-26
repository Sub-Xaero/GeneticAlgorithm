package ga

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func TestDefaultGenerateCandidate(t *testing.T) {
	t.Parallel()
	var genA = NewGeneticAlgorithm()
	genA.SetSeed(3)
	genA.SetOutputFunc(func(a ...interface{}) { t.Log(a...) })

	expectedLength := 10
	bitstring, err := genA.GenerateCandidate(expectedLength, genA.RandomEngine)

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
		val, err := strconv.Atoi(string(i))
		check(err)
		sum += val
	}
	if sum >= 20 {
		t.Error("String is not correct value.", "Expected:", expectedSum, "Got:", sum)
	} else {
		t.Log("String is correct length.", "Expected:", expectedSum, "Got:", sum)
	}
}

func TestGenerateCandidate (t *testing.T) {
	t.Parallel()
	expectedLength := 9
	setup := func () (Bitstring, error) {
		var genA= NewGeneticAlgorithm()
		genA.SetGenerateCandidate(func(length int, random *rand.Rand) (Bitstring, error) {
			var sequence Bitstring
			for i := 1; i <= length; i++ {
				sequence = append(sequence, strconv.Itoa(i))
			}
			return sequence, nil
		})
		return genA.GenerateCandidate(expectedLength, genA.RandomEngine)
	}

	t.Run("BadLength", func(t *testing.T) {
		t.Parallel()
		var genA = NewGeneticAlgorithm()
		genA.SetSeed(3)
		genA.SetOutputFunc(func(a ...interface{}) { t.Log(a...) })
		expected := 0
		_, err := genA.GenerateCandidate(expected, genA.RandomEngine)
		if err == nil {
			t.Error("Bad candidate length did not throw error as it should. Err:", err)
		} else {
			t.Log("Successfully threw and caught error:", err)
		}
	})
	t.Run("GeneratesCandidate", func(t *testing.T) {
		t.Parallel()
		_, err := setup()
		if err == nil {
			t.Log("Successfully generated candidate")
		} else {
			t.Error("Unexpected error:", err)
		}
	})
	t.Run("CorrectStringLength", func(t *testing.T) {
		t.Parallel()
		sequence, _ := setup()
		length := len(sequence)
		if length != expectedLength {
			t.Error("String is not correct length.", "Expected:", expectedLength, "Got:", length)
		} else {
			t.Log("String is correct length. Expected:", expectedLength, "Got:", length)
		}
	})
	t.Run("CorrectStringOutput", func(t *testing.T) {
		t.Parallel()
		sequence, _ := setup()
		expected := "[1 2 3 4 5 6 7 8 9]"
		if fmt.Sprint(sequence) != expected {
			t.Error("String is not correct string.", "Expected:", expected, "Got:", sequence)
		} else {
			t.Log("String is correct string.", "Expected:", expected, "Got:", sequence)
		}
	})
}