package aoc2022

import (
	"aoc/utils"
	"bufio"
	"strconv"
	"strings"
)

func getCaloriesByElf(scanner *bufio.Scanner) []int {
	var TotalCaloriesArray []int
	elfTotalCalories := 0
	var previousLenEqZero bool = true
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "\f\t\r\n ")
		if len(line) > 0 {
			previousLenEqZero = false
			itemCalories, _ := strconv.Atoi(line)
			elfTotalCalories += itemCalories
		} else {
			if previousLenEqZero {
				continue
			}
			previousLenEqZero = true
			TotalCaloriesArray = append(TotalCaloriesArray, elfTotalCalories)
			elfTotalCalories = 0
		}
	}
	return TotalCaloriesArray
}

func TotalMostCalories(scanner *bufio.Scanner) int {
	TotalCaloriesArray := getCaloriesByElf(scanner)
	_, max := utils.MinMax(TotalCaloriesArray)
	return max
}

func TopThreeTotalMostCalories(scanner *bufio.Scanner) int {
	TotalCaloriesByElf := getCaloriesByElf(scanner)
	elfTotalCalories := 0
	for i := 1; i <= 3; i++ {
		_, max := utils.MinMax(TotalCaloriesByElf)
		TotalCaloriesByElf = utils.RemoveFromArray(TotalCaloriesByElf, max)
		elfTotalCalories += max
	}
	return elfTotalCalories
}
