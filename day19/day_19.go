package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var rules map[int]string // Save all rule strings in a map with integer identifyer
var ruleZero string      // regex for rule 0.

func replace(s string) string {
	n, _ := strconv.Atoi(s)
	return strings.Trim(rules[n], " ")
}

func tst(s string) bool {
	re := regexp.MustCompile(ruleZero)
	if re.MatchString(s) {
		return true
	}
	return false
}

func makeRegex(n int) string {
	rule := rules[n]
	rn := regexp.MustCompile(`([0-9])+`)
	for rn.MatchString(rule) {
		rule = rn.ReplaceAllStringFunc(rule, replace)
	}
	rule = strings.ReplaceAll(rule, string('"'), "")
	rule = "^" + strings.ReplaceAll(rule, string(" "), "") + "$"
	return rule
}

// =====================================================================================

// checkRules solves part 1
func checkRules(s *bufio.Scanner) int {

	readRules := true
	rules = make(map[int]string)
	sum := 0

	for s.Scan() {

		r := strings.Trim(s.Text(), "\n")

		if len(r) == 0 {
			ruleZero = makeRegex(0)
			readRules = false
			continue
		}

		if readRules {
			// Get the integer number
			spl := strings.Split(r, ":")
			n, _ := strconv.Atoi(spl[0])
			// Get the rul
			s := spl[1]
			// If we have an or, add parentheses
			if strings.Contains(s, "|") {
				s = "(" + s + ")"
			}
			rules[n] = s
		} else {
			if tst(r) {
				sum++
			}
		}

	}

	if err := s.Err(); err != nil {
		panic(err)
	}
	return sum
}

// =====================================================================================

func solve(l []string, change bool) int {

	// Create "ending" rule for 11
	var tmp []string
	for i := 1; i < 5; i++ {
		tmp = append(tmp, strings.Repeat("42 ", i)+strings.Repeat("31 ", i))
	}

	rule := strings.Join(tmp[:], "| ")

	if change {
		// New rules
		rules[8] = "42+" // will be just a bunch of 42's
		rules[11] = "(" + rule + ")"
	}

	// Create monster regex
	ruleZero = makeRegex(0)

	// Check which lines follow the rules
	sum := 0
	for _, s := range l {
		if tst(s) {
			sum++
		}
	}

	return sum
}

// checkNewRules solves part 2
func checkNewRules(s *bufio.Scanner, change bool) int {
	readRules := true
	rules = make(map[int]string)

	var lines []string

	for s.Scan() {

		r := strings.Trim(s.Text(), "\n")

		if len(r) == 0 {
			readRules = false
			continue
		}

		if readRules {
			// Get the integer number
			spl := strings.Split(r, ":")
			n, _ := strconv.Atoi(spl[0])
			// Get the rule
			s := spl[1]
			// If we have an or, add parentheses
			if strings.Contains(s, "|") {
				s = "(" + s + ")"
			}
			rules[n] = s
		} else {
			lines = append(lines, r)
		}

	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	return solve(lines, change)
}

// =====================================================================================

func main() {
	part2 := true
	file, err := os.Open("data_19.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !part2 {
		n := checkRules(scanner)
		fmt.Println("Day 19-1:", n)
	} else {
		n := checkNewRules(scanner, true)
		fmt.Println("Day 19-2:", n)
	}
}
