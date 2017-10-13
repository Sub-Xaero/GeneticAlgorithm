package ga

import (
	"fmt"
	"testing"
)

func TestDefaultGenerateCandidate(t *testing.T) {
	bitstring := GenerateCandidate(10)

	if len(bitstring) != 10 {
		t.Error("String is not correct length")
	}
	sum := 0
	for _, i := range bitstring {
		sum += i
	}
	if sum >= 20 {
		t.Error("String is not correct value")
	}
}

func TestCustomGenerateCandidate(t *testing.T) {
	SetGenerateCandidate(func(length int) []int {
		var sequence []int
		for i := 1; i <= length; i++ {
			sequence = append(sequence, i)
		}
		return sequence
	})
	bitstring := GenerateCandidate(9)

	if len(bitstring) != 9 {
		t.Error("String is not correct length", bitstring)
	}

	expected := "[1 2 3 4 5 6 7 8 9]"
	if fmt.Sprint(bitstring) != expected {
		t.Error("String is not correct string:", bitstring, "Expected:", expected)
	}
}

func TestGenome_ToString(t *testing.T) {
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
