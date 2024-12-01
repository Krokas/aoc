package aoc2024

import (
	"bufio"
	"sort"
	"strconv"
	"strings"
)

type TwoDirectionLists struct {
	listOne []int
	listTwo []int
}

type Directions struct {
	one int
	two int
}

func GetDistanceSum(scanner *bufio.Scanner) int {
	var directionLists TwoDirectionLists;
	directionDistanceSum := 0;
	for scanner.Scan() {
		line := scanner.Text()
		directions := parseDirections(line)
		directionLists.listOne = append(directionLists.listOne, directions.one)
		directionLists.listTwo = append(directionLists.listTwo, directions.two)
	}

	sort.Slice(directionLists.listOne, func(i, j int) bool {
		return directionLists.listOne[i] < directionLists.listOne[j]
	})

	sort.Slice(directionLists.listTwo, func(i, j int) bool {
		return directionLists.listTwo[i] < directionLists.listTwo[j]
	})
	
	for i := 0; i < len(directionLists.listOne); i++ {
		diff := directionLists.listOne[i] - directionLists.listTwo[i]

		if (diff < 0) {
			diff *= -1
		}

		directionDistanceSum += diff
	}

	return directionDistanceSum
}

func parseDirections(line string) Directions {
	chars := strings.Split(line, " ")
	var numbers []int;
	for _, char := range chars {
		number, err := strconv.Atoi(strings.Trim(char, " "))

		if err == nil {
			numbers = append(numbers, number)
		}
	}

	if len(numbers) > 0 {
		return Directions{one: numbers[0], two: numbers[len(numbers) - 1]}
	}
	return Directions{one: 0, two: 0}
}