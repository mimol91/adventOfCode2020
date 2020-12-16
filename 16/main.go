package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var pattern = regexp.MustCompile(`(.+): (\d+)-(\d+) or (\d+)-(\d+)`)
func main() {
	fmt.Printf("Part1 %d\n", part1())
	fmt.Printf("Part2 %d\n", part2())
}

type Range struct {
	min, max int
}

func (r Range) IsValid(val int) bool {
	return val >= r.min && val <= r.max
}

type TicketInfo struct {
	name string
	a    Range
	b    Range
}
type TicketsInfo []TicketInfo

func (r TicketsInfo) IsValid(val int) bool {
	for _, ticket := range r {
		if ticket.IsValid(val) == true {
			return true
		}
	}
	return false
}

type TicketsNr []int

func (r TicketInfo) IsValid(val int) bool {
	return r.a.IsValid(val) || r.b.IsValid(val)
}
func part1() int {
	var result int
	text := readInput()
	ticketsInfo, _, nearbyTickets := parse(text)
	for _, ticketsNr := range nearbyTickets {
		for _, nr := range ticketsNr {
			if ticketsInfo.IsValid(nr) == false {
				result += nr
			}
		}
	}

	return result
}

func part2() int {
	text := readInput()
	ticketsInfo, yourTickets, nearbyTickets := parse(text)
	validTickets := make([]TicketsNr, 0)
	for _, ticketsNr := range nearbyTickets {
		isValid := true
		for _, nr := range ticketsNr {
			if !ticketsInfo.IsValid(nr) {
				isValid = false
				break
			}
		}
		if isValid {
			validTickets = append(validTickets, ticketsNr)
		}
	}

	size := len(validTickets[0])
	cols := make([]TicketsNr, size)
	validTicketsSize := len(validTickets)
	for i := 0; i < size; i++ {
		cols[i] = make(TicketsNr, validTicketsSize)
		for j, ticket := range validTickets {
			cols[i][j] = ticket[i]
		}
	}

	validOptions := make(map[string][]int)

	for colNr, col := range cols {
		for _, ticketInfo := range ticketsInfo {
			isValid := true
			for _, nr := range col {
				if !ticketInfo.IsValid(nr) {
					isValid = false
					break
				}
			}
			if isValid {
				validOptions[ticketInfo.name] = append(validOptions[ticketInfo.name], colNr)
			}
		}
	}

	uniqueCols := make(map[int]struct{})
	for len(uniqueCols) != len(ticketsInfo) {
		for _, possibleColumns := range validOptions {
			if len(possibleColumns) == 1 {
				uniqueCols[possibleColumns[0]] = struct{}{}
			}
		}
		for key, possibleColumns := range validOptions {
			if len(possibleColumns) == 1 {
				continue
			}
			// @todo implement binary search as they are sorted
			for i := len(possibleColumns) - 1; i >= 0; i-- {
				for uniqueNr := range uniqueCols {
					if possibleColumns[i] == uniqueNr {
						validOptions[key] = append(validOptions[key][:i], validOptions[key][i+1:]...)
					}
				}
			}
		}
	}
	result := 1
	for key, col := range validOptions {
		if strings.HasPrefix(key, "departure") {
			result *= yourTickets[col[0]]
		}
	}
	return result
}

func parse(text string) (TicketsInfo, TicketsNr, []TicketsNr) {
	sections := strings.Split(text, "\n\n")

	ticketsInfoLines := strings.Split(sections[0], "\n")
	ticketsInfo := make(TicketsInfo, len(ticketsInfoLines))
	for i, line := range ticketsInfoLines {
		re := pattern.FindStringSubmatch(line)
		ticketInfo := TicketInfo{
			name: re[1],
			a:    Range{min: atoi(re[2]), max: atoi(re[3])},
			b:    Range{min: atoi(re[4]), max: atoi(re[5])},
		}
		ticketsInfo[i] = ticketInfo
	}

	yourTicketsLines := strings.Split(sections[1], "\n")[1]
	yourTickets := aatoi(strings.Split(yourTicketsLines, ","))

	nearbyTicketsLines := strings.Split(sections[2], "\n")[1:]
	nearbyTickets := make([]TicketsNr, len(nearbyTicketsLines))
	for i, line := range nearbyTicketsLines {
		nearbyTickets[i] = aatoi(strings.Split(line, ","))
	}

	return ticketsInfo, yourTickets, nearbyTickets
}

func aatoi(vals []string) []int {
	result := make([]int, len(vals))
	for i, val := range vals {
		result[i] = atoi(val)
	}
	return result
}

func atoi(val string) int {
	if res, err := strconv.Atoi(val); err != nil {
		panic("unable to convert")
	} else {
		return res
	}
}

func readInput() string {
	b, err := ioutil.ReadFile("16/input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
