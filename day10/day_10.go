package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// keep track of recursion results we have already seen
var branchesDone map[int]int

// readlines parses all lines in a scanner to ints and returns them as a slice.
func readlines(s *bufio.Scanner) []int {
	var lines []int

	// Read lines and convert to int, then append to slice
	for s.Scan() {
		row := strings.Trim(s.Text(), "\n")
		n, _ := strconv.Atoi(row)
		lines = append(lines, n)
	}

	if err := s.Err(); err != nil {
		panic(err)
	}
	return lines
}

func counter(start int, branches map[int]bool) int {

	// If we already did this one, return what we found then
	cnt, found := branchesDone[start]
	if found {
		return cnt
	}

	res := 0        // how many branches for this start
	nothing := true // Do we branch from here?

	for i := 1; i <= 3; i++ {
		// Check if we have follow up numbers with 3 or less difference
		nextStart := start + i
		_, ok := branches[nextStart]
		// If we do, branch from that one
		if ok {
			res += counter(nextStart, branches)
			nothing = false
		}
	}

	// If we did branch, add one, because that is the only option from this branch
	if nothing {
		res++
	}

	// Add result to branch start
	branchesDone[start] = res

	// return result
	return res
}

func adapterArangements(s *bufio.Scanner) int {
	// we are lucky that all numbers in the list are different. So we can use
	// that to keep track of wether we branched on that point or not.

	lines := readlines(s)                        // input
	sort.Ints(lines)                             // Sort ints
	lines = append([]int{0}, lines...)           // add 0 for power outlet
	lines = append(lines, lines[len(lines)-1]+3) // add 3 for output

	// keep track of what branch start are already done, accessable globaly
	branchesDone = make(map[int]int)

	// Create a map with branch starts
	branches := make(map[int]bool)
	for _, l := range lines {
		branches[l] = true
	}

	// start recursing over all branch starts
	return counter(0, branches)

}

func joltageDifference(s *bufio.Scanner) int {
	lines := readlines(s)                        // Read lines
	sort.Ints(lines)                             // Sort ints
	lines = append([]int{0}, lines...)           // add 0 for power outlet
	lines = append(lines, lines[len(lines)-1]+3) // add 3 for output

	// counters
	ones := 0
	threes := 0
	for i := 1; i < len(lines); i++ {
		d := lines[i] - lines[i-1]
		if d == 1 {
			ones++
		}
		if d == 3 {
			threes++
		}
	}
	return ones * threes
}

func main() {
	part2 := true
	file, err := os.Open("data_10.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !part2 {
		n := joltageDifference(scanner)
		fmt.Println("Day 10-1:", n)
	} else {
		n := adapterArangements(scanner)
		fmt.Println("Day 10-2:", n)
	}
}
