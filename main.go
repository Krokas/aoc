package main

import (
	aoc2021 "aoc/2021"
	aoc2022 "aoc/start/2022"
	"aoc/utils"
	"bufio"
)

func main() {
	// Day 1
	sonar := aoc2021.SonarSweep{
		Input: "./2021/input/day1.txt",
	}
	sonar.CountDepthMesurementIncrease()
	sonar.ThreeMeasurementsWindow()

	// Day 2
	submarine := aoc2021.Submarine{
		Input:      "./2021/input/day2.txt",
		Depth:      0,
		Horizontal: 0,
		Aim:        0,
	}

	submarine.Navigate()
	submarine.BackToStartingPosition()
	submarine.Steer()

	// Day 3
	diagnostics := aoc2021.Diagnostics{
		Input:                 "./2021/input/day3.txt",
		PowerConsumption:      0,
		GammaRate:             0,
		EpsilonRate:           0,
		OxygenGeneratorRating: 0,
		Co2ScrubberRating:     0,
		LifeSupportRating:     0,
	}

	diagnostics.GetPowerConsumption()
	diagnostics.GetLifeSupportRating()

	// 2022
	utils.Warn("THIS IS A START TO 2022!")
	utils.Warn("DAY 1")
	file := utils.OpenFile("./2022/input/day1.txt")
	defer file.Close()
	
	caloryListScanner := bufio.NewScanner(file)
	maxCalories := aoc2022.TotalMostCalories(caloryListScanner)
	utils.WarnWithIntValue("Top Total Calories by elf", maxCalories)

	file2 := utils.OpenFile("./2022/input/day1.txt")
	defer file2.Close()

	topThreecaloryListScanner := bufio.NewScanner(file2)
	topThreeTotalColories := aoc2022.TopThreeTotalMostCalories(topThreecaloryListScanner)
	utils.WarnWithIntValue("Top Three Total Calories by elf", topThreeTotalColories)
	

}
