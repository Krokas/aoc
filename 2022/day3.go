package aoc2022

import (
	"bufio"
	"strings"
)

func getCompartments(rucksack string) (string, string) {
	compartmentSize := len(rucksack) / 2
	firstCompartment := rucksack[:compartmentSize]
	secondCompartment := rucksack[compartmentSize:]
	return firstCompartment, secondCompartment
}

func getMatchingItemTypes(first string, second string) string {
	var matches string
	for _, char := range first {
		if strings.Contains(second, string(char)) && !strings.Contains(matches, string(char)) {
			matches += string(char)
		}
	}
	return matches
}

func getPriorityByType(itemType string) int {
	var lowerCase string
	for i := 'a'; i <= 'z'; i++ {
        lowerCase += string(i)
    }
	var upperCase string
	for i := 'A'; i <= 'Z'; i++ {
		upperCase += string(i)
	}
	for pos, char := range lowerCase {
		if itemType == string(char) {
			return pos + 1
		}
	}

	for pos, char := range upperCase {
		if itemType == string(char) {
			return pos + 1 + len(lowerCase)
		}
	}
	return 0
} 

func TotalPrioritizedItemsInRucksacks(scanner *bufio.Scanner) int {
	totalPriority := 0
	for scanner.Scan() {
		line := scanner.Text()
		firstCompartment, secondCompartment := getCompartments(line)
		matchingTypes := getMatchingItemTypes(firstCompartment, secondCompartment)
		for _, char := range matchingTypes {
			totalPriority += getPriorityByType(string(char))
		}
	}
	return totalPriority
}

func TotalBadgePriorities(scanner *bufio.Scanner) int {
	TotalBadgePriorities := 0
	elfIndex := 0;
	var elfOneInventory string
	var elfTwoInventory string
	var elfThreeInventory string
	for scanner.Scan() {
		elfIndex++
		switch elfIndex {
		case 1:
			elfOneInventory = scanner.Text()
			break
		case 2:
			elfTwoInventory = scanner.Text()
			break
		case 3:
			elfThreeInventory = scanner.Text()
			break
		}
		if elfIndex == 3 {
			elfIndex = 0
			oneAndTwoMatches := getMatchingItemTypes(elfOneInventory, elfTwoInventory)
			matchThree := getMatchingItemTypes(oneAndTwoMatches, elfThreeInventory)
			for _, char := range matchThree {
				TotalBadgePriorities += getPriorityByType(string(char))
			}
		}
	}
	return TotalBadgePriorities
}