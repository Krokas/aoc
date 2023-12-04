package aoc2023

import (
	"bufio"
	"strconv"
	"strings"
)

type ColoredCubeCount struct {
	red int;
	green int;
	blue int;
}

func GetSumOfGameIds(scanner *bufio.Scanner) int {
	var correctGameIds []int
	for scanner.Scan() {
		line := scanner.Text()
		id := getGameId(line)

		maxCubes := getMaxQubesPerGame(line)
		
		if isGamePossible(maxCubes) {
			correctGameIds = append(correctGameIds, id)
		}
		

	}

	var sum int;
	for _, gameId := range correctGameIds {
		sum = sum + gameId
	}

	return sum;
}

func GetPowerOfGamesSum(scanner *bufio.Scanner) int {
	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		
		maxCubes := getMaxQubesPerGame(line)

		power := maxCubes.red * maxCubes.green * maxCubes.blue

		sum = sum + power
	}

	return sum
}

func getGameId(line string) int {
	var gameId int = -1;
	var err error
	gameIdentifier := strings.Split(line, ":")
	if len(gameIdentifier) > 1 {
		identifierArray := strings.Split(gameIdentifier[0], " ")
		if len(identifierArray) > 1 {
			gameId, err = strconv.Atoi(identifierArray[1])

			if err != nil {
				return -1
			}
		}
	}

	return gameId
}

func getMaxQubesPerGame(line string) ColoredCubeCount {
	var maxCubes ColoredCubeCount
	maxCubes.red = 0
	maxCubes.blue = 0
	maxCubes.green = 0

	gameContentArray := strings.Split(line, ": ")

	if len(gameContentArray) > 1 {
		sets := strings.Split(gameContentArray[1], "; ")

		for _, set := range sets {
			qubeData := strings.Split(set, ", ")
			for _, quad := range qubeData {
				quadDataArray := strings.Split(quad, " ")
				
				if strings.Contains(quadDataArray[1], "red") {
					count, err := strconv.Atoi(quadDataArray[0])

					if err == nil && count > maxCubes.red {
						maxCubes.red = count
					}
				}

				if strings.Contains(quadDataArray[1], "blue") {
					count, err := strconv.Atoi(quadDataArray[0])

					if err == nil && count > maxCubes.blue {
						maxCubes.blue = count
					}
				}

				if strings.Contains(quadDataArray[1], "green") {
					count, err := strconv.Atoi(quadDataArray[0])

					if err == nil && count > maxCubes.green {
						maxCubes.green = count
					}
				}

			}
		}
	}

	return maxCubes
}

func isGamePossible(log ColoredCubeCount) bool {
		// 12 red cubes, 13 green cubes, and 14 blue cubes

		var correctCombination ColoredCubeCount
		correctCombination.red = 12
		correctCombination.green = 13
		correctCombination.blue = 14

		var isCorrect bool = true

		if correctCombination.red < log.red {
			isCorrect = false
		}

		if correctCombination.green < log.green {
			isCorrect = false
		}

		if correctCombination.blue < log.blue {
			isCorrect = false
		}

		return isCorrect
}