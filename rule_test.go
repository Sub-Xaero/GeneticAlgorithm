package ga

import (
	"bufio"
	"math/rand"
	"os"
	"testing"
)

var InputRuleBase []Rule

func TestRule_String(t *testing.T) {
	t.Parallel()
	rule1 := Rule{bitstring{"1", "0"}, "0"}

	expected := "[1 0] 0"
	got := rule1.String()
	if expected != got {
		t.Error("Rule.toString does not match expected. Expected:", expected, "Got:", got)
	} else {
		t.Log("Rule.toString matches expected. Expected:", expected, "Got:", got)
	}
}

func TestRule_Matches_BadLength(t *testing.T) {
	t.Parallel()
	rule1 := Rule{bitstring{"1", "0", "0"}, "0"}
	rule2 := Rule{bitstring{"1", "0"}, "1"}
	got, err := rule1.Matches(rule2)
	if err == nil {
		t.Error("Expected error, strings are not same length. No error thrown.")
	} else {
		t.Log("Expected error, got error")
	}
	if false != got {
		t.Error("Rules match, expected not to match.", rule1, "!=", rule2, "=", got)
	} else {
		t.Log("Rules do not match, expected not to match.", rule1, "!=", rule2, "=", got)
	}
}

func TestRule_NotMatches(t *testing.T) {
	t.Parallel()
	rule1 := Rule{bitstring{"1", "0"}, "0"}
	rule2 := Rule{bitstring{"1", "0"}, "1"}
	got, err := rule1.Matches(rule2)
	if err != nil {
		t.Error("Did not expect error, error thrown.", err)
	} else {
		t.Log("No errors thrown, Rules match length")
	}
	if false != got {
		t.Error("Rules match, expected not to match.", rule1, "!=", rule2, "=", got)
	} else {
		t.Log("Rules do not match, expected not to match.", rule1, "!=", rule2, "=", got)
	}
}

func TestRule_NotMatchesWildcard(t *testing.T) {
	t.Parallel()
	rule1 := Rule{bitstring{"1", "0"}, "1"}
	rule2 := Rule{bitstring{"#", "1"}, "1"}
	got, err := rule1.Matches(rule2)
	if err != nil {
		t.Error("Did not expect error, error thrown.", err)
	} else {
		t.Log("No errors thrown, Rules match length")
	}
	if false != got {
		t.Error("Rules match, expected not to match.", rule1, "!=", rule2, "=", got)
	} else {
		t.Log("Rules do not match, expected not to match.", rule1, "!=", rule2, "=", got)
	}
}

func TestRule_NotMatchesMultipleWildcard(t *testing.T) {
	t.Parallel()
	rule1 := Rule{bitstring{"1", "1", "0"}, "1"}
	rule2 := Rule{bitstring{"#", "#", "1"}, "1"}
	got, err := rule1.Matches(rule2)
	if err != nil {
		t.Error("Did not expect error, error thrown.", err)
	} else {
		t.Log("No errors thrown, Rules match length")
	}
	if false != got {
		t.Error("Rules match, expected not to match.", rule1, "!=", rule2, "=", got)
	} else {
		t.Log("Rules do not match, expected not to match.", rule1, "!=", rule2, "=", got)
	}
}

func TestRule_Matches(t *testing.T) {
	t.Parallel()
	rule1 := Rule{bitstring{"1", "0"}, "1"}
	rule2 := Rule{bitstring{"1", "0"}, "1"}
	got, err := rule1.Matches(rule2)
	if err != nil {
		t.Error("Did not expect error, error thrown.", err)
	} else {
		t.Log("No errors thrown, Rules match length")
	}

	if true != got {
		t.Error("Rules do not match, expected to match.", rule1, "!=", rule2, "=", got)
	} else {
		t.Log("Rules match, expected to match.", rule1, "!=", rule2, "=", got)
	}
}

func TestRule_MatchesWildcard(t *testing.T) {
	t.Parallel()
	rule1 := Rule{bitstring{"1", "0", "#"}, "1"}
	rule2 := Rule{bitstring{"1", "0", "1"}, "1"}
	got, err := rule1.Matches(rule2)
	if err != nil {
		t.Error("Did not expect error, error thrown.", err)
	} else {
		t.Log("No errors thrown, Rules match length")
	}

	if true != got {
		t.Error("Rules do not match, expected to match.", rule1, "!=", rule2, "=", got)
	} else {
		t.Log("Rules match, expected to match.", rule1, "!=", rule2, "=", got)
	}
}

func TestRule_MatchesMultipleWildcards(t *testing.T) {
	t.Parallel()
	rule1 := Rule{bitstring{"#", "0", "#"}, "1"}
	rule2 := Rule{bitstring{"1", "0", "1"}, "1"}
	got, err := rule1.Matches(rule2)
	if err != nil {
		t.Error("Did not expect error, error thrown.", err)
	} else {
		t.Log("No errors thrown, Rules match length")
	}

	if true != got {
		t.Error("Rules do not match, expected to match.", rule1, "!=", rule2, "=", got)
	} else {
		t.Log("Rules match, expected to match.", rule1, "!=", rule2, "=", got)
	}
}

func TestRuleGA(t *testing.T) {
	var geneticAlgorithm = NewGeneticAlgorithm()
	geneticAlgorithm.SetSeed(3)
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
				matches, err := InputRule.Matches(GeneratedRule)
				check(err)
				if matches {
					fitnessValue++
					break
				}
			}
		}
		return fitnessValue
	})

	geneticAlgorithm.SetMutateFunc(func(gene Genome, random *rand.Rand) Genome {
		choice := random.Int() % len(gene.Sequence)
		choice2 := random.Int() % 2

		switch gene.Sequence[choice] {
		case "0":
			switch choice2 {
			case 0:
				gene.Sequence[choice] = "1"
			case 1:
				gene.Sequence[choice] = "#"
			}
		case "1":
			switch choice2 {
			case 0:
				gene.Sequence[choice] = "0"
			case 1:
				gene.Sequence[choice] = "#"
			}
		case "#":
			switch choice2 {
			case 0:
				gene.Sequence[choice] = "0"
			case 1:
				gene.Sequence[choice] = "1"
			}
		}
		return gene
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
