package main

import (
	"fmt"
	"strconv"
	"strings"
)

// parseInput converts a stirng of comma separated ints to a list of ints
func parseInput(s string) []int {
	var ret []int
	for _, n := range strings.Split(s, ",") {
		nn, _ := strconv.Atoi(n)
		ret = append(ret, nn)
	}
	return ret
}

func playGame(input []int, it int) int {
	answer := make(map[int]int) // save number + the turn it was asnwered
	var turn, last int          // turn number and last said number. last is what we will return

	// Fill answer with numbers in list, excluding the last one
	for _, n := range input {
		if turn > 0 {
			answer[last] = turn
		}
		last = n
		turn++
	}

	// Run through all turns
	for turn < it {

		// If we already had last as an answer
		if t, ok := answer[last]; ok {

			answer[last] = turn // update turn beloning to last
			last = turn - t     // update answer
		} else {
			// We did not see last untill now

			answer[last] = turn // add last to seen answers
			last = 0            // last is 0, because previous answer was unique
		}
		turn++
	}
	return last
}

func main() {
	part2 := true
	input := "17,1,3,16,19,0"
	if !part2 {
		n := playGame(parseInput(input), 2020)
		fmt.Println("Day 15-1:", n)
	} else {
		n := playGame(parseInput(input), 30000000)
		fmt.Println("Day 15-2:", n)
	}
}
