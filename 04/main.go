package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	BirthDayYear   int
	IssueYear      int
	ExpirationYear int
	Height         string
	HairColor      string
	EyeColor       string
	PassportID     string
	CountryID      string
}

func (p *Passport) Set(key, val string) {
	switch key {
	case "byr":
		p.BirthDayYear = p.toInt(val)
	case "iyr":
		p.IssueYear = p.toInt(val)
	case "eyr":
		p.ExpirationYear = p.toInt(val)
	case "hgt":
		p.Height = val
	case "hcl":
		p.HairColor = val
	case "ecl":
		p.EyeColor = val
	case "cid":
		p.CountryID = val
	case "pid":
		p.PassportID = val
	}
}

func (Passport) toInt(val string) int {
	num, err := strconv.Atoi(val)
	if err != nil {
		panic(val)
	}
	return num
}

func (p Passport) IsValid() bool {
	return p.BirthDayYear != 0 &&
		p.IssueYear != 0 &&
		p.ExpirationYear != 0 &&
		p.Height != "" &&
		p.HairColor != "" &&
		p.EyeColor != "" &&
		p.PassportID != ""
}

func (p Passport) IsStrictValid() bool {
	if !p.IsValid() {
		return false
	}
	if !p.hasValidEyeColor() {
		return false
	}
	if !p.inRange(p.BirthDayYear, 1920, 2002) {
		return false
	}
	if !p.inRange(p.IssueYear, 2010, 2020) {
		return false
	}
	if !p.inRange(p.ExpirationYear, 2020, 2030) {
		return false
	}
	if !p.hasValidHeight() {
		return false
	}
	if !p.hasValidHairColor() {
		return false
	}
	if !p.hasValidPassportID() {
		return false
	}

	return true
}

func (p Passport) hasValidHeight() bool {
	var val int
	var metric string
	if _, err := fmt.Sscanf(p.Height, "%d%s", &val, &metric); err != nil {
		return false
	}
	switch metric {
	case "cm":
		return p.inRange(val, 150, 193)
	case "in":
		return p.inRange(val, 59, 96)
	}
	return false
}

func (p Passport) hasValidHairColor() bool {
	pattern := regexp.MustCompile("^#[a-fA-F0-9]{6}$")
	if !pattern.Match([]byte(p.HairColor)) {
		return false
	}
	return true
}

func (p *Passport) hasValidEyeColor() bool {
	validColors := map[string]struct{}{
		"amb": {},
		"blu": {},
		"brn": {},
		"gry": {},
		"grn": {},
		"hzl": {},
		"oth": {},
	}
	_, ok := validColors[p.EyeColor]
	return ok
}

func (p Passport) hasValidPassportID() bool {
	if len(p.PassportID) != 9 {
		return false
	}
	if _, err := strconv.Atoi(p.PassportID); err != nil {
		return false
	}
	return true
}

func (Passport) inRange(val, min, max int) bool {
	return val >= min && val <= max
}

type Passports []Passport

func (p Passports) CountValid(fn func(passport Passport) bool) int {
	var result int
	for _, passport := range p {
		if fn(passport) {
			result++
		}
	}

	return result
}
func main() {
	fmt.Printf("Part1 %d\n", part1())
	fmt.Printf("Part2 %d\n", part2())

}

func part1() int {
	text := readInput()
	elements := strings.Split(text, "\n\n")
	passports := make(Passports, len(elements))

	for i, element := range elements {
		var passport Passport
		val := strings.ReplaceAll(element, "\n", " ")
		chunks := strings.Split(val, " ")
		for _, chunk := range chunks {
			parts := strings.Split(chunk, ":")
			if len(parts) != 2 {
				panic("invalid data")
			}
			passport.Set(parts[0], parts[1])

		}
		passports[i] = passport
	}

	return passports.CountValid(Passport.IsValid)
}
func part2() int {
	text := readInput()
	elements := strings.Split(text, "\n\n")
	passports := make(Passports, len(elements))

	for i, element := range elements {
		var passport Passport
		val := strings.ReplaceAll(element, "\n", " ")
		chunks := strings.Split(val, " ")
		for _, chunk := range chunks {
			parts := strings.Split(chunk, ":")
			if len(parts) != 2 {
				panic("invalid data")
			}
			passport.Set(parts[0], parts[1])

		}
		passports[i] = passport
	}

	return passports.CountValid(Passport.IsStrictValid)
}

func readInput() string {
	b, err := ioutil.ReadFile("04/input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
