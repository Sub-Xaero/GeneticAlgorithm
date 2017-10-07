package ga

import (
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestSort(t *testing.T) {
	SetFitnessFunc(func(gene Genome) int {
		return strings.Count(gene.Sequence, "1")
	})

	population := []Genome{
		{"111111"},
		{"011111"},
		{"001111"},
		{"000111"},
		{"000011"},
		{"000001"},
	}

	after := []Genome{
		{"111111"},
		{"011111"},
		{"001111"},
		{"000111"},
		{"000011"},
		{"000001"},
	}
	sort.Sort(ByFitness(population))

	if len(population) != len(after) {
		t.Error("Populations do not match length")
	}

	for i := range population {
		if population[i] != after[len(after)-i-1] {
			t.Error("Not sorted")
			break
		}
	}

}

func TestDefaultGenerateBitString(t *testing.T) {
	bitstring := GenerateBitString(10)

	if len(bitstring) != 10 {
		t.Error("String is not correct length")
	}
	_, err := strconv.ParseInt(bitstring, 2, 0)
	if err != nil {
		t.Error("String is not correct length")
	}
}

func TestCustomGenerateBitString(t *testing.T) {
	SetGenerateBitString(func(length int) string {
		var bitstring = ""
		for i := 1; i <= length; i++ {
			bitstring += strconv.Itoa(i)
		}
		return bitstring
	})
	bitstring := GenerateBitString(9)

	if len(bitstring) != 9 {
		t.Error("String is not correct length", bitstring)
	}

	expected := "123456789"
	if bitstring != expected {
		t.Error("String is not correct string:", bitstring, "Expected:", expected)
	}
}

func TestGenome_ToString(t *testing.T) {
	SetFitnessFunc(func(gene Genome) int {
		return strings.Count(gene.Sequence, "1")
	})

	outputString := Genome{"1111"}.String()
	expected := "{1111,   4}"
	if outputString != expected {
		t.Error("Incorrect string:", outputString, "Expected:", expected)
	}

	outputString = Genome{"1010101010"}.String()
	expected = "{1010101010,   5}"
	if outputString != expected {
		t.Error("Incorrect string:", outputString, "Expected:", expected)
	}

	outputString = Genome{"1111111111"}.String()
	expected = "{1111111111,  10}"
	if outputString != expected {
		t.Error("Incorrect string:", outputString, "Expected:", expected)
	}

	outputString = Genome{"111111111111"}.String()
	expected = "12"
	if outputString != expected {
		t.Error("Incorrect string:", outputString, "Expected:", expected)
	}
}
