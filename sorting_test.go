package ga

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
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

	population := []Genome{
		{[]int{1, 1, 1, 1, 1, 1}},
		{[]int{0, 1, 1, 1, 1, 1}},
		{[]int{0, 0, 1, 1, 1, 1}},
		{[]int{0, 0, 0, 1, 1, 1}},
		{[]int{0, 0, 0, 0, 1, 1}},
		{[]int{0, 0, 0, 0, 0, 1}},
	}

	after := []Genome{
		{[]int{1, 1, 1, 1, 1, 1}},
		{[]int{0, 1, 1, 1, 1, 1}},
		{[]int{0, 0, 1, 1, 1, 1}},
		{[]int{0, 0, 0, 1, 1, 1}},
		{[]int{0, 0, 0, 0, 1, 1}},
		{[]int{0, 0, 0, 0, 0, 1}},
	}

	t.Log("Population Before:", population)
	sort.Sort(ByFitness(population))
	t.Log("Population After:", population)

	if len(population) != len(after) {
		t.Error("Populations do not match length")
	} else {
		t.Log("Population lengths match")
	}

	sorted := true
	for i := range population {
		left := fmt.Sprint(population[i])
		right := fmt.Sprint(after[len(after)-i-1])
		if left != right {
			t.Error("Population is not sorted")
			sorted = false
			break
		}
	}
	if sorted {
		t.Log("Successfully sorted population")
	}

}
