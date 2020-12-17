package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y, z int
}

func parseInput(s *bufio.Scanner) map[Point]bool {
	grid := make(map[Point]bool)
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
				grid[Point{x, y, 0}] = true
			}
		}
	}

	return grid
}

func generateNeighbours(p Point) []Point {
	var n []Point
	for x := p.x - 1; x < p.x+2; x++ {
		for y := p.y - 1; y < p.y+2; y++ {
			for z := p.z - 1; z < p.z+2; z++ {
				if x == p.x && y == p.y && z == p.z {
					continue
				}
				n = append(n, Point{x, y, z})
			}
		}
	}
	return n
}

func findMinMax(g map[Point]bool) (Point, Point) {
	minX := 10000
	maxX := 0
	minY := 10000
	maxY := 0
	minZ := 10000
	maxZ := 0
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
	// We need to expand our grid every strp, so '-1' and '+1'
	return Point{minX - 1, minY - 1, minZ - 1}, Point{maxX + 1, maxY + 1, maxZ + 1}
}

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

func ticketScaningErrorRate(s *bufio.Scanner) int {

	grid := parseInput(s)

	c := 0
	for c < 6 {
		minP, maxP := findMinMax(grid)

		gridPoints := generateGridPoints(minP, maxP)
		newGrid := make(map[Point]bool)

		for _, p := range gridPoints {
			nl := generateNeighbours(p) // neighbour list

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

func main() {
	part2 := false
	file, err := os.Open("data_17.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !part2 {
		n := ticketScaningErrorRate(scanner)
		fmt.Println("Day 17-1:", n)
	}
	// else {
	// 	n := ticketScaningErrorRate(scanner)
	// 	fmt.Println("Day 17-2:", n)
	// }
}
