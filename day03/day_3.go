package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type slope struct {
	right int
	down  int
	pos   int
	count int
}

// TreesInRouteProduct returns the product of number of trees found in multiple routes
func TreesInRouteProduct(s *bufio.Scanner) int {

	var row string
	var ll int // length of row

	// Counters
	first := true   // Skip counting trees in first row
	rowCounter := 0 // keep track of row number

	// Init slopes and their counters
	slopes := []slope{
		slope{1, 1, 0, 0},
		slope{3, 1, 0, 0},
		slope{5, 1, 0, 0},
		slope{7, 1, 0, 0},
		slope{1, 2, 0, 0},
	}

	// Scan the file and handle line by line
	for s.Scan() {
		// We don't check trees in the first line but we need to know the length
		row = strings.Trim(s.Text(), "\n")
		if first {
			ll = len(row)
			first = false
			rowCounter++
			continue
		}

		// Handle each line for all slopes
		for i := range slopes {
			// Check if the line should be checks, i.e. down step is 1 or 2
			if rowCounter%slopes[i].down == 0 {
				// Step to the right
				slopes[i].pos += slopes[i].right
				// Check if there is a tree at the position
				if row[slopes[i].pos%ll] == '#' {
					slopes[i].count++
				}
			}
		}
		rowCounter++
	}
	// Handle scanner error
	if err := s.Err(); err != nil {
		panic(err)
	}

	// Take the product of the counts
	prod := 1
	for i := range slopes {
		prod *= slopes[i].count
	}
	return prod
}

// TreesInRoute counts trees for a single slope
func TreesInRoute(s *bufio.Scanner) int {
	n := 0         // horizontal position
	first := true  // do not count the first row
	treeCount := 0 // counter
	var row string
	var ll int // lenght of row

	// Scan the file and handle line by line
	for s.Scan() {
		// We don't check trees in the first line but we need to know the length
		row = strings.Trim(s.Text(), "\n")
		if first {
			ll = len(row)
			first = false
			continue
		}
		// Step and count
		n += 3
		if row[n%ll] == '#' {
			treeCount++
		}
	}

	// Handle scanner error
	if err := s.Err(); err != nil {
		panic(err)
	}
	return treeCount
}

func main() {
	part2 := true
	file, err := os.Open("data_3.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	if !part2 {
		fmt.Println("Day 3-1:", TreesInRoute(scanner))
	} else {
		fmt.Println("Day 3-1:", TreesInRouteProduct(scanner))
	}
}
