package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type RawTile struct {
	id    int
	s     []string
	fixed bool
	found bool
}

// rotate rotates the []string of the tile 90 degrees to the right
func (t *RawTile) rotate() {
	var r []string
	for i := 0; i < len(t.s[0]); i++ {
		tmp := ""
		for j := len(t.s) - 1; j >= 0; j-- {
			tmp += string(t.s[j][i])
		}
		r = append(r, tmp)
	}
	t.s = r
}

func (t *RawTile) flipHorizontal() {
	var r []string
	for _, s := range t.s {
		tmp := ""

		for i := len(s) - 1; i >= 0; i-- {
			tmp += string(s[i])
		}

		r = append(r, tmp)
	}
	t.s = r
}

func (t *RawTile) topString() string {
	return t.s[0]
}

// bottom string is read from right to left
func (t *RawTile) bottomString() string {
	return t.s[len(t.s)-1]
}

// Left string is read from bottom to top
func (t *RawTile) leftString() string {
	s := ""
	for _, i := range t.s {
		s += string(i[0])
	}

	return s
}

// right string is read from top to bottom
func (t *RawTile) rightString() string {
	s := ""
	for _, i := range t.s {
		s += string(i[len(i)-1])
	}
	return s
}

func (t *RawTile) flipVertical() {
	var r []string
	for i := len(t.s) - 1; i >= 0; i-- {
		r = append(r, t.s[i])
	}
	t.s = r
}

func (t *RawTile) printTile() {
	fmt.Println()
	for _, i := range t.s {
		fmt.Println(i)
	}
}

type Tile struct {
	id                       int
	s                        []string
	top, bottom, left, right int
}

// reverseString reverses the order of the characters in a sting
func reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

// parseInput scans a file line by line and creates Tile object appending them to a
// list. This list is returned.
func parseInput(s *bufio.Scanner) []RawTile {
	var rawTiles []RawTile             // output list
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
			rawTiles = append(rawTiles, RawTile{id, rawTile, false, false})
			rawTile = []string{}
			continue
		}
		// Just add raw tile string to tileList
		rawTile = append(rawTile, r)
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return rawTiles
}

func rotator(t RawTile, n int) RawTile {
	tmp := RawTile{t.id, t.s, t.fixed, t.found}
	for i := 0; i < n; i++ {
		tmp.rotate()
	}
	return tmp
}

func flipper(t RawTile, vert bool) RawTile {

	tmp := RawTile{t.id, t.s, t.fixed, t.found}
	if vert {
		tmp.flipVertical()
		return tmp
	}
	tmp.flipHorizontal()
	return tmp
}

func stitch(rawTiles []RawTile) []Tile {
	var tileList []Tile      // save tiles with neighbours here
	rawTiles[0].fixed = true // start with the first tile

	for {

		var rt RawTile // look for neighbours from this
		var rtI int    // this is its index
		stop := true   // stop looking when we have a tile to start with
		for i, t := range rawTiles {
			if t.fixed && !t.found {
				stop = false
				rtI = i
				rt = t
			}
		}

		if stop {
			break
		}

		newTile := Tile{
			id: rt.id,
			s:  rt.s,
		}

		for i := range rawTiles {
			if rawTiles[i].id == rt.id {
				continue
			}

			// possible neighbour
			pn := rawTiles[i]

			found := false
			var pns []RawTile

			if pn.fixed {
				pns = []RawTile{pn}
			} else {
				pns = []RawTile{
					pn,                             // 0,
					flipper(pn, false),             // 0, hflip
					rotator(pn, 1),                 // 90,
					flipper(rotator(pn, 1), true),  // 90, vflip
					rotator(pn, 2),                 // 180,
					flipper(rotator(pn, 2), false), // 180, hflip
					rotator(pn, 3),                 // 270,
					flipper(rotator(pn, 3), true),  // 270 vflip
				}
			}

			for _, p := range pns {
				if rt.bottomString() == p.topString() {
					newTile.bottom = p.id
					rawTiles[i] = p
					found = true
					break
				} else if rt.topString() == p.bottomString() {
					newTile.top = p.id
					rawTiles[i] = p
					found = true
					break
				} else if rt.rightString() == p.leftString() {
					newTile.right = p.id
					rawTiles[i] = p
					found = true
					break
				} else if rt.leftString() == p.rightString() {
					newTile.left = p.id
					rawTiles[i] = p
					found = true
					break
				}
			}

			if found {
				rawTiles[i].fixed = true
			}

		}
		rawTiles[rtI].found = true
		tileList = append(tileList, newTile)
	}
	return tileList
}

