package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type board [][]string

// readlines parses all lines in a scanner to ints and returns them as a slice.
func readlines(s *bufio.Scanner) board {
	var b board

	// Read lines and convert to int, then append to slice
	for s.Scan() {
		row := strings.Trim(s.Text(), "\n")
		spl := strings.Split(row, "")
		b = append(b, spl)
	}

	if err := s.Err(); err != nil {
		panic(err)
	}
	return b
}

func findNeighours(x, y int, b board) int {
	is := []int{x - 1, x, x + 1}
	js := []int{y - 1, y, y + 1}
	cnt := 0
	for _, i := range is {
		for _, j := range js {
			if i < 0 || j < 0 || i >= len(b) || j >= len(b[0]) {
				continue
			}
			if i == x && j == y {
				continue
			}
			if b[i][j] == "#" {
				cnt++
			}

		}
	}
	return cnt
}

func bPrint(b board) {
	for _, i := range b {
		fmt.Println(i)
	}
	fmt.Println()
}

func copyBoard(b board) board {
	nb := make(board, len(b))
	for i := range b {
		nb[i] = make([]string, len(b[i]))
		copy(nb[i], b[i])
	}
	return nb
}

func modelSeats(s *bufio.Scanner) int {
	b := readlines(s)  // orginal board
	nb := copyBoard(b) // copy of the board to make changes on
	changed := false   // if this does not change we reached equilibrium
	c := 0

	for {
		for i := 0; i < len(nb); i++ {
			for j := 0; j < len(nb[0]); j++ {

				occ := findNeighours(i, j, b)
				if b[i][j] == "L" && occ == 0 {
					nb[i][j] = "#"
					changed = true
				}

				if b[i][j] == "#" && occ >= 4 {
					nb[i][j] = "L"
					changed = true
				}
			}
		}
		c++

		// Break out if we did not change anything
		if !changed {
			break
		}

		b = copyBoard(nb)
		nb = copyBoard(b)
		changed = false

	}
	cnt := 0
	for i := 0; i < len(nb); i++ {
		for j := 0; j < len(nb[0]); j++ {
			if nb[i][j] == "#" {
				cnt++
			}
		}
	}
	return cnt
}

// Part 2 ==============================================================================

// Coord saves x and y
type Coord struct {
	x, y int
}

func parseLines(b board) map[Coord]bool {
	m := make(map[Coord]bool)
	for x := 0; x < len(b[0]); x++ {
		for y := 0; y < len(b); y++ {
			if b[y][x] != "." {
				c := Coord{x, y}
				if b[y][x] == "#" {
					m[c] = true
				} else {
					m[c] = false
				}
			}
		}
	}
	return m
}

func modelSeatsDistance(s *bufio.Scanner) int {

	area := readlines(s)
	bPrint(area)
	seats := parseLines(area)

	for {
		changed := false
		newSeats := make(map[Coord]bool)

		// fmt.Println(seats)
		for coord, occ := range seats {

			count := 0

			// Up
			for y := coord.y - 1; y >= 0; y-- {
				seat, ok := seats[Coord{coord.x, y}]
				if ok {
					if seat {
						count++
					}
					break
				}
			}

			// down
			for y := coord.y + 1; y < len(area); y++ {
				seat, ok := seats[Coord{coord.x, y}]
				if ok {
					if seat {
						count++
					}
					break
				}
			}

			// left
			for x := coord.x - 1; x >= 0; x-- {
				seat, ok := seats[Coord{x, coord.y}]
				if ok {
					if seat {
						count++
					}
					break
				}
			}

			// right
			for x := coord.x + 1; x < len(area[0]); x++ {
				seat, ok := seats[Coord{x, coord.y}]
				if ok {
					if seat {
						count++
					}
					break
				}
			}

			// diagonals
			ul := false
			ur := false
			dl := false
			dr := false
			for i := 1; i < 2*len(area); i++ {

				// top left
				if !ul {
					xx := coord.x - i
					yy := coord.y + i
					if xx >= 0 && yy < len(area) {
						seat, ok := seats[Coord{xx, yy}]
						if ok {
							if seat {
								count++
							}
							ul = true
						}
					}
				}

				// top right
				if !ur {
					xx := coord.x + i
					yy := coord.y + i
					if xx < len(area[0]) && yy < len(area) {
						seat, ok := seats[Coord{xx, yy}]
						if ok {
							if seat {
								count++
							}
							ur = true
						}
					}
				}

				// bottom left
				if !dl {
					xx := coord.x - i
					yy := coord.y - i
					if xx >= 0 && yy >= 0 {
						seat, ok := seats[Coord{xx, yy}]
						if ok {
							if seat {
								count++
							}
							dl = true
						}
					}
				}

				// bottom right
				if !dr {
					xx := coord.x + i
					yy := coord.y - i
					if xx < len(area[0]) && yy >= 0 {
						seat, ok := seats[Coord{coord.x + i, coord.y - i}]
						if ok {
							if seat {
								count++
							}
							dr = true
						}
					}
				}
			}

			// fmt.Println(coord, count)
			// Based on neightbour counts, change the seats occupance
			if !occ && count == 0 {
				// no neighbours, sit here
				newSeats[coord] = true
				changed = true
			} else if occ && count >= 5 {
				// 5 or more neighbours. go away
				newSeats[coord] = false
				changed = true
			} else {
				// nothing happens
				newSeats[coord] = occ
			}

		}

		// if nothing changed we have reached an equilibrium
		if !changed {
			break
		}

		seats = newSeats

	}

	// Count number of occupied seats
	cnt := 0
	for _, v := range seats {
		if v {
			cnt++
		}
	}

	return cnt
}

func main() {
	part2 := true
	file, err := os.Open("data_11.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !part2 {
		n := modelSeats(scanner)
		fmt.Println("Day 11-1:", n)
	} else {
		n := modelSeatsDistance(scanner)
		fmt.Println("Day 11-2:", n)
	}
}
