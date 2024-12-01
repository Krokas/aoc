package aoc2024

import (
	"bufio"
	"strings"
	"testing"
)

var listSample string = `3   4
4   3
2   5
1   3
3   9
3   3
`

func TestGetDistanceSum(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(listSample))
	got := GetDistanceSum(scanner)
	if got != 11 {
		t.Errorf("GetDistanceSum = %d, expected 11", got)
	}
}