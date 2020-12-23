package main

import (
	"fmt"
	"strconv"
	"strings"
)

// maxVal returns the largest value from a list of ints
func maxVal(l []int) int {
	m := 0
	for _, n := range l {
		if n > m {
			m = n
		}
	}
	return m
}

// CupGame keeps track of the turn we are in, which is related to the target cup as
// well as the list of cups.
type CupGame struct {
	turn int
	cups []int
}

// slice splits the cup list into a pickup list and a tmp cup list
func (c *CupGame) slice(s int) ([]int, []int) {
	var pickUp []int
	var tmpCups []int

	// Keep track of indices picked up
	selector := make(map[int]bool)

	// generate pickup indixes
	for i := s + 1; i < s+4; i++ {
		pickUp = append(pickUp, c.cups[i%len(c.cups)])
		selector[i%len(c.cups)] = true
	}

	// Create a list ignoring picked up indices
	for i := range c.cups {
		if _, ok := selector[i]; !ok {
			tmpCups = append(tmpCups, c.cups[i])
		}
	}

	return pickUp, tmpCups
}

// rotate puts the target cup back to its original position
func (c *CupGame) rotate(n int, newCups []int) {
	newIndex := 0
	for i, nn := range newCups {
		if nn == c.cups[n] {
			newIndex = i
		}
	}

	// We will shift by this much
	shift := newIndex - n

	var rotatedCups []int
	for i := range newCups {
		rotatedCups = append(rotatedCups, newCups[(i+shift)%len(newCups)])
	}

	c.cups = rotatedCups
}

// move does a single move to the cup list
func (c *CupGame) move() {
	n := c.turn % len(c.cups)

	pickUp, tmpCups := c.slice(n)                                   // take out three cups
	destinationCupIndex := selectDestinationCup(tmpCups, c.cups[n]) // find destination cup
	end := append(pickUp, tmpCups[destinationCupIndex+1:]...)       // construct end of new cup list
	newCups := append(tmpCups[:destinationCupIndex+1], end...)      // add end to beginning of new cup list

	c.rotate(n, newCups) // rotate the cup list to have the target at its original position
	c.turn++
}

// selectDestinationCup finds the cup to place the pickup cups behind
func selectDestinationCup(l []int, target int) int {
	nt := target
	for {
		nt--

		// we should not choose the target
		if nt == target {
			nt--
		}
		// we wrap around
		if nt == -1 {
			nt = maxVal(l)
		}

		// we find the cup we are looking for
		for i, n := range l {
			if n == nt {
				return i
			}
		}
	}
}

// newCupGame creates a new CupGame struct at turn 0
func newCupGame(s string) CupGame {
	var cups []int
	for _, s := range strings.Split(s, "") {
		n, _ := strconv.Atoi(s)
		cups = append(cups, n)
	}
	return CupGame{0, cups}
}

// createAnswer converts a list of ints to a single big number with all numbers after
// "1" in the list (wrap around)
func createAnswer(l []int) int {
	// find at which index 1 is
	oneIndex := 0
	for i := range l {
		if l[i] == 1 {
			oneIndex = i
			break
		}
	}
	// collect all numbers after one into a single string
	s := ""
	for i := oneIndex + 1; i < len(l)+oneIndex; i++ {
		s += strconv.Itoa(l[i%len(l)])
	}
	// convert back to a number and retunr
	n, _ := strconv.Atoi(s)
	return n
}

// crabCupspt1 calculates the answer to part 1
func crabCupspt1(input string, turns int) int {
	cups := newCupGame(input)
	for i := 0; i < turns; i++ {
		cups.move()
	}
	return createAnswer(cups.cups)
}

// newCupGame creates a new CupGame struct at turn 0
func newBigCupGame(s string) CupGame {
	var cups []int
	for _, s := range strings.Split(s, "") {
		n, _ := strconv.Atoi(s)
		cups = append(cups, n)
	}
	mv := maxVal(cups)
	for i := mv + 1; i < 1000001; i++ {
		cups = append(cups, i)
	}
	return CupGame{0, cups}
}

func createBigAnswer(l []int) int {
	oneIndex := 0
	for i := range l {
		if l[i] == 1 {
			oneIndex = i
		}
	}
	return l[oneIndex+1] * l[oneIndex+2]
}

func crabCupspt2(input string, turns int) int {
	cups := newBigCupGame(input)
	for i := 0; i < turns; i++ {
		if i%1000 == 0 {
			fmt.Println("Turn", i)
		}
		cups.move()
	}
	return createBigAnswer(cups.cups)
}

func main() {
	part2 := false
	if !part2 {
		n := crabCupspt1("792845136", 100)
		fmt.Println("Day 22-1:", n)
	} else {
		n := crabCupspt1("792845136", 10000000)
		fmt.Println("Day 22-2:", n)
	}
}
