package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countCommonAnswers(s *bufio.Scanner) int {
	total := 0                  // total unique yes answers per group
	groupSize := 0              // Counter to check if the whole group answered yes
	group := make(map[byte]int) // group map

	for s.Scan() {

		row := strings.Trim(s.Text(), "\n")

		if len(row) == 0 {

			// If the number of 'yes' answers is equal to group size, we add it to the
			// total
			for _, v := range group {
				if v == groupSize {
					total++
				}
			}

			// Reset the counters
			group = make(map[byte]int)
			groupSize = 0
		} else {
			for i := 0; i < len(row); i++ {
				group[row[i]]++
			}
			groupSize++
		}
	}

	// UGLY: also do last group
	for _, v := range group {
		if v == groupSize {
			total++
		}
	}

	// Handle io error
	if err := s.Err(); err != nil {
		panic(err)
	}

	return total
}

func countUniqueAnswers(s *bufio.Scanner) int {
	total := 0                  // Count the number of 'yes'es
	group := make(map[byte]int) // Per answer count the number of yes

	for s.Scan() {

		row := strings.Trim(s.Text(), "\n")

		if len(row) == 0 {
			// If group is checked. Add the number of yes answers to the total
			total += len(group)
			group = make(map[byte]int)
		} else {
			for i := 0; i < len(row); i++ {
				group[row[i]]++
			}
		}
	}

	// UGLY: also do last group
	total += len(group)

	// Handle io error
	if err := s.Err(); err != nil {
		panic(err)
	}

	return total
}

func main() {
	fmt.Println("DAY 6")
	part2 := true
	file, err := os.Open("data_6.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !part2 {
		fmt.Println("Day 6-1:", countUniqueAnswers(scanner))
	} else {
		fmt.Println("Day 6-2:", countCommonAnswers(scanner))
	}
}
