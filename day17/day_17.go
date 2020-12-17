package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// =====================================================================================
// PART II
// =====================================================================================

// Part II is similar to part I, but with an added dimension

// Point4 is a 4 dimensional coordinate
type Point4 struct {
	x, y, z, w int
}

func parseInput4(s *bufio.Scanner) map[Point4]bool {
	grid := make(map[Point4]bool)
	var tmp [][]string
	for s.Scan() {
		r := strings.Split(strings.Trim(s.Text(), "\n"), "")
		tmp = append(tmp, r)
	}
	if err := s.Err(); err != nil {
		panic(err)
	}

	for x := 0; x < len(tmp); x++ {
		for y := 0; y < len(tmp); y++ {
			if tmp[x][y] == "#" {
				grid[Point4{x, y, 0, 0}] = true // z and w are 0
			}
		}
	}

	return grid
}

func generateNeighbours4(p Point4) []Point4 {
	var n []Point4
	for x := p.x - 1; x < p.x+2; x++ {
		for y := p.y - 1; y < p.y+2; y++ {
			for z := p.z - 1; z < p.z+2; z++ {
				for w := p.w - 1; w < p.w+2; w++ {
					if x == p.x && y == p.y && z == p.z && w == p.w {
						continue
					}
					n = append(n, Point4{x, y, z, w})
				}
			}
		}
	}
	return n
}

func findMinMax4(g map[Point4]bool) (Point4, Point4) {
	minX := 10000
	maxX := 0
	minY := 10000
	maxY := 0
	minZ := 10000
	maxZ := 0
	minW := 10000
	maxW := 0
	for p := range g {
		if p.x < minX {
			minX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.z < minZ {
			minZ = p.z
		}
		if p.w < minW {
			minW = p.w
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
		if p.z > maxZ {
			maxZ = p.z
		}
		if p.w > maxW {
			maxW = p.w
		}
	}
	// We need to expand our grid every strp, so '-1' and '+1'
	pMin := Point4{minX - 1, minY - 1, minZ - 1, minW - 1}
	pMax := Point4{maxX + 1, maxY + 1, maxZ + 1, maxW + 1}
	return pMin, pMax
}

func generateGridPoints4(min, max Point4) []Point4 {
	var ps []Point4
	for x := min.x; x < max.x+1; x++ {
		for y := min.y; y < max.y+1; y++ {
			for z := min.z; z < max.z+1; z++ {
				for w := min.w; w < max.w+1; w++ {
					ps = append(ps, Point4{x, y, z, w})
				}
			}
		}
	}
	return ps
}

func simulateGrid4(s *bufio.Scanner) int {

	grid := parseInput4(s)

	c := 0
	for c < 6 {
		minP, maxP := findMinMax4(grid)

		gridPoints := generateGridPoints4(minP, maxP)
		newGrid := make(map[Point4]bool)

		for _, p := range gridPoints {
			nl := generateNeighbours4(p) // neighbour list

			nc := 0
			for _, n := range nl {
				if grid[n] {
					nc++
				}
			}

			if grid[p] {
				if nc == 2 || nc == 3 {
					newGrid[p] = true
				}
			} else {
				if nc == 3 {
					newGrid[p] = true
				}
			}
		}
		grid = newGrid
		c++
	}

	return len(grid)
}

// =====================================================================================
// PART I
// =====================================================================================

// Point is a data structure to save coordinates in
type Point struct {
	x, y, z int
}

// parsInput reads a file and convert the "#"'s to a Point object which is used as a key
// in a map.
func parseInput(s *bufio.Scanner) map[Point]bool {

	// Read each line and split on ",", append to tmp
	var tmp [][]string
	for s.Scan() {
		r := strings.Split(strings.Trim(s.Text(), "\n"), "")
		tmp = append(tmp, r)
	}
	if err := s.Err(); err != nil {
		panic(err)
	}

	// Transform the list of list of strings to a single map[Point]bool
	grid := make(map[Point]bool)
	for x := 0; x < len(tmp); x++ {
		for y := 0; y < len(tmp); y++ {
			if tmp[x][y] == "#" {
				grid[Point{x, y, 0}] = true
			}
		}
	}

	return grid
}

// generateNeighbours takes a point and returns all neighbour Points in all directions.
func generateNeighbours(p Point) []Point {
	var n []Point
	for x := p.x - 1; x < p.x+2; x++ {
		for y := p.y - 1; y < p.y+2; y++ {
			for z := p.z - 1; z < p.z+2; z++ {
				// Do not include input point p in the neighbours list
				if x == p.x && y == p.y && z == p.z {
					continue
				}
				n = append(n, Point{x, y, z})
			}
		}
	}
	return n
}

// findMinMax loops over the grid and returns the two Points. One containing the minimum
// x,y,z and one containing the maximum x,y,z.
func findMinMax(g map[Point]bool) (Point, Point) {
	// initialize mins and maxes
	minX := 10000
	maxX := 0
	minY := 10000
	maxY := 0
	minZ := 10000
	maxZ := 0

	// For each point in the grid, check if we have smaller or larger coordinates then
	// the current mins and maxes.
	for p := range g {
		if p.x < minX {
			minX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.z < minZ {
			minZ = p.z
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
		if p.z > maxZ {
			maxZ = p.z
		}
	}
	// We need to expand our grid every step, so '-1' and '+1'
	return Point{minX - 1, minY - 1, minZ - 1}, Point{maxX + 1, maxY + 1, maxZ + 1}
}

// generateGridPoints generates all points on a grid between min and max+1
func generateGridPoints(min, max Point) []Point {
	var ps []Point
	for x := min.x; x < max.x+1; x++ {
		for y := min.y; y < max.y+1; y++ {
			for z := min.z; z < max.z+1; z++ {
				ps = append(ps, Point{x, y, z})
			}
		}
	}
	return ps
}

// simulateGrid simulates 6 timesteps of a grid in three dimension
func simulateGrid(s *bufio.Scanner) int {
	grid := parseInput(s) // read the initial grid state

	c := 0 // number of steps to simulate
	for c < 6 {
		minP, maxP := findMinMax(grid)               // get grid boundaries
		gridPoints := generateGridPoints(minP, maxP) // get all points in the grid
		newGrid := make(map[Point]bool)              // save the new grid in this map

		// Loop over all points in the grid
		for _, p := range gridPoints {
			nl := generateNeighbours(p) // neighbour list

			// Count the neighbours of point p
			nc := 0
			for _, n := range nl {
				if grid[n] {
					nc++
				}
			}

			// Depending on the number of neighbours, put point p in newGrid
			if grid[p] {
				if nc == 2 || nc == 3 {
					newGrid[p] = true
				}
			} else {
				if nc == 3 {
					newGrid[p] = true
				}
			}
		}

		grid = newGrid // Set grid to be newGrid for the next step
		c++            // increase simulation steps
	}

	// return number of active points
	return len(grid)
}

func main() {
	part2 := true
	file, err := os.Open("data_17.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !part2 {
		n := simulateGrid(scanner)
		fmt.Println("Day 17-1:", n)
	} else {
		n := simulateGrid4(scanner)
		fmt.Println("Day 17-2:", n)
	}
}
