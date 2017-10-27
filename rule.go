package ga

import (
	"errors"
	"fmt"
)

type Rule struct {
	Condition Bitstring
	Output    string
}

func (rule1 Rule) Matches(rule2 Rule) (bool, error) {
	if len(rule1.Condition) != len(rule2.Condition) {
		return false, errors.New("conditions are not same length")
	}
	conditionMatches := true
	for i := range rule1.Condition {
		if string(rule1.Condition[i]) == "#" || string(rule2.Condition[i]) == "#" {
			continue
		} else if string(rule1.Condition[i]) != string(rule2.Condition[i]) {
			conditionMatches = false
			break
		}
	}
	outputMatches := rule1.Output == rule2.Output
	return outputMatches && conditionMatches, nil
}

func (r Rule) String() string {
	return fmt.Sprintf("%v %v", r.Condition, r.Output)
}
