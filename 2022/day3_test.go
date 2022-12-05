package aoc2022

import (
	"bufio"
	"strings"
	"testing"
)


var rucksacks string = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`

func TestTotalPrioritizedItemsInRucksacks(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(rucksacks))
	got := TotalPrioritizedItemsInRucksacks(scanner)
	if got != 157 {
		t.Errorf("TotalPrioritizedItemsInRucksacks = %d; expected 157", got)
	}
}

func TestTotalBadgePriorities(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(rucksacks))
	got := TotalBadgePriorities(scanner)
	if got != 70 {
		t.Errorf("TotalBadgePriorities = %d; expected 70", got)
	}
}