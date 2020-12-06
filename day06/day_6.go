package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countCommonAnswers(s *bufio.Scanner) int {
	total := 0
	groupSize := 0
	group := make(map[byte]int)
	for s.Scan() {
		row := strings.Trim(s.Text(), "\n")
		if len(row) == 0 {
			for _, v := range group {
				if v == groupSize {
					total++
				}
			}
			group = make(map[byte]int)
			groupSize = 0
		} else {
			for i := 0; i < len(row); i++ {
				group[row[i]]++
			}
			groupSize++
		}
	}

	// TODO: HOW COME I HAVE TO DO THIS?
	for _, v := range group {
		if v == groupSize {
			total++
		}
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return total
}

func countUniqueAnswers(s *bufio.Scanner) int {
	total := 0
	group := make(map[byte]int)
	for s.Scan() {
		row := strings.Trim(s.Text(), "\n")
		if len(row) == 0 {
			total += len(group)
			group = make(map[byte]int)
		} else {
			for i := 0; i < len(row); i++ {
				group[row[i]]++
			}
		}
	}
	total += len(group)
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
