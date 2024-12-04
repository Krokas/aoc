package aoc2024

import (
	"bufio"
	"strings"
	"testing"
)

var wordSearchSample string = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestGetSumOfWordOccuranceInWordSearch(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(wordSearchSample))
	got := GetSumOfWordOccuranceInWordSearch(scanner)
	if got != 18 {
		t.Errorf("GetSumOfWordOccuranceInWordSearch = %d, expected 18", got)
	}
}