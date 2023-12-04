package aoc2023

import (
	"bufio"
	"strings"
	"testing"
)

var gameLog string = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`

var powerGameLog string = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`

func TestGetSumOfGameIds(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(gameLog))
	got := GetSumOfGameIds(scanner)
	if got != 8 {
		t.Errorf("GetSumOfGameIds = %d; expected 8", got)
	}
}

func TestGetPowerOfGamesSum(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(powerGameLog))
	got := GetPowerOfGamesSum(scanner)
	if got != 2286 {
		t.Errorf("GetPowerOfGamesSum = %d; expected 2286", got)
	}
}