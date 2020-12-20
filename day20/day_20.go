package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// reverseString reverses the order of the characters in a sting
func reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

type Tile struct {
	id      int
	raw     []string
	borders map[string]bool
}

// borderStrings collects all strings of the borders of the tile and adds them to
// the borders map being set to false
func (t *Tile) borderStrings() {
	// top
	topString := t.raw[0]
	t.borders[topString] = false
	t.borders[reverseString(topString)] = false

	// bottom
	bottomString := t.raw[len(t.raw)-1]
	t.borders[bottomString] = false
	t.borders[reverseString(bottomString)] = false

	// extract left and right side
	leftString := ""
	rightString := ""
	for _, s := range t.raw {
		leftString += string(s[0])
		rightString += string(s[len(s)-1])
	}

	// left
	t.borders[leftString] = false
	t.borders[reverseString(leftString)] = false

	// right
	t.borders[rightString] = false
	t.borders[reverseString(rightString)] = false
}

// parseInput scans a file line by line and creates Tile object appending them to a
// list. This list is returned.
func parseInput(s *bufio.Scanner) []Tile {
	var tileList []Tile                // final tile list
	var rawTile []string               // save tile strings in
	var id int                         // id of tile
	re := regexp.MustCompile(`[0-9]+`) // regex to get tile id

	for s.Scan() {

		// Get ID for tile
		r := strings.Trim(s.Text(), "\n")
		if re.MatchString(r) {
			s := re.FindString(r)
			id, _ = strconv.Atoi(s)
			continue
		}

		// We have an empty line, so we've seen a full tile. Add it to our list
		if len(r) == 0 {
			m := make(map[string]bool)
			tileList = append(tileList, Tile{id, rawTile, m})
			rawTile = []string{}
			continue
		}

		// Just add raw tile string to tileList
		rawTile = append(rawTile, r)
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	return tileList
}

func arrangeTiles(s *bufio.Scanner) int {
	// Read input
	tileList := parseInput(s)

	// For each side, extract forward and reverse as a string
	for _, t := range tileList {
		fmt.Println(t)
		t.borderStrings()
	}

	// For each tile check if one of the borderstrings is present in one of the other
	// borderstrings
	for i := range tileList {
		for j := range tileList {
			// do not compare with itself
			if i != j {
				// compare keys, the border strings
				for k := range tileList[j].borders {
					if _, ok := tileList[i].borders[k]; ok {
						// set this string to true
						tileList[i].borders[k] = true
						tileList[j].borders[k] = true
					}
				}
			}
		}
	}

	// Count number of strings that have been seen per tile.
	var ids []int
	for _, t := range tileList {
		s := 0
		for _, v := range t.borders {
			if v {
				s++
			}
		}
		// If we found 4 borderstrings set to true, this is a corner tile
		if s == 4 {
			ids = append(ids, t.id)
		}
	}

	// Calculate the product of the corner tile id's
	prod := 1
	for _, id := range ids {
		prod *= id
	}

	return prod
}

func main() {
	part2 := false
	file, err := os.Open("data_20.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !part2 {
		n := arrangeTiles(scanner)
		fmt.Println("Day 20-1:", n)
	}
	// else {
	//n := checkNewRules(scanner, true)
	//fmt.Println("Day 20-2:", n)
	//}
}
