package aoc2024

import (
	"bufio"
	"strings"
	"testing"
)

var corruptedMemorySample string = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

func TestGetSumOfMultipiedCorruptedMemory(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(corruptedMemorySample))
	got := GetSumOfMultipiedCorruptedMemory(scanner)
	if got != 161 {
		t.Errorf("GetSumOfMultipiedCorruptedMemory = %d, expected 161", got)
	}
}