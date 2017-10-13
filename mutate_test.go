package ga

import (
	"fmt"
	"testing"
)

func TestSetMutateFunc(t *testing.T) {
	SetMutateFunc(func(gene Genome) Genome {
		return Genome{[]int{1, 2, 3, 4}}
	})

	if fmt.Sprint(Genome{[]int{}}.Mutate()) != "{[1 2 3 4],   1}" {
		t.Error("Mutate function not set")
	}
}
