package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Bag type contains the whole ruleset of bags. This works since the rules are only
// 1 deep.
type Bag map[string]map[string]int

func contains(bags Bag, b, target string) bool {
	// If target is in current bag: stop
	if _, ok := bags[b][target]; ok {
		return true
	}

	// If target is in one of the children: stop
	for b := range bags[b] {
		if contains(bags, b, target) {
			return true
		}
	}
	// We did not find the target
	return false
}

func count(bags Bag, b string) int {
	c := 0
	// Iterate over all children of bags[b]
	for name, num := range bags[b] {
		// - If num does not fit any bags, it still counts as 1
		// - We need to multiply the num with the number the children can hold:
		// 		if num = 2 and has a single child that can hold 2 children we have
		//		2 * 2 + 2
		// 		If the final children also hold 2 children:
		//		2 * (2 * 2 + 2) + 2 = 14
		//		etc....
		//		(see tests)
		c += num * (count(bags, name) + 1)
	}
	return c
}

func parseRow(row string, bags Bag) Bag {

	// Split on spaces, and create name for current bag
	spl := strings.Split(row, " ")
	name := spl[0] + "_" + spl[1] // first two words

	// handle no bags inside
	if spl[4] == "no" {
		bags[name] = map[string]int{}
		return bags
	}

	// If we have bags inside, create children and add to Bag[name] with their number
	children := strings.Split(row, "contain")
	// Children are comma separated, we trim it before separating.
	childrenList := strings.Split(strings.Trim(children[1], " "), ",")

	// Create map to start filling with children
	bags[name] = map[string]int{}
	for _, c := range childrenList {
		// Clean up the children
		c = strings.Trim(c, " ")
		s := strings.Split(c, " ")
		// Create child name and parse number. Place inside the parent map
		childName := s[1] + "_" + s[2]
		bags[name][childName], _ = strconv.Atoi(s[0])

	}
	return bags
}

func findParentBags(s *bufio.Scanner) int {

	bags := Bag{} // Create empty Bag

	for s.Scan() {
		// Trim and parse row
		row := strings.Trim(s.Text(), "\n")
		bags = parseRow(row, bags)
	}

	// Handle io error
	if err := s.Err(); err != nil {
		panic(err)
	}

	// For each bag in Bag we need to find if it can fit a "shiny_gold" bag. This bag
	// can have children, so something recursive is needed.
	counter := 0
	for b := range bags {
		if contains(bags, b, "shiny_gold") {
			counter++
		}
	}

	return counter
}

func countBagsInside(s *bufio.Scanner) int {
	// Parsing is identiacal to findParentBags, should have abstracted this:

	bags := Bag{} // Create empty Bag

	for s.Scan() {

		row := strings.Trim(s.Text(), "\n")
		bags = parseRow(row, bags)
	}

	// Handle io error
	if err := s.Err(); err != nil {
		panic(err)
	}

	// For "shiny_gold" we will loop over all children and the children's children
	// counting the numbers:
	return count(bags, "shiny_gold")
}

func main() {
	part2 := true
	file, err := os.Open("data_7.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !part2 {
		fmt.Println("Day 7-1:", findParentBags(scanner))
	} else {
		fmt.Println("Day 7-2:", countBagsInside(scanner))
	}
}
