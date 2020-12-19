package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type SearchMap map[int]map[int]struct{} //Key Rule=> Val list of other rules depends on
type Rules struct {
	Rules       []Rule
	SearchMap   SearchMap
	SolvedRules map[int]struct{}
}

type Element struct {
	RulesNr []int
	Solved  bool
	Value   string
}
type Rule struct {
	Elements []Element
}

func (r Rule) IsSolved() bool {
	for _, el := range r.Elements {
		if el.Solved == false {
			return false
		}
	}
	return true
}

func (r Rules) Print() {
	var sb strings.Builder
	for id, rule := range r.Rules {
		sb.WriteString(fmt.Sprintf("%d:", id))

		rulesText := make([]string, 0)
		for _, e := range rule.Elements {

			rulesText = append(rulesText, fmt.Sprintf("%s %d", e.Value, e.RulesNr))
		}

		sb.WriteString(strings.Join(rulesText, " | "))
		sb.WriteString("\n")
	}
	sb.WriteString("\n")
	fmt.Println(sb.String())
}

func (r *Rules) Solve() {
	var i int
	for len(r.SolvedRules) != len(r.Rules) {
		i++
		fmt.Println(fmt.Sprintf("-- %d --", i))
		r.Print()
		for ruleNr := range r.SolvedRules {
			dependRulesNr := r.SearchMap[ruleNr]
			if len(dependRulesNr) == 0 {
				continue
			}

			for dependedRuleNr := range dependRulesNr {
				rule := r.Rules[dependedRuleNr]
				for _, el := range rule.Elements {
					firstRuleNr := el.RulesNr[0]
					if firstRuleNr != ruleNr {
						continue
					}

					rulesLen := len(el.RulesNr)
					solvedRule := r.Rules[firstRuleNr]
					r.Rules[dependedRuleNr].Elements = r.Rules[dependedRuleNr].Elements[1:] //remove 1st
					for _, solvedEl := range solvedRule.Elements {
						element := Element{
							Solved: rulesLen == 1,
							Value:  el.Value + solvedEl.Value,
						}
						if !element.Solved {
							element.RulesNr = el.RulesNr[1:]
						}
						r.Rules[dependedRuleNr].Elements = append(r.Rules[dependedRuleNr].Elements, element)
					}
					if r.Rules[dependedRuleNr].IsSolved() {
						r.SolvedRules[dependedRuleNr] = struct{}{}
					}
				}
			}
		}
	}
}

func main() {
	fmt.Printf("Part1 %d\n", part1())
	//fmt.Printf("Part2 %d\n", part2())
}

func part1() int {
	text := readInput()
	rules := parse(text)

	rules.Solve()
	return 0
}

func part2() int {

	return 0
}

func parse(text string) Rules {
	rules := Rules{
		SearchMap:   make(SearchMap),
		SolvedRules: make(map[int]struct{}),
	}
	lines := strings.Split(text, "\n")
	result := make([]Rule, len(lines))

	for _, line := range lines {
		parts := strings.Split(line, ":")
		ruleID := atoi(parts[0])

		if strings.HasPrefix(parts[1], ` "`) {
			//parts[1] = ` "a"`
			result[ruleID] = Rule{
				Elements: []Element{{
					Solved: true,
					Value:  parts[1][2:3], //@todo
				}},
			}
			rules.SolvedRules[ruleID] = struct{}{}
		} else {
			el := strings.Split(parts[1], "|")
			elements := make([]Element, 0)

			for _, text := range el {
				numbers := strings.Split(strings.TrimSpace(text), " ")

				rulesNr := make([]int, len(numbers))
				for j, nr := range numbers {
					val := atoi(nr)
					rulesNr[j] = val

					// Add to searchMap
					if _, ok := rules.SearchMap[val]; !ok {
						rules.SearchMap[val] = make(map[int]struct{}, 0)
					}
					rules.SearchMap[val][ruleID] = struct{}{}
				}
				elements = append(elements, Element{RulesNr: rulesNr})
			}
			result[ruleID] = Rule{Elements: elements}
		}
	}

	rules.Rules = result
	return rules
}
func atoi(val string) int {
	if res, err := strconv.Atoi(val); err != nil {
		panic("unable to convert")
	} else {
		return res
	}
}
func readInput() string {
	b, err := ioutil.ReadFile("19/input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
