package ga

import (
	"bufio"
	"math/rand"
	"os"
	"testing"
)

var InputRuleBase []Rule

func TestRuleGA(t *testing.T) {
	rand.Seed(3)
	var geneticAlgorithm = NewGeneticAlgorithm()
	geneticAlgorithm.SetOutputFunc(func(a ...interface{}) { t.Log(a) })

	InputRuleBase = make([]Rule, 0)

	//ConditionLength := 5

	file, err := os.OpenFile("data.txt", os.O_RDONLY, 7777)
	scanner := bufio.NewScanner(file)
	if err != nil {
		t.Error(err)
	}
	for i := 0; scanner.Scan(); i++ {
		if i == 0 {
			continue
		}
		text := scanner.Text()
		ruleSequence := make(bitstring, 0)
		for char := 0; char < 5; char++ {
			num := string(text[char])
			check(err)
			ruleSequence = append(ruleSequence, num)
		}
		output := string(text[6])
		check(err)
		InputRuleBase = append(InputRuleBase, Rule{ruleSequence, output})
	}

	geneticAlgorithm.SetFitnessFunc(func(gene Genome) int {
		NewRuleBase := make([]Rule, 0)
		fitnessValue := 0
		for i := 0; i < len(gene.Sequence)-6-1; i += 6 {
			condition := make(bitstring, len(gene.Sequence[i:i+5]))
			copy(condition, gene.Sequence[i:i+5])
			rule := Rule{condition, gene.Sequence[i+5]}
			NewRuleBase = append(NewRuleBase, rule)
		}
		for _, InputRule := range InputRuleBase {
			for _, GeneratedRule := range NewRuleBase {
				if InputRule.Matches(GeneratedRule) {
					fitnessValue++
					break
				}
			}
		}
		return fitnessValue
	})

	geneticAlgorithm.Run(10, 60, 10, true, true, false)
	geneticAlgorithm.Output(geneticAlgorithm.BestCandidate, geneticAlgorithm.Population)

	expectedFitness := 8
	gotFitness := geneticAlgorithm.Fitness(geneticAlgorithm.BestCandidate)

	if gotFitness < expectedFitness {
		t.Error("GA did not produce a suitable candidate.", "Expected at least:", expectedFitness, "Got:", gotFitness)
	} else {
		t.Log("GA produced a suitable candidate.", "Expected at least:", expectedFitness, "Got:", gotFitness)
	}
}
