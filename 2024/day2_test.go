package aoc2024

import (
	"bufio"
	"strings"
	"testing"
)

var reportsSample string = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestGetSafeReportSum(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(reportsSample))
	got := GetSafeReportSum(scanner)
	if got != 2 {
		t.Errorf("GetSafeReportSum = %d, expected 2", got)
	}
}

func TestGetSafeReportWithProblemDampener(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(reportsSample))
	got := GetSafeReportWithProblemDampener(scanner)
	if got != 4 {
		t.Errorf("GetSafeReportWithProblemDampener = %d, expected 4", got)
	}
}