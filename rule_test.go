package ga

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"testing"
)

var InputRuleBase []Rule

func TestSetRulesMatchFunc(t *testing.T) {
	t.Parallel()
	var genA = NewGeneticAlgorithm()
	genA.SetSeed(3)
	genA.SetOutputFunc(func(a ...interface{}) { t.Log(a...) })

	genA.SetMutateFunc(func(gene Genome, random *rand.Rand) Genome {
		return Genome{Bitstring{"1", "2", "3", "4"}}
	})

	output := fmt.Sprint(genA.Mutate(Genome{}, genA.RandomEngine))
	expectedOutput := "{[1 2 3 4 ]}"
	if output != expectedOutput {
		t.Error("Mutate function not set. Expected:", expectedOutput, "Got:", output)
	} else {
		t.Log("Mutate function was set successfully. Expected:", expectedOutput, "Got:", output)
	}
}

func TestDefaultRulesMatchFunc(t *testing.T) {
	t.Parallel()
	genA := NewGeneticAlgorithm()
	gene := Genome{Bitstring{"1", "0", "1", "0", "1"}}
	geneOutput := genA.Mutate(gene, genA.RandomEngine)

	if gene.String() == geneOutput.String() {
		t.Error("Mutate did not change bitstrings. At least one mutation should occur. Was:", gene, "Mutated:", geneOutput)
	} else {
		t.Log("Mutate successfully changed bitstrings. At least one mutation should occur. Was:", gene, "Mutated:", geneOutput)
	}
}

func TestRule_String(t *testing.T) {
	t.Parallel()
	rule1 := Rule{Bitstring{"1", "0"}, "0"}

	expected := "[1 0 ] 0"
	got := rule1.String()
	if expected != got {
		t.Error("Rule.toString does not match expected. Expected:", expected, "Got:", got)
	} else {
		t.Log("Rule.toString matches expected. Expected:", expected, "Got:", got)
	}
}

func TestRule_DefaultRulesMatch(t *testing.T) {
	t.Parallel()
	genA := NewGeneticAlgorithm()
	checkResult := func(rule1, rule2 Rule, expectedResult, result bool) {
		if expectedResult != result {
			t.Error("Rules match, expected not to match.", rule1, "!=", rule2, "=", result)
		} else {
			t.Log("Rules do not match, expected not to match.", rule1, "!=", rule2, "=", result)
		}
	}

	t.Run("MatchesBadLength", func(t *testing.T) {
		t.Parallel()
		rule1 := Rule{Bitstring{"1", "0", "0"}, "0"}
		rule2 := Rule{Bitstring{"1", "0"}, "1"}
		got, err := genA.RulesMatch(rule1, rule2)
		if err == nil {
			t.Error("Expected error, strings are not same length. No error thrown.")
		} else {
			t.Log("Expected error, got error")
		}
		checkResult(rule1, rule2, false, got)
	})
	t.Run("NotMatches", func(t *testing.T) {
		t.Parallel()
		rule1 := Rule{Bitstring{"1", "0"}, "0"}
		rule2 := Rule{Bitstring{"1", "0"}, "1"}
		got, err := genA.RulesMatch(rule1, rule2)
		if err != nil {
			t.Error("Did not expect error, error thrown.", err)
		} else {
			t.Log("No errors thrown, Rules match length")
		}
		checkResult(rule1, rule2, false, got)
	})
	t.Run("NotMatchesWildcard", func(t *testing.T) {
		t.Parallel()
		rule1 := Rule{Bitstring{"1", "0"}, "1"}
		rule2 := Rule{Bitstring{"#", "1"}, "1"}
		got, err := genA.RulesMatch(rule1, rule2)
		if err != nil {
			t.Error("Did not expect error, error thrown.", err)
		} else {
			t.Log("No errors thrown, Rules match length")
		}
		checkResult(rule1, rule2, false, got)
	})
	t.Run("NotMatchesMultipleWildcard", func(t *testing.T) {
		t.Parallel()
		rule1 := Rule{Bitstring{"1", "1", "0"}, "1"}
		rule2 := Rule{Bitstring{"#", "#", "1"}, "1"}
		got, err := genA.RulesMatch(rule1, rule2)
		if err != nil {
			t.Error("Did not expect error, error thrown.", err)
		} else {
			t.Log("No errors thrown, Rules match length")
		}
		checkResult(rule1, rule2, false, got)
	})
	t.Run("Matches", func(t *testing.T) {
		t.Parallel()
		rule1 := Rule{Bitstring{"1", "0"}, "1"}
		rule2 := Rule{Bitstring{"1", "0"}, "1"}
		got, err := genA.RulesMatch(rule1, rule2)
		if err != nil {
			t.Error("Did not expect error, error thrown.", err)
		} else {
			t.Log("No errors thrown, Rules match length")
		}
		checkResult(rule1, rule2, true, got)
	})
	t.Run("MatchesWildcard", func(t *testing.T) {
		t.Parallel()
		rule1 := Rule{Bitstring{"1", "0", "#"}, "1"}
		rule2 := Rule{Bitstring{"1", "0", "1"}, "1"}
		got, err := genA.RulesMatch(rule1, rule2)
		if err != nil {
			t.Error("Did not expect error, error thrown.", err)
		} else {
			t.Log("No errors thrown, Rules match length")
		}
		checkResult(rule1, rule2, true, got)
	})
	t.Run("MatchesMultipleWildcard", func(t *testing.T) {
		t.Parallel()
		rule1 := Rule{Bitstring{"#", "0", "#"}, "1"}
		rule2 := Rule{Bitstring{"1", "0", "1"}, "1"}
		got, err := genA.RulesMatch(rule1, rule2)
		if err != nil {
			t.Error("Did not expect error, error thrown.", err)
		} else {
			t.Log("No errors thrown, Rules match length")
		}
		checkResult(rule1, rule2, true, got)
	})
}

