package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// Allergen is a container to save the possible ingredients in
type Allergen struct {
	options  map[string]int
	allergen string
}

// parseFile reads the input and fills a few maps with the data
func parseFile(s *bufio.Scanner) (map[string]Allergen, map[string]int, map[string]int) {
	allergenList := make(map[string]Allergen)
	ingredientList := make(map[string]int)
	allergenCount := make(map[string]int)

	for s.Scan() {

		// Read and split line
		r := strings.Trim(s.Text(), "\n")
		sp := strings.Split(r, "(")

		// Handle allergens
		allergens := strings.Split(strings.Trim(sp[1], ")"), ",")
		for i := range allergens {
			allergens[i] = strings.Trim(strings.ReplaceAll(allergens[i], "contains", ""), " ")
		}

		// Parse out the ingredients
		ingredients := strings.Split(strings.Trim(sp[0], " "), " ")

		for _, a := range allergens {
			allergenCount[a]++
			// allergen already in list
			if _, ok := allergenList[a]; ok {
				// add n ingredients
				for _, i := range ingredients {
					allergenList[a].options[i]++
				}
			} else {
				tmp := make(map[string]int)
				for _, i := range ingredients {
					tmp[i]++
				}
				allergenList[a] = Allergen{tmp, a}
			}
		}

		for _, i := range ingredients {
			ingredientList[i]++
		}

	}

	if err := s.Err(); err != nil {
		panic(err)
	}
	return allergenList, allergenCount, ingredientList
}

// analyseInput filters down the ingredients and allergens
func analyseInput(s *bufio.Scanner) (map[string]int, map[string]string) {
	// get input data in lots of different containers
	allergenList, allergenCount, ingredientList := parseFile(s)
	result := make(map[string]string)

	for {
		// if we do not find anything this will result in exiting the loop
		done := true

		// save the found ingredient and allergen
		var eliminateIngredient string
		var eliminateAllergen string

		for a, in := range allergenList {

			// If we have a single option for an allergen that should be a hit
			if len(in.options) == 1 {
				eliminateAllergen = a
				var tmpI string
				for k := range in.options {
					tmpI = k
				}
				eliminateIngredient = tmpI
				done = false
				break
			}

			// Look for a case where the number of ingredient occurences is equal to
			// the number of allergen occurences
			c := 0
			var tmpI string
			for i, cnt := range in.options {
				if cnt == allergenCount[a] {
					tmpI = i
					c++
				}
			}
			if c == 1 {
				eliminateAllergen = a
				eliminateIngredient = tmpI
				done = false
				break
			}

		}

		// Save result
		result[eliminateAllergen] = eliminateIngredient

		// Delete stuff we found
		delete(allergenList, eliminateAllergen)
		delete(allergenCount, eliminateAllergen)
		delete(ingredientList, eliminateIngredient)
		for i := range allergenList {
			delete(allergenList[i].options, eliminateIngredient)
		}

		// stop looking for hits
		if done || len(allergenList) == 0 {
			break
		}
	}
	return ingredientList, result

}

// checkAllergens solves part 1
func checkAllergens(s *bufio.Scanner) int {
	ingredientList, _ := analyseInput(s)
	sum := 0
	for _, v := range ingredientList {
		sum += v
	}
	return sum
}

// Pair is a simple container to sort based on allergen
type Pair struct {
	al, in string
}

// stockRaft gives the answer to part 2
func stockRaft(s *bufio.Scanner) string {
	_, result := analyseInput(s) // parse input

	// Convert to slice of Pair
	var p []Pair
	for k, v := range result {
		p = append(p, Pair{k, v})
	}
	// sort p
	sort.Slice(p, func(i, j int) bool {
		return p[i].al < p[j].al
	})

	// create slice of string
	r := []string{}
	for _, i := range p {
		r = append(r, i.in)
	}

	// join slice with "," separator
	return strings.Join(r[:], ",")
}

func main() {
	part2 := true
	file, err := os.Open("data_21.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !part2 {
		n := checkAllergens(scanner)
		fmt.Println("Day 21-1:", n)
	} else {
		n := stockRaft(scanner)
		fmt.Println("Day 19-2:", n)
	}
}
