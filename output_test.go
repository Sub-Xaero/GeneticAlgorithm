package ga

import (
	"testing"
	"fmt"
)

func TestSetOutputFunc(t *testing.T) {
	t.Parallel()
	var genA GeneticAlgorithm
	var gotOutput string
	output := func (s string) {
		gotOutput = s
	}

	genA.SetOutputFunc(func(a ...interface{}) {
		output(fmt.Sprint(a))
	})

	genA.Output("output string")

	expectedOutput := "[output string]"
	if expectedOutput != gotOutput {
		t.Error("Output func not set. Expected:", expectedOutput, "Got:", gotOutput)
	} else {
		t.Log("Output func set correctly. Expected:", expectedOutput, "Got:", gotOutput)
	}
}