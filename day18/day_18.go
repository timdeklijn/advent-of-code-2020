package main

import (
	"bufio"
	"fmt"
	"os"
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
			// spl := strings.Split(s[i+1:], ")")
			// fmt.Println(spl[0])
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

func main() {
	part2 := false
	file, err := os.Open("data_18.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !part2 {
		n := homeworkPart1(scanner)
		fmt.Println("Day 18-1:", n)
	}
	// else {
	// 	n := simulateGrid4(scanner)
	// 	fmt.Println("Day 18-2:", n)
	// }
}
