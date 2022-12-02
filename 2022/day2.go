package aoc2022

import (
	"bufio"
	"fmt"
	"strings"
)

type RoundOutcome int

const (
	WIN RoundOutcome = 6
	DRAW RoundOutcome = 3
	LOSS RoundOutcome = 0
)

type Shape string

const (
	OPPONENT_ROCK Shape = "A"
	OPPONENT_PAPER Shape = "B"
	OPPONENT_SCISSORS Shape = "C"

	YOUR_ROCK Shape = "X"
	YOUR_PAPER Shape = "Y"
	YOUR_SCISSORS Shape = "Z"
)

const (
	SHOULD_WIN string = "Z"
	SHOULD_LOSE string = "X"
	SHOULD_DRAW string = "Y"
)

func getShapesOfRound(line string) (string, string) {
	shapes := strings.Split(line, " ")
	if len(shapes) != 2 {
		fmt.Errorf("getShapesOfRound - shape is the size of %d", len(shapes))
		fmt.Errorf("getShapesOfRound - shape in question: %s", shapes)
	}
	return shapes[0], shapes[1]
}

func getRoundOutcome(opponent string, you string) RoundOutcome {
	var outcome RoundOutcome
	switch opponent {
	case string(OPPONENT_ROCK):
		switch you {
		case string(YOUR_PAPER):
			outcome = WIN
		case string(YOUR_ROCK):
			outcome = DRAW
		case string(YOUR_SCISSORS):
			outcome = LOSS
		}
	case string(OPPONENT_PAPER):
		switch you {
		case string(YOUR_PAPER):
			outcome = DRAW
		case string(YOUR_ROCK):
			outcome = LOSS
		case string(YOUR_SCISSORS):
			outcome = WIN
		}
	case string(OPPONENT_SCISSORS):
		switch you {
		case string(YOUR_PAPER):
			outcome = LOSS
		case string(YOUR_ROCK):
			outcome = WIN
		case string(YOUR_SCISSORS):
			outcome = DRAW
		}
	}

	return outcome
}

func getPointsForShape(you string) int {
	var points int
	switch you {
	case string(YOUR_ROCK):
		points = 1
	case string(YOUR_PAPER):
		points = 2
	case string(YOUR_SCISSORS):
		points = 3
	}
	return points
}

func getYourFullStrategyShape(opponent string, you string) Shape {
	var yourShape Shape
	switch opponent {
	case string(OPPONENT_ROCK):
		switch you {
		case SHOULD_WIN:
			yourShape = YOUR_PAPER
		case SHOULD_DRAW:
			yourShape = YOUR_ROCK
		case SHOULD_LOSE:
			yourShape = YOUR_SCISSORS
		}
	case string(OPPONENT_PAPER):
		switch you {
		case SHOULD_WIN:
			yourShape = YOUR_SCISSORS
		case SHOULD_DRAW:
			yourShape = YOUR_PAPER
		case SHOULD_LOSE:
			yourShape = YOUR_ROCK
		}
	case string(OPPONENT_SCISSORS):
		switch you {
		case SHOULD_WIN:
			yourShape = YOUR_ROCK
		case SHOULD_DRAW:
			yourShape = YOUR_SCISSORS
		case SHOULD_LOSE:
			yourShape = YOUR_PAPER
		}
	}
	return yourShape
}

func getScoreByRound(line string) int {
	opponent, you := getShapesOfRound(line)
	outcome := getRoundOutcome(opponent, you)
	shapePoints := getPointsForShape(you)
	return int(outcome) + shapePoints
}

func getFullStrategyScoreByRound(line string) int {
	opponent, you := getShapesOfRound(line)
	// Determine what you should play
	yourShape := getYourFullStrategyShape(opponent, you)
	shapePoints := getPointsForShape(string(yourShape))
	// calculate outcome and points
	var outcome int
	switch you {
	case SHOULD_WIN:
		outcome = int(WIN)
	case SHOULD_DRAW:
		outcome = int(DRAW)
	case SHOULD_LOSE:
		outcome = int(LOSS)
	}
	return shapePoints + outcome
}

func CheatingTotalScore(scanner *bufio.Scanner) int {
	var totalScore int 
	for scanner.Scan() {
		line := scanner.Text()
		gameScore := getScoreByRound(line)
		totalScore += gameScore
	} 
	return totalScore
}

func UltraTopSecretGuideForCheating(scanner *bufio.Scanner) int {
	var totalScore int 
	for scanner.Scan() {
		line := scanner.Text()
		roundScore := getFullStrategyScoreByRound(line)
		totalScore += roundScore
	}
	return totalScore
}