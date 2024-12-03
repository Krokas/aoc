package aoc2024

import (
	"bufio"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type MultiplicationPair struct {
	first int
	second int
}

func GetSumOfMultipiedCorruptedMemory(scanner *bufio.Scanner) int {
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		pairs := getPairs(line)
	
		for i := 0; i < len(pairs); i++ {
			sum += pairs[i].first * pairs[i].second
		}
	
	}
	return sum
}

func getPairs(line string) []MultiplicationPair {
	var pairs []MultiplicationPair
	regex, err := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
	if err != nil {
		log.Fatalf("regexp.Compile: %s", err)
	}

	found := regex.FindAllString(line, -1)
	if len(found) > 0 {
		for i := 0; i < len(found); i++ {
			intermediate := strings.Split(found[i], "(")
			intermediate = strings.Split(intermediate[1], ")")
			pair := strings.Split(intermediate[0], ",")
			if len(pair) == 2 {
				var multimultiplicationPair MultiplicationPair
				number, err := strconv.Atoi(pair[0])

				if err == nil {
					multimultiplicationPair.first = number
				}

				number, err = strconv.Atoi(pair[1])

				if err == nil {
					multimultiplicationPair.second = number
				}

				pairs = append(pairs, multimultiplicationPair)
			}
		}
	}
	return pairs
}