package ga

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSetMutateFunc(t *testing.T) {
	rand.Seed(time.Now().Unix())
	SetMutateFunc(DefaultMutateFunc)
	SetSelectionFunc(TournamentSelection)
	SetFitnessFunc(DefaultFitnessFunc)
	SetGenerateCandidate(DefaultGenerateCandidate)
	SetCrossoverFunc(DefaultCrossoverFunc)

	SetMutateFunc(func(gene Genome) Genome {
		return Genome{[]int{1, 2, 3, 4}}
	})

	output := fmt.Sprint(Genome{[]int{}}.Mutate())
	if output != "{[1 2 3 4],   1}" {
		t.Error("Mutate function not set", output)
	}
}
