package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Game struct {
	p1, p2 []int
}

// getCardLists reads lines from a file scanner and returns two lists with card
// numbers, one for each player.
func getCardLists(s *bufio.Scanner) ([]int, []int) {
	player1 := true // put number in p1 or p2
	p1 := []int{}   // save p1 deck
	p2 := []int{}   // save p2 deck
	for s.Scan() {
		r := strings.Trim(s.Text(), "\n")

		// Switch putting cards into p1 deck into p2 deck
		if r == "Player 2:" {
			player1 = false
		}
		// If we can convert a line to a number, we put it in a deck.
		if n, err := strconv.Atoi(r); err == nil {
			if player1 {
				p1 = append(p1, n)
			} else {
				p2 = append(p2, n)
			}
		}
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	return p1, p2
}

// calcWinningScore, given two lists, calculates the score of the winner
func calcWinningScore(p1, p2 []int) int {
	// Put winning list in winner
	var winner []int
	if len(p1) > len(p2) {
		winner = p1
	} else {
		winner = p2
	}

	// Calculate winning score, scale is decreasing.
	sum := 0
	scale := len(winner)
	for _, w := range winner {
		sum += w * scale
		scale--
	}

	return sum

}

func playGame(s *bufio.Scanner) int {
	// Load card decks
	p1Cards, p2Cards := getCardLists(s)

	// play until 1 deck is empty
	for {

		// handle cases one player wins
		if p1Cards[0] > p2Cards[0] {
			p1Cards = append(p1Cards, p1Cards[0], p2Cards[0])
			// remove played cards from front of the decks
			p1Cards = p1Cards[1:]
			p2Cards = p2Cards[1:]
		} else {
			p2Cards = append(p2Cards, p2Cards[0], p1Cards[0])
			// remove played cards from front of the decks
			p2Cards = p2Cards[1:]
			p1Cards = p1Cards[1:]
		}

		// escape if one deck is empty
		if len(p1Cards) == 0 || len(p2Cards) == 0 {
			break
		}
	}

	return calcWinningScore(p1Cards, p2Cards)
}

// gamePlayed checks in mem(ory) if a game is already played
func gamePlayed(mem []Game, g Game) bool {
	for _, gg := range mem {
		if reflect.DeepEqual(gg.p1, g.p1) && reflect.DeepEqual(gg.p2, g.p2) {
			return true
		}
	}
	return false
}

// play does turns in the card game recursively.
func play(p1, p2 []int, mem []Game) ([]int, []int, bool) {

	// Check if we had this exact case already, if we have we return
	g := Game{p1, p2} // create something to recognise in a map
	if gamePlayed(mem, g) {
		p1 = append(p1[1:], p1[0], p2[0])
		return p1, p2[1:], true
	} else {
		mem = append(mem, g)
	}

	// Escape because one deck is empty
	if len(p1) == 0 || len(p2) == 0 {
		if len(p2) == 0 {
			return p1, p2, true
		}
		return p1, p2, false
	}

	// start playing a normal turn
	c1 := p1[0] // play first cards
	c2 := p2[0]
	p1 = p1[1:] // remove first cards from deck
	p2 = p2[1:]

	p1Win := false
	// Check if we need a sub game or not
	if c1 <= len(p1) && c2 <= len(p2) {
		tmp1 := append([]int{}, p1[:c1]...)
		tmp2 := append([]int{}, p2[:c2]...)
		// Play a subgame, this will have its own memory, thats why we add []Game{}.
		_, _, player1Win := play(tmp1, tmp2, []Game{})
		if player1Win {
			p1Win = true
		}
	} else {
		// play normal turn, check who has the highest card.
		if c1 > c2 {
			p1Win = true
		}
	}

	// create new decks, by giving the winner the played cards.
	if p1Win {
		p1 = append(p1, c1, c2)
	} else {
		p2 = append(p2, c2, c1)
	}

	// Play enother round
	return play(p1, p2, mem)
}

// playRecursiveGame starts the recursive game and calculates the score of the winner.
func playRecursiveGame(s *bufio.Scanner) int {
	p1Cards, p2Cards := getCardLists(s)           // Load decks
	r1, r2, _ := play(p1Cards, p2Cards, []Game{}) // play recursive game
	return calcWinningScore(r1, r2)               // return winning score
}

func main() {
	part2 := true
	file, err := os.Open("data_22.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !part2 {
		n := playGame(scanner)
		fmt.Println("Day 22-1:", n)
	} else {
		n := playRecursiveGame(scanner)
		fmt.Println("Day 22-2:", n)
	}
}
