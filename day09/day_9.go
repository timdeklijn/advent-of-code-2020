package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// readlines parses all lines in a scanner to ints and returns them as a slice.
func readlines(s *bufio.Scanner) []int {
	var lines []int
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

// findMinMax returns the min and max numbers of a slice of ints
func findMinMax(l []int) (int, int) {
	min := 1000000000000
	max := 0
	for _, i := range l {
		if i < min {
			min = i
			continue
		}
		if i > max {
			max = i
			continue
		}
	}
	return min, max
}

// encryptionWeakness is the answer to part 2
func encryptionWeakness(s *bufio.Scanner, preamble, target int) (int, error) {
	lines := readlines(s) // inputs

	// Move the start of our search up
	for start := 0; start < len(lines); start++ {

		sum := 0     // needs to become the target
		var ll []int // save numbers used in

		// Sum numbers from start until the sum equals the target or overshoots it.
		for i := start; i < len(lines); i++ {
			sum += lines[i]
			ll = append(ll, lines[i])

			// We found it
			if sum == target {
				// find the min and max from the list of numbers used to sum and return
				// min + max
				min, max := findMinMax(ll)
				return min + max, nil
			}

			// Sum is larger then target, move start and try again.
			if sum > target {
				break
			}
		}
	}

	// Nothing found return error
	return 0, fmt.Errorf("No sum found")
}

// findWeakness finds the answer to part 1
func findWeakness(s *bufio.Scanner, preamble int) (int, error) {
	lines := readlines(s) // inputs
	start := 0            // start of preamble

	// Loop over all data
	for i := preamble + 1; i < len(lines); i++ {
		ok := false // Did we find a sum or not?

		// Double loop over preamble to find the sum
		for j := start; j < i; j++ {
			for k := start; k < i; k++ {

				// Cannot be the same number
				if j == k || lines[j] == lines[k] {
					continue
				}

				// If they sum up to the number after the preamble we found it
				if lines[j]+lines[k] == lines[i] {
					ok = true
					break
				}
			}

			// We found a sum, so no weakness
			if ok {
				break
			}
		}

		// If we did not find a sum for lines[i], this is the weekness
		if !ok {
			return lines[i], nil
		}

		// Preamble moves over by one
		start++
	}
	return 0, fmt.Errorf("No weakness found")
}

func main() {
	part2 := true
	file, err := os.Open("data_9.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !part2 {
		n, _ := findWeakness(scanner, 25)
		fmt.Println("Day 9-1:", n)
	} else {
		n, _ := encryptionWeakness(scanner, 25, 3199139634)
		fmt.Println("Day 9-2:", n)
	}
}
