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
	SetOutputFunc(func(a ...interface{}) { t.Log(a...) })

	SetMutateFunc(func(gene Genome) Genome {
		return Genome{[]int{1, 2, 3, 4}}
	})

	output := fmt.Sprint(Genome{[]int{}}.Mutate())
	expectedOutput := "{[1 2 3 4],   1}"
	if output != expectedOutput {
		t.Error("Mutate function not set. Expected:", expectedOutput, "Got:", output)
	} else {
		t.Log("Mutate function was set successfully. Expected:", expectedOutput, "Got:", output)
	}
}
