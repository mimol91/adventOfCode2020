package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Rules struct {
	Rules       []Rule
	SearchMap   SearchMap
	SolvedRules map[int]struct{}
}
type SearchMap map[int]map[int]struct{} //Key Rule=> Val list of other rules depends on
type Element struct {
	RulesNr []int
	Solved  bool
	Value   string
}
type Elements []Element

func (e Elements) Values() []string {
	result := make([]string, len(e))
	for i, el := range e {
		result[i] = el.Value
	}
	return result
}

type Rule struct {
	Elements Elements
}

func (r Rule) IsSolved() bool {
	for _, el := range r.Elements {
		if el.Solved == false {
			return false
		}
	}
	return true
}

func (r *Rules) Solve() {
	for len(r.SolvedRules) != len(r.Rules) {

		for ruleNr := range r.SolvedRules {
			dependRulesNr := r.SearchMap[ruleNr]
			if len(dependRulesNr) == 0 {
				continue
			}

			for dependedRuleNr := range dependRulesNr {
				//rule := r.Rules[dependedRuleNr]
				if r.Rules[dependedRuleNr].IsSolved() {
					continue
				}

				for i := len(r.Rules[dependedRuleNr].Elements) - 1; i >= 0; i-- {
					el := r.Rules[dependedRuleNr].Elements[i]
					if el.Solved {
						continue
					}
					firstRuleNr := el.RulesNr[0]
					if firstRuleNr != ruleNr {
						continue
					}
					rulesLen := len(el.RulesNr)
					solvedRule := r.Rules[firstRuleNr]

					//Remove processing element
					r.Rules[dependedRuleNr].Elements = append(r.Rules[dependedRuleNr].Elements[:i], r.Rules[dependedRuleNr].Elements[i+1:]...)
					i++
					for _, solvedEl := range solvedRule.Elements {
						element := Element{
							Solved: rulesLen == 1,
							Value:  el.Value + solvedEl.Value,
						}
						if !element.Solved {
							element.RulesNr = el.RulesNr[1:]
						}
						i--
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
	fmt.Printf("Part2 %d\n", part2())
}

func part1() int {
	var result int
	text := readInput()
	rules, messages := parse(text)
	rules.Solve()
	possibleVals := make(map[string]struct{}, len(rules.Rules[0].Elements))

	for _, elements := range rules.Rules[0].Elements {
		possibleVals[elements.Value] = struct{}{}
	}

	for _, message := range messages {
		if _, ok := possibleVals[message]; ok {
			result++
		}
	}

	return result
}

func part2() int {
	var result int
	text := readInput()
	rules, messages := parse(text)
	rules.Solve()

	part31 := fmt.Sprintf("(%s)", strings.Join(rules.Rules[31].Elements.Values(), "|"))
	part42 := fmt.Sprintf("(%s)", strings.Join(rules.Rules[42].Elements.Values(), "|"))

	// Rule 8 -> one or more times rule 42
	rule8String := fmt.Sprintf("(%s)+", part42)


	makeRegexp := func(num int) *regexp.Regexp {
		// rule 11:  equal number of 42 and 31 rules
		pattern := fmt.Sprintf("^%s%s{%d}%s{%d}$", rule8String, part42, num, part31, num)
		return regexp.MustCompile(pattern)
	}

	for _, m := range messages {
		for i := 1; i < 10; i++ {
			pattern := makeRegexp(i)
			if pattern.MatchString(m) {
				result++
				break
			}
		}
	}

	return result
}

func parse(text string) (Rules, []string) {
	parts := strings.Split(text, "\n\n")
	rules := parseRules(parts[0])
	messages := strings.Split(parts[1], "\n")

	return rules, messages
}
func parseRules(text string) Rules {
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
					Value:  parts[1][2:3],
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
