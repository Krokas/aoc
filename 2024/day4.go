package aoc2024

import (
	"bufio"
	"fmt"
	"strings"
)

type WordOccurance struct {
	found bool
	position Position
	oreantation int
}

type Position struct {
	x int
	y int
}

const (
	RIGHT = 0
	BOTTOM_RIGHT = 1
	BOTTOM = 2
	BOTTOM_LEFT = 3
	LEFT = 4
	TOP_LEFT = 5
	TOP = 6
	TOP_RIGHT = 7
)

func GetSumOfWordOccuranceInWordSearch(scanner *bufio.Scanner) int {
	var wordSearch []string
	var lineWidth int
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			chars := strings.Split(line, "")
			wordSearch = append(wordSearch, chars...)

			lineWidth = len(chars)
		}
	}
	occuranceCount := 0
	var foundOccurances []WordOccurance 
	for i := 0; i < len(wordSearch); i++ {
		// retest
		for r := 0; r < 8; r++ {

			occurance := findWordOccurance("XMAS", i, wordSearch, lineWidth, r)
			if occurance.found {
				if !isOccuranceFound(occurance, foundOccurances) {
					foundOccurances = append(foundOccurances, occurance)
					occuranceCount++
				}
			}
		}
		
	}

	return occuranceCount
}

func GetSumOfXCrossedWords(scanner *bufio.Scanner) int {
	var wordSearch []string
	var lineWidth int
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			chars := strings.Split(line, "")
			wordSearch = append(wordSearch, chars...)

			lineWidth = len(chars)
		}
	}
	occuranceCount := 0
	var foundOccurances []WordOccurance 
	for i := 0; i < len(wordSearch); i++ {
		// retest
		for r := 0; r < 8; r++ {

			occurance := findWordOccurance("MAS", i, wordSearch, lineWidth, r)
			if occurance.found {
				if !isOccuranceFound(occurance, foundOccurances) {
					foundOccurances = append(foundOccurances, occurance)
				}
			}
		}
		
	}

	var xPossible []WordOccurance

	for i := 0; i < len(foundOccurances); i++ {
		oc := foundOccurances[i]
		if oc.oreantation == TOP_LEFT || oc.oreantation == TOP_RIGHT || oc.oreantation == BOTTOM_LEFT || oc.oreantation == BOTTOM_RIGHT {
			xPossible = append(xPossible, oc)
			if oc.position.x == 7 && oc.position.y == 1 {
				fmt.Println(oc)
			}
		}
	}

	for a := 0; a < len(xPossible); a++ {
		for b := 0; b < len(xPossible); b++ {
			occA := xPossible[a]
			occB := xPossible[b]

			// from the bottom - 5
			if occB.position.x + 2 == occA.position.x  && occB.position.y == occA.position.y && occB.oreantation == TOP_RIGHT && occA.oreantation == TOP_LEFT {
				occuranceCount++
			}

			// from the top - 1
			if occB.position.x + 2 == occA.position.x && occB.position.y == occA.position.y && occB.oreantation == BOTTOM_RIGHT && occA.oreantation == BOTTOM_LEFT {
				occuranceCount++
			}

			// if occB.position.x == 5 && occB.position.y == 1 && occA.oreantation == BOTTOM_LEFT {
			// 	fmt.Println("occA", occA, occB.oreantation)
			// }

			// from the left 2
			if occB.position.y == occA.position.y + 2 && occB.position.x == occA.position.x && occB.oreantation == TOP_RIGHT && occA.oreantation == BOTTOM_RIGHT {
				occuranceCount++
			}

			// from the right 2 (1 radau)
			if occB.position.y == occA.position.y + 2 && occB.position.x == occA.position.x && occB.oreantation == TOP_LEFT && occA.oreantation == BOTTOM_LEFT {
				occuranceCount++
			}
		}
	}

	return occuranceCount
}

