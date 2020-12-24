package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Coord struct {
	x, y int
}

// matchToCoordinate loops through a list of directions and returns the final
// coordinate as a Coord.
func matchToCoordinate(l []string) Coord {
	x := 0
	y := 0
	for _, d := range l {
		switch d {
		case "e":
			x++
		case "w":
			x--
		case "ne":
			y++
		case "nw":
			x--
			y++
		case "se":
			x++
			y--
		case "sw":
			y--
		}
	}
	return Coord{x, y}
}

// readTiles reads the input line by line, gets the coordinate and reverses the bool in
// a map[Coord]bool to simulate flipping tiles.
func readTiles(s *bufio.Scanner) map[Coord]bool {

	pointList := make(map[Coord]bool)

	for s.Scan() {
		r := strings.Trim(s.Text(), "\n")

		// Split the text, first try a two character solution, if the does not fit. Take a
		// single character
		var l []string
		for len(r) > 1 {
			long := r[:2]
			if long == "se" || long == "sw" || long == "ne" || long == "nw" {
				l = append(l, long)
				r = r[2:]
			} else {
				l = append(l, string(r[0]))
				r = r[1:]
			}
		}

		// Include the final character
		if len(r) == 1 {
			l = append(l, r)
		}

		// Calculate the coordinate this instruction points to
		c := matchToCoordinate(l)
		pointList[c] = !pointList[c]
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	return pointList
}

// tileFloor is the answer to pt. 1. We parse the file to a map[Coord]bool and then
// count the number of true's in the map.
func tileFloor(s *bufio.Scanner) (sum int) {
	l := readTiles(s)
	sum = 0
	for _, v := range l {
		if v {
			sum++
		}
	}
	return
}

// getMinMax finds the min and max x's and y's and returns them decremented by 1
func getMinMax(m map[Coord]bool) (int, int, int, int) {
	minX, minY := 1000000, 1000000
	maxX, maxY := -1000000, -1000000
	for c := range m {
		if c.x < minX {
			minX = c.x
		}
		if c.x > maxX {
			maxX = c.x
		}

		if c.y < minY {
			minY = c.y
		}

		if c.y > maxY {
			maxY = c.y
		}
	}
	return minX - 1, maxX + 1, minY - 1, maxY + 1
	//return minX, maxX, minY, maxY
}

// createPossibleCoordinates creates a list of coordinates to check for neighbours in
// this hexagonal game of life.
func createPossibleCoordinates(m map[Coord]bool) []Coord {
	minX, maxX, minY, maxY := getMinMax(m) // get mins and maxes
	var l []Coord

	// create all possible coordinates
	for x := minX; x < maxX+1; x++ {
		for y := minY; y < maxY+1; y++ {
			l = append(l, Coord{x, y})
		}
	}
	return l
}

func countNeighbours(c Coord, m map[Coord]bool) int {
	// create neighbours
	ns := []Coord{
		{c.x + 1, c.y},
		{c.x - 1, c.y},
		{c.x, c.y + 1},
		{c.x - 1, c.y + 1},
		{c.x + 1, c.y - 1},
		{c.x, c.y - 1},
	}

	cnt := 0
	for _, n := range ns {

		// check if neighbour is in our map
		b, found := m[n]

		// if not, it is white and we increment
		if !found {
			cnt++
		}

		// if we found it and b is false we increment
		if found && !b {
			cnt++
		}

	}
	return cnt
}

func simulateTiles(s *bufio.Scanner) int {
	m := readTiles(s) // read day0 tiles

	// do 100 steps in the simulation
	i := 0
	for i < 100 {
		// new map to write chages in
		newTiles := make(map[Coord]bool)
		// possible coordinates
		coordList := createPossibleCoordinates(m)
		for _, c := range coordList {

			// what is our current colour?
			black := false
			if v, ok := m[c]; ok {
				black = v
			}

			// Count white neighbours and subtract that from 6 to get the number of black
			// neightbours.
			n := 6 - countNeighbours(c, m)

			// do logic based on color and number of black neighbours
			if black && (n == 0 || n > 2) {
				newTiles[c] = false
			} else if !black && n == 2 {
				newTiles[c] = true
			} else {
				newTiles[c] = m[c]
			}

		}

		// overwrite m with result of timestep
		m = newTiles
		i++
	}

	// count number of black tiles
	sum := 0
	for _, t := range m {
		if t {
			sum++
		}
	}

	return sum
}

func main() {
	part2 := true
	file, err := os.Open("data_24.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !part2 {
		n := tileFloor(scanner)
		fmt.Println("Day 24-1:", n)
	} else {
		n := simulateTiles(scanner)
		fmt.Println("Day 24-2:", n)
	}
}
