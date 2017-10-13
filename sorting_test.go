package ga

import (
	"fmt"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
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
	sort.Sort(ByFitness(population))

	if len(population) != len(after) {
		t.Error("Populations do not match length")
	}

	for i := range population {
		left := fmt.Sprint(population[i])
		right := fmt.Sprint(after[len(after)-i-1])
		if left != right {
			t.Error("Not sorted")
			break
		}
	}

}