func findWordOccurance(word string, position int, wordSearch []string, width int, rotation int) WordOccurance {

	height := len(wordSearch) / width
	pos := Position{}
	pos.y = position / width
	pos.x = position % width
	wordSlice := strings.Split(word, "")
	direction := -1
	var charactersFound int

	if wordSearch[position] == wordSlice[0] {
		charactersFound++
		if pos.x < width - 1 && wordSlice[1] == getCharacterInDirection(RIGHT, wordSearch, pos.x, pos.y, width, 1) && rotation == RIGHT {
			direction = RIGHT
			charactersFound++
		} else
		if pos.x < width - 1 && pos.y < height - 1 && wordSlice[1] == getCharacterInDirection(BOTTOM_RIGHT, wordSearch, pos.x, pos.y, width, 1) && rotation == BOTTOM_RIGHT {
			direction = BOTTOM_RIGHT
			charactersFound++
			
		} else
		if pos.y < height - 1 && wordSlice[1] == getCharacterInDirection(BOTTOM, wordSearch, pos.x, pos.y, width, 1) && rotation == BOTTOM {
			direction = BOTTOM
			charactersFound++
		} else
		if pos.x - 1 >= 0 && pos.y < height - 1 && wordSlice[1] == getCharacterInDirection(BOTTOM_LEFT, wordSearch, pos.x, pos.y, width, 1) && rotation == BOTTOM_LEFT {
			direction = BOTTOM_LEFT
			charactersFound++
		} else

		if pos.x - 1 >= 0 && wordSlice[1] == getCharacterInDirection(LEFT, wordSearch, pos.x, pos.y, width, 1) && rotation == LEFT {
			direction = LEFT
			charactersFound++
		} else

		if pos.x - 1 >= 0 && pos.y -1 >= 0 && wordSlice[1] == getCharacterInDirection(TOP_LEFT, wordSearch, pos.x, pos.y, width, 1) && rotation == TOP_LEFT {
			direction = TOP_LEFT
			charactersFound++
		} else

		if pos.y -1 >= 0 && wordSlice[1] == getCharacterInDirection(TOP, wordSearch, pos.x, pos.y, width, 1) && rotation == TOP {
			direction = TOP
			charactersFound++
		} else

		if pos.x < width - 1 && pos.y -1 >= 0 && wordSlice[1] == getCharacterInDirection(TOP_RIGHT, wordSearch, pos.x, pos.y, width, 1) && rotation == TOP_RIGHT {
			direction = TOP_RIGHT
			charactersFound++
		}

		for j := 2; j < len(wordSlice); j++ {
			if direction == RIGHT && pos.x < width - j && wordSlice[j] == getCharacterInDirection(RIGHT, wordSearch, pos.x, pos.y, width, j) {
				charactersFound++
				continue
			}

			if direction == BOTTOM_RIGHT && pos.x < width - j && pos.y < height - j && wordSlice[j] == getCharacterInDirection(BOTTOM_RIGHT, wordSearch, pos.x, pos.y, width, j) {
				charactersFound++
				continue
			}

			if direction == BOTTOM && pos.y < height - j && wordSlice[j] == getCharacterInDirection(BOTTOM, wordSearch, pos.x, pos.y, width, j) {
				charactersFound++
				continue
			}

			if direction == BOTTOM_LEFT && pos.x - j >= 0 && pos.y < height - j && wordSlice[j] == getCharacterInDirection(BOTTOM_LEFT, wordSearch, pos.x, pos.y, width, j) {
				charactersFound++
				continue
			}

			if direction == LEFT && pos.x - j >= 0 && wordSlice[j] == getCharacterInDirection(LEFT, wordSearch, pos.x, pos.y, width, j) {
				charactersFound++
				continue
			}

			if direction == TOP_LEFT && pos.x - j >= 0 && pos.y -j >= 0 && wordSlice[j] == getCharacterInDirection(TOP_LEFT, wordSearch, pos.x, pos.y, width, j) {
				charactersFound++
				continue
			}

			if direction == TOP && pos.y -j >= 0 && wordSlice[j] == getCharacterInDirection(TOP, wordSearch, pos.x, pos.y, width, j) {
				charactersFound++
				continue
			}

			if direction == TOP_RIGHT && pos.x < width - j && pos.y -j >= 0 && wordSlice[j] == getCharacterInDirection(TOP_RIGHT, wordSearch, pos.x, pos.y, width, j) {
				charactersFound++
				continue
			}

		}

		if len(wordSlice) == charactersFound {
			return WordOccurance{found: true, position: Position{x:pos.x, y:pos.y}, oreantation: direction}
		}
	} else {
		return WordOccurance{found: false}
	}
			
	return WordOccurance{found: false}
}

func getCharacterInDirection(direction int, wordSearch []string, x int, y int, width int, offset int) string {
	switch direction {
	case RIGHT:
		return wordSearch[y * width + (x + offset)]
		
	case BOTTOM_RIGHT:
		return wordSearch[(y + offset) * width + (x + offset)]
		
	case BOTTOM:
		return wordSearch[(y + offset) * width + x]
		
	case BOTTOM_LEFT:
		return wordSearch[(y + offset) * width + (x - offset)]
		
	case LEFT:
		return wordSearch[y * width + (x - offset)]
		
	case TOP_LEFT:
		return wordSearch[(y - offset) * width + (x - offset)]
		
	case TOP:
		return wordSearch[(y - offset) * width + x]
		
	case TOP_RIGHT:
		return wordSearch[(y - offset) * width + (x + offset)]
		
	}
	return ""
}

func isOccuranceFound(occurance WordOccurance, occurancesFound []WordOccurance) bool {
	for i := 0; i < len(occurancesFound); i++ {
		if occurance.position.x == occurancesFound[i].position.x && occurance.position.y == occurancesFound[i].position.y && occurance.oreantation == occurancesFound[i].oreantation {
			return true
		}
	}
	return false
}