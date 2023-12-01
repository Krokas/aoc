package aoc2023

import (
	"bufio"
	"strings"
	"testing"
)

var calibrationValues string = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`

var spelledCalibrationValues string = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`

func TestGetCalibrationValuesSum(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(calibrationValues))
	got := GetCalibrationValuesSum(scanner)
	if got != 142 {
		t.Errorf("GetCalibrationValuesSum = %d; expected 142", got)
	}
}

func TestGetSpelledCalibrationValuesSum(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(spelledCalibrationValues))
	got := GetSpelledCalibrationValuesSum(scanner)
	
	if got != 281 {
		t.Errorf("GetSpelledCalibrationValuesSum = %d; expected 281", got)
	}

}