func arrangeTiles(s *bufio.Scanner) int {
	rawTiles := parseInput(s) // Read input
	ts := stitch(rawTiles)
	sum := 1
	for _, i := range ts {
		tmp := 0
		if i.top > 0 {
			tmp++
		}
		if i.bottom > 0 {
			tmp++
		}
		if i.left > 0 {
			tmp++
		}
		if i.right > 0 {
			tmp++
		}
		if tmp == 2 {
			sum *= i.id
		}
	}
	return sum
}

func findByID(id int, ts []Tile) Tile {
	for _, i := range ts {
		if i.id == id {
			return i
		}
	}
	return ts[0]
}

func shrinkS(s []string) []string {
	var tmp []string
	for _, s := range s[1 : len(s)-1] {
		tmp = append(tmp, s[1:len(s)-1])
	}
	return tmp
}

func getSea(ts []Tile) []string {
	// find left top
	var start Tile
	for _, i := range ts {
		if i.left == 0 && i.top == 0 {
			start = i
		}
	}
	// save bottom id
	bottomID := start.bottom

	// containers
	var sea []string
	var tmp []string
	for {

		// add start to tmp
		smallStart := shrinkS(start.s)

		if len(tmp) == 0 {
			for _, s := range smallStart {
				tmp = append(tmp, s)
			}
		} else {
			for i, s := range smallStart {
				tmp[i] += s
			}
		}

		// find right neighbour
		start = findByID(start.right, ts)
		smallStart = shrinkS(start.s)

		// if new start has no neighbour we need to go to the next row
		if start.right == 0 {
			// add last element to tmp
			for i, s := range smallStart {
				tmp[i] += s
			}
			// add tmp to sea
			for _, t := range tmp {
				sea = append(sea, t)
			}
			// reset tmp
			tmp = []string{}

			// if the bottomID == 0, we are done there are no more rows to go
			if bottomID == 0 {
				break
			}

			// find first element of next row using bottom id
			start = findByID(bottomID, ts)
			bottomID = start.bottom
		}
	}
	return sea

}

func monster() ([]int, []int) {

	m := []string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	}
	xs := []int{}
	ys := []int{}

	for y, i := range m {
		for x, ii := range i {
			if ii == '#' {
				xs = append(xs, x)
				ys = append(ys, y)
			}
		}
	}

	return xs, ys
}

func replaceAtIndex(in string, i int) string {
	out := []rune(in)
	out[i] = 'O'
	return string(out)
}

func findSeaMonsters(s *bufio.Scanner) int {
	rawTiles := parseInput(s) // Read input
	ts := stitch(rawTiles)
	ss := getSea(ts)
	xs, ys := monster()

	rt := RawTile{0, ss, false, false}

	seas := []RawTile{
		rt,                             // 0,
		flipper(rt, false),             // 0, hflip
		rotator(rt, 1),                 // 90,
		flipper(rotator(rt, 1), true),  // 90, vflip
		rotator(rt, 2),                 // 180,
		flipper(rotator(rt, 2), false), // 180, hflip
		rotator(rt, 3),                 // 270,
		flipper(rotator(rt, 3), true),  // 270 vflip
	}

	maxCount := 0
	var maxSea []string

	for _, s := range seas {
		sea := s.s

		monsterCounter := 0
		// Loop over sea coordinates - monster size
		for y := range sea[0 : len(sea)-3] {
			for x := range sea[y][0 : len(sea[y])-20] {

				// loop over monster coordinates
				cnt := true
				for i := range xs {
					if sea[y+ys[i]][x+xs[i]] != '#' {
						cnt = false
						break
					}
				}
				if cnt == true {
					monsterCounter++
					for i := range xs {
						sea[y+ys[i]] = replaceAtIndex(sea[y+ys[i]], x+xs[i])
					}
				}
			}
		}
		if monsterCounter > maxCount {
			maxCount = monsterCounter
			maxSea = sea
		}
	}

	for _, s := range maxSea {
		fmt.Println(s)
	}
	fmt.Println("\n", maxCount, "monsters found!!!! HELP")

	seaCount := 0
	re := regexp.MustCompile("#")

	for _, s := range maxSea {
		c := re.FindAllStringIndex(s, -1)
		seaCount += len(c)
	}

	return seaCount
}

func main() {
	part2 := true
	file, err := os.Open("data_20.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !part2 {
		n := arrangeTiles(scanner)
		fmt.Println("Day 20-1:", n)
	} else {
		n := findSeaMonsters(scanner)
		fmt.Println("Day 20-2:", n)
	}
}
