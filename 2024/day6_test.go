package aoc2024

import (
	"bufio"
	"strings"
	"testing"
)

var areaMapSample string = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestGetGuardVisitedAreaCount(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(areaMapSample))
	got := GetGuardVisitedAreaCount(scanner)
	if got != 41 {
		t.Errorf("GetGuardVisitedAreaCount = %d, expected 41", got)
	}
}