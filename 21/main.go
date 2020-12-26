package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Printf("Part1 %d\n", part1())
	//fmt.Printf("Part2 %d\n", part2())
}

type Food struct {
	ingredients []string
	allergens   []string
}
type FoodList []Food
type AllergensMap map[string][]string

func (m AllergensMap) solve() {
	unique := make(map[string]struct{})
	for len(unique) != len(m) {
		var toRemove []string
		for allergen, ingredients := range m {
			if len(ingredients) == 1 {
				if _, ok := unique[allergen]; !ok {
					toRemove = append(toRemove, ingredients[0])
					unique[allergen] = struct{}{}
				}
			}
		}

		for allergen, ingredients := range m {
			if len(ingredients) == 1 {
				continue
			}
			m[allergen] = remove(ingredients, toRemove)
		}
	}

}

func (r FoodList) AsMap() AllergensMap {
	result := make(map[string][]string)
	for _, food := range r {
		for _, allergen := range food.allergens {
			result[allergen] = intersection(result[allergen], food.ingredients)
		}
	}
	return result
}

func (r FoodList) CountNoAllergens(allergensMap AllergensMap) int {
	var result int

	allergens := make(map[string]struct{}, len(allergensMap))
	for _, allergen := range allergensMap {
		allergens[allergen[0]] = struct{}{}
	}

	for _, food := range r {
		for _, ingredient := range food.ingredients {
			if _, ok := allergens[ingredient]; !ok {
				result++
			}
		}
	}
	return result
}

func part1() int {
	text := readInput()
	foodList := parse(text)
	allergensMap := foodList.AsMap()
	allergensMap.solve()

	return foodList.CountNoAllergens(allergensMap)
}

func parse(text string) FoodList {
	text = strings.Replace(text, "(", "", -1)
	text = strings.Replace(text, ")", "", -1)
	lines := strings.Split(text, "\n")
	result := make(FoodList, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " contains ")
		result[i] = Food{
			ingredients: strings.Split(parts[0], " "),
			allergens:   strings.Split(parts[1], ", "),
		}
	}
	return result
}

func part2() int {

	return 0
}

func intersection(a, b []string) []string {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	var result []string
	lookUp := make(map[string]struct{}, len(a))
	for _, el := range a {
		lookUp[el] = struct{}{}
	}
	for _, el := range b {
		if _, ok := lookUp[el]; ok {
			result = append(result, el)
		}
	}
	return result
}
func remove(list []string, itemsToRemove []string) []string {
	var result []string
	lookUp := make(map[string]struct{}, len(itemsToRemove))
	for _, item := range itemsToRemove {
		lookUp[item] = struct{}{}
	}
	for _, el := range list {
		if _, ok := lookUp[el]; !ok {
			result = append(result, el)
		}
	}
	return result
}

func readInput() string {
	b, err := ioutil.ReadFile("21/input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
