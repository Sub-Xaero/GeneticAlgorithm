package ga

import (
	"fmt"
	"errors"
)

type Rule struct {
	condition Bitstring
	output    string
}

func (rule1 Rule) Matches(rule2 Rule) (bool, error) {
	if len(rule1.condition) != len(rule2.condition) {
		return false, errors.New("conditions are not same length")
	}
	conditionMatches := true
	for i := range rule1.condition {
		if string(rule1.condition[i]) == "#" || string(rule2.condition[i]) == "#" {
			continue
		} else if string(rule1.condition[i]) != string(rule2.condition[i]) {
			conditionMatches = false
			break
		}
	}
	outputMatches := rule1.output == rule2.output
	return outputMatches && conditionMatches, nil
}

func (r Rule) String() string {
	return fmt.Sprintf("%v %v", r.condition, r.output)
}
