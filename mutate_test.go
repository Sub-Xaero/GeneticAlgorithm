package ga

import (
	"fmt"
	"testing"
)

func TestSetMutateFunc(t *testing.T) {
	SetFitnessFunc(DefaultFitnessFunc)
	SetMutateFunc(func(gene Genome) Genome {
		return Genome{[]int{1, 2, 3, 4}}
	})

	output := fmt.Sprint(Genome{[]int{}}.Mutate())
	if output != "{[1 2 3 4],   1}" {
		t.Error("Mutate function not set", output)
	}
}
