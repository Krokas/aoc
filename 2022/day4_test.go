package aoc2022

import (
	"bufio"
	"strings"
	"testing"
)

var pairs string = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func TestTotalFullyContainedPairs(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(pairs))
	got := TotalFullyContainedPairs(scanner)
	if got != 2 {
		t.Errorf("TotalFullyContainedPairs = %d; expected 2", got)
	}
}

func TestTotalOverlapingPairs(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(pairs))
	got := TotalOverlapingPairs(scanner)
	if got != 4 {
		t.Errorf("TotalOverlapingPairs = %d; expected 4", got)
	}
}