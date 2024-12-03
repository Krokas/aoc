package aoc2024

import (
	"bufio"
	"strings"
	"testing"
)

var corruptedMemorySample string = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

var disableableCorruptedMemorySample string = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func TestGetSumOfMultipiedCorruptedMemory(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(corruptedMemorySample))
	got := GetSumOfMultipiedCorruptedMemory(scanner)
	if got != 161 {
		t.Errorf("GetSumOfMultipiedCorruptedMemory = %d, expected 161", got)
	}
}

func TestGetSumofDisableableMultipieldCorruptedMemory(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(disableableCorruptedMemorySample))
	got := GetSumofDisableableMultipieldCorruptedMemory(scanner)
	if got != 48 {
		t.Errorf("GetSumofDisableableMultipieldCorruptedMemory = %d, expected 48", got)
	}
}