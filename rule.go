package ga

import "fmt"

type Rule struct {
	condition []int
	output    int
}

func (rule1 Rule) Matches(rule2 Rule) bool {
	if len(rule1.condition) != len(rule2.condition) {
		return false
	}
	conditionMatches := true
	for i := range rule1.condition {
		if rule1.condition[i] == 2 || rule2.condition[i] == 2 {
			continue
		} else if rule1.condition[i] != rule2.condition[i] {
			conditionMatches = false
			break
		}
	}
	outputMatches := rule1.output == rule2.output
	return outputMatches && conditionMatches
}

func (r Rule) String() string {
	return fmt.Sprintf("%v %v", r.condition, r.output)
}