func TestRuleGA(t *testing.T) {
	var geneticAlgorithm = NewGeneticAlgorithm()
	geneticAlgorithm.SetSeed(3)
	geneticAlgorithm.SetOutputFunc(func(a ...interface{}) { t.Log(a) })

	conditionLength := 5
	outputLength := 1
	ruleLength := conditionLength + outputLength
	numRules := 32

	InputRuleBase = make([]Rule, 0)

	file, err := os.OpenFile("data/data1.txt", os.O_RDONLY, 7777)
	scanner := bufio.NewScanner(file)
	if err != nil {
		t.Error(err)
	}
	for i := 0; scanner.Scan(); i++ {
		if i == 0 {
			continue
		}
		text := scanner.Text()
		ruleSequence := make(Bitstring, 0)
		for char := 0; char < conditionLength; char++ {
			num := string(text[char])
			check(err)
			ruleSequence = append(ruleSequence, num)
		}
		output := string(text[conditionLength+1:])
		check(err)
		InputRuleBase = append(InputRuleBase, Rule{ruleSequence, output})
	}

	geneticAlgorithm.SetFitnessFunc(func(gene Genome) int {
		fitnessValue := 0
		NewRuleBase, err := geneticAlgorithm.DecodeRules(gene.Sequence, conditionLength, ruleLength)
		check(err)
		for _, InputRule := range InputRuleBase {
			for _, GeneratedRule := range NewRuleBase {
				matches, err := geneticAlgorithm.RulesMatch(InputRule, GeneratedRule)
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
		NewRuleBase, err := geneticAlgorithm.DecodeRules(gene.Sequence, conditionLength, ruleLength)
		check(err)
		gene = gene.Copy()
		for rule := range NewRuleBase {
			chance := random.Int() % 100
			if chance < 5 {
				choice2 := random.Int() % len(NewRuleBase[rule].Condition)
				operators := []string{"0", "1", "#"}
				for index, val := range operators {
					if NewRuleBase[rule].Condition[choice2] == val {
						operators = append(operators[0:index], operators[index+1:]...)
					}
				}
				choice3 := random.Int() % len(operators)
				NewRuleBase[rule].Condition[choice2] = operators[choice3]
			}
		}
		sequence, err := geneticAlgorithm.EncodeRules(NewRuleBase)
		check(err)
		return Genome{sequence}
	})

	geneticAlgorithm.Run(20, numRules*ruleLength, 20, true, true, false)
	geneticAlgorithm.Output(geneticAlgorithm.BestCandidate, geneticAlgorithm.Candidates)

	geneticAlgorithm.Output(geneticAlgorithm.DecodeRules(geneticAlgorithm.BestCandidate.Sequence, conditionLength, ruleLength))

	expectedFitness := 26
	gotFitness := geneticAlgorithm.Fitness(geneticAlgorithm.BestCandidate)

	if gotFitness < expectedFitness {
		t.Error("GA did not produce a suitable candidate.", "Expected at least:", expectedFitness, "Got:", gotFitness)
	} else {
		t.Log("GA produced a suitable candidate.", "Expected at least:", expectedFitness, "Got:", gotFitness)
	}
}
