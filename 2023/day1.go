package aoc2023

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
)

func GetCalibrationValuesSum(scanner *bufio.Scanner) int {

	var sum int;

	for scanner.Scan() {
		line := scanner.Text()
		numbers := getNumericFromLine(line)

		
		first := numbers[0]
		last := numbers[len(numbers)-1]

		var buffer bytes.Buffer;
		buffer.WriteString(strconv.Itoa(first))
		buffer.WriteString(strconv.Itoa(last))

		calibrationValue, err := strconv.Atoi(buffer.String())

		if err == nil {
			sum = sum + calibrationValue
		}
	}
 	return sum;
}

func GetSpelledCalibrationValuesSum(scanner *bufio.Scanner) int {
	var sum int;
	var numbers []int;
	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < len(line); i++ {
			// utils.Warn(line[i:])
			number, err := strconv.Atoi(string(line[i]))

			if err == nil {
				numbers = append(numbers, number)
			} else {
				spelledNumber := getSpelledNumber(line[i:])
				if (spelledNumber > 0) {
					numbers = append(numbers, spelledNumber)
				}
			}
		}

		first := numbers[0]
		last := numbers[len(numbers)-1]

		var buffer bytes.Buffer;
		buffer.WriteString(strconv.Itoa(first))
		buffer.WriteString(strconv.Itoa(last))

		calibrationValue, err := strconv.Atoi(buffer.String())

		if err == nil {
			sum = sum + calibrationValue
		}

		numbers = numbers[:0]
	}
	return sum
}

func getNumericFromLine(line string) []int {
	chars := strings.Split(line, "")
	var numbers []int;
	for _, char := range chars {
		number, err := strconv.Atoi(char)

		if err == nil {
			numbers = append(numbers, number)
		}
	}

	return numbers;
}

func getSpelledNumber(line string) int {
	if strings.Contains(line, "one") {
		index := strings.Index(line, "one")
		
		if index == 0 {
			return 1
		}
	}
	
	if strings.Contains(line, "two") {
		index := strings.Index(line, "two")
		
		if index == 0 {
			return 2
		}
	}

	if strings.Contains(line, "three") {
		index := strings.Index(line, "three")
		
		if index == 0 {
			return 3
		}
	}

	if strings.Contains(line, "four") {
		index := strings.Index(line, "four")
		
		if index == 0 {
			return 4
		}
	}

	if strings.Contains(line, "five") {
		index := strings.Index(line, "five")
		
		if index == 0 {
			return 5
		}
	}

	if strings.Contains(line, "six") {
		index := strings.Index(line, "six")
		
		if index == 0 {
			return 6
		}
	}

	if strings.Contains(line, "seven") {
		index := strings.Index(line, "seven")
		
		if index == 0 {
			return 7
		}
	}

	if strings.Contains(line, "eight") {
		index := strings.Index(line, "eight")
		
		if index == 0 {
			return 8
		}
	}

	if strings.Contains(line, "nine") {
		index := strings.Index(line, "nine")
		
		if index == 0 {
			return 9
		}
	}

	return -1
}