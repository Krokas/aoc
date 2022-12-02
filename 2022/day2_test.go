package aoc2022

import (
	"bufio"
	"strings"
	"testing"
)

var rounds string = `A Y
B X
C Z
`

func TestCheatingTotalScore(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(rounds))
	got := CheatingTotalScore(scanner)
	if got != 15 {
		t.Errorf("CheatingTotalScore = %d; expected 15", got)
	}
}

func TestUltraTopSecretGuideForCheating(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(rounds))
	got := UltraTopSecretGuideForCheating(scanner)
	if got != 12 {
		t.Errorf("UltraTopSecretGuideForCheating = %d; expected 12", got)
	}
}