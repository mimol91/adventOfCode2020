package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	Amount    int
	Color     string
	InnerBags []Bag
}

var myRegexp = regexp.MustCompile(`([\d]) (.+?)bag`)
var rep = strings.NewReplacer(".", "", ",", "", "bags", "bag")

type Q []Bag

func (r *Q) Empty() bool   { return len(*r) == 0 }
func (r *Q) Enqueue(p Bag) { *r = append(*r, p) }
func (r *Q) Dequeue() Bag {
	if len(*r) == 0 {
		panic("opsie, Q is empty")
	}
	el := (*r)[0]
	*r = (*r)[1:]
	return el
}

func (b *Bag) Parse(line string) {
	elements := strings.Split(rep.Replace(line), "bag contain ")
	reaming := elements[1]
	if reaming == "no other bags" {
		return
	}

	b.Amount = 1
	b.Color = strings.TrimSpace(elements[0])
	inner := myRegexp.FindAllStringSubmatch(reaming, -1)
	for _, parts := range inner {
		b.InnerBags = append(b.InnerBags, Bag{
			Amount: atoi(parts[1]),
			Color:  strings.TrimSpace(parts[2]),
		})
	}
}

func main() {
	fmt.Printf("Part1 %d\n", part1())
	fmt.Printf("Part2 %d\n", part2())
}

func part1() int {
	var q Q
	names := make(map[string]struct{})
	lines := strings.Split(readInput(), "\n")
	bags := make(map[string]Bag)
	for _, line := range lines {
		bag := Bag{}
		bag.Parse(line)
		bags[bag.Color] = bag
	}
	traverse(bags, &q, "shiny gold", names)

	for !q.Empty() {
		el := q.Dequeue()
		traverse(bags, &q, el.Color, names)
	}

	return len(names)
}

func traverse(bags map[string]Bag, q *Q, name string, names map[string]struct{}) {
	for i, bag := range bags {
		for _, innerBag := range bag.InnerBags {
			if innerBag.Color == name {
				names[bag.Color] = struct{}{}
				q.Enqueue(bags[i])
			}
		}
	}

}

func part2() int {
	lines := strings.Split(readInput(), "\n")
	bags := make(map[string]Bag)
	for _, line := range lines {
		bag := Bag{}
		bag.Parse(line)
		bags[bag.Color] = bag
	}

	shinyGold := bags["shiny gold"]
	res := getNested(bags, shinyGold)
	return res
}
func getNested(bags map[string]Bag, bag Bag) int {
	if len(bag.InnerBags) == 0 {
		return 0
	}

	var sum int
	for _, innerBag := range bag.InnerBags {
		sum += innerBag.Amount
		sum += innerBag.Amount * getNested(bags, bags[innerBag.Color])
	}
	return bag.Amount * sum
}

func atoi(val string) int {
	if res, err := strconv.Atoi(val); err != nil {
		panic("unable to convert")
	} else {
		return res
	}
}

func readInput() string {
	b, err := ioutil.ReadFile("07/input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
