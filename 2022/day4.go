package aoc2022

import (
	"bufio"
	"strconv"
	"strings"
)

type Pair struct {
	start int;
	end int;
}

func getRangePairs(elfPair string) (Pair, Pair) {
	var start, end Pair

	pair := strings.Split(elfPair, ",")

	firstElf := pair[0]
	secondElf := pair[1]

	firstElfSplice := strings.Split(firstElf, "-")
	secondElfSplice := strings.Split(secondElf, "-")


	start.start, _ = strconv.Atoi(firstElfSplice[0])
	start.end, _ = strconv.Atoi(firstElfSplice[1])

	end.start, _ = strconv.Atoi(secondElfSplice[0])
	end.end, _ = strconv.Atoi(secondElfSplice[1])

	return start, end
}

func TotalFullyContainedPairs(scanner *bufio.Scanner) int {
	var TotalFullyContainedPairs int = 0
	for scanner.Scan() {
		line := scanner.Text()
		firstElfSectors, secondElfSectors := getRangePairs(line)
		// First contains the second
		if firstElfSectors.start <= secondElfSectors.start && firstElfSectors.end >= secondElfSectors.end {
			TotalFullyContainedPairs++
		// Second contains the first
		} else if secondElfSectors.start <= firstElfSectors.start && secondElfSectors.end >= firstElfSectors.end {
			TotalFullyContainedPairs++
		} 
	}
	return TotalFullyContainedPairs
}

func TotalOverlapingPairs(scanner *bufio.Scanner) int {
	var TotalOverlapingPairs int = 0

	for scanner.Scan() {
		line := scanner.Text()
		firstElfSectors, secondElfSectors := getRangePairs(line)
		// Second overlaps in the first
		if firstElfSectors.start <= secondElfSectors.start && firstElfSectors.end >= secondElfSectors.start {
			TotalOverlapingPairs++
		} else if firstElfSectors.start <= secondElfSectors.end && firstElfSectors.end >= secondElfSectors.end {
			TotalOverlapingPairs++
		// Second overlaps in second
		} else if secondElfSectors.start <= firstElfSectors.start && secondElfSectors.end >= firstElfSectors.start {
			TotalOverlapingPairs++
		} else if secondElfSectors.start <= firstElfSectors.end && secondElfSectors.end >= firstElfSectors.end {
			TotalOverlapingPairs++
		}
	}

	return TotalOverlapingPairs
}