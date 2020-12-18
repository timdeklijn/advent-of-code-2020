package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// doOp either sums or adds n1 and n2
func doOp(n1, n2 int, op string) int {
	switch op {
	case "*":
		return n1 * n2
	default:
		return n1 + n2
	}
}

// parSplit extracts a string from the outer most left parentheses.
func parSplit(s string) string {
	n := 0
	ns := ""
	for _, c := range s {
		cc := string(c)
		if cc == ")" && n == 0 {
			return ns
		}

		if cc == ")" && n > 0 {
			ns += cc
			n--
			continue
		}

		if cc == "(" {
			ns += cc
			n++
			continue
		}

		ns += cc

	}

	return ns
}

// calc is a recursive function that walks over a string and returns the answer of the
// sum in that string.
func calc(s string, tot int) int {
	op := "+"
	i := 0
	for i < len(s) {

		// convert bytes to string
		ss := string(s[i])

		// If we encounter a number:
		if n, err := strconv.Atoi(ss); err == nil {
			tot = doOp(tot, n, op)
		}

		// Check for ops:
		if ss == "+" || ss == "*" {
			op = ss
		}

		// handle parentheses
		if ss == "(" {
			ns := parSplit(s[i+1:])
			// Handle stuff inside parentheses
			tmp := calc(ns, 0)
			// Increment i, keep parentheses in mind
			i += len(ns) + 2
			// Do operation with result of that
			tot = doOp(tot, tmp, op)
			continue
		}

		i++

	}
	return tot
}

func advancedCalc(s string) string {

	fields := strings.Fields(strings.Trim(s, "()"))
	tot, _ := strconv.Atoi(fields[0])

	for i := 1; i < len(fields); i += 2 {
		switch n, _ := strconv.Atoi(fields[i+1]); fields[i] {
		case "+":
			tot += n
		case "*":
			tot *= n
		}
	}

	return strconv.Itoa(tot)
}

func noParam(s string) string {
	rr := regexp.MustCompile(`\d+ \+ \d+`)
	for rr.MatchString(s) {
		s = rr.ReplaceAllStringFunc(s, advancedCalc)
	}
	s = advancedCalc(s)
	return s
}

func calcAdvanced(s string) int {

	re := regexp.MustCompile(`\([^\(\)]+\)`)

	for re.MatchString(s) {
		s = re.ReplaceAllStringFunc(s, noParam)
	}

	ret, _ := strconv.Atoi(noParam(s))

	return ret
}

func homeworkPart1(s *bufio.Scanner) int {
	tot := 0
	for s.Scan() {
		// Prep strings
		r := strings.Trim(s.Text(), "\n")
		r = strings.ReplaceAll(r, " ", "")
		// Do calculation
		tot += calc(r, 0)
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return tot
}

func homeworkPart2(s *bufio.Scanner) int {
	tot := 0
	for s.Scan() {
		// Prep strings
		r := strings.Trim(s.Text(), "\n")
		r = strings.TrimSpace(r)
		// Do calculation
		tot += calcAdvanced(r)
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return tot
}

func main() {
	part2 := true
	file, err := os.Open("data_18.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !part2 {
		n := homeworkPart1(scanner)
		fmt.Println("Day 18-1:", n)
	} else {
		n := homeworkPart2(scanner)
		fmt.Println("Day 18-2:", n)
	}
}
