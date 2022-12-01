package aoc2022

import (
	"bufio"
	"strings"
	"testing"
)

var calories string = `1000
	2000
	3000
	
	4000
	
	5000
	6000
	
	7000
	8000
	9000
	
	10000
	`

func TestTotalMostCalories(t *testing.T) {	
	scanner := bufio.NewScanner(strings.NewReader(calories))
	got := TotalMostCalories(scanner)
	if got != 24000 {
		t.Errorf("TestTotalMostCalories = %d; expected 24000", got)
	}
}

func TestTopThreeTotalMostCalories(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(calories))
	got := TopThreeTotalMostCalories(scanner)
	if got != 45000 {
		t.Errorf("TopThreeTotalMostCalories = %d; expected 45000", got)
	}
}