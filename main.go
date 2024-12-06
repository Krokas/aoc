package main

import (
	aoc2021 "aoc/2021"
	aoc2022 "aoc/start/2022"
	aoc2023 "aoc/start/2023"
	aoc2024 "aoc/start/2024"
	"aoc/utils"
	"bufio"
)

func run2021() {
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
}

func run2022() {
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
	
	utils.Warn("DAY 2")
	day2Part1File := utils.OpenFile("./2022/input/day2.txt")
	defer day2Part1File.Close()

	RPCScanner := bufio.NewScanner(day2Part1File)
	totalRPCScore := aoc2022.CheatingTotalScore(RPCScanner)
	utils.WarnWithIntValue("Total Rock Paper Scissors Score", totalRPCScore)

	day2Part2File := utils.OpenFile("./2022/input/day2.txt")
	defer day2Part2File.Close()

	RPCFullStrategyScanner := bufio.NewScanner(day2Part2File)
	totalRPCFullStrategyScore := aoc2022.UltraTopSecretGuideForCheating(RPCFullStrategyScanner)
	utils.WarnWithIntValue("Total Rock Paper Scissors: Full Strategy Score", totalRPCFullStrategyScore)

	utils.Warn("DAY 3")
	day3Part1File := utils.OpenFile("./2022/input/day3.txt")
	defer day3Part1File.Close()

	day3part1Scanner := bufio.NewScanner(day3Part1File)
	totalPriorityInRucksacks := aoc2022.TotalPrioritizedItemsInRucksacks(day3part1Scanner)
	utils.WarnWithIntValue("Total Priority of Items by Type in rucksacks", totalPriorityInRucksacks)

	day3Part2File := utils.OpenFile("./2022/input/day3.txt")
	defer day3Part2File.Close()

	day3Part2Scanner := bufio.NewScanner(day3Part2File)
	totalBadgePriority := aoc2022.TotalBadgePriorities(day3Part2Scanner)
	utils.WarnWithIntValue("Total Priority of Badge Items", totalBadgePriority)

	utils.Warn("DAY 4")
	day4Part1File := utils.OpenFile("./2022/input/day4.txt")
	defer day4Part1File.Close()

	day4Part1Scanner := bufio.NewScanner(day4Part1File)
	totalFullyContainerdPairs := aoc2022.TotalFullyContainedPairs(day4Part1Scanner)
	utils.WarnWithIntValue("Total Fully contained Payrs", totalFullyContainerdPairs)

	day4Part2File := utils.OpenFile("./2022/input/day4.txt")
	defer day4Part2File.Close()

	day4Part2Scanner := bufio.NewScanner(day4Part2File)
	totalOverlapingPairs := aoc2022.TotalOverlapingPairs(day4Part2Scanner)
	utils.WarnWithIntValue("Total overlaping pairs",  totalOverlapingPairs)
}

func run2023() {
	day1Part1File := utils.OpenFile("./2023/input/day1part1.txt")
	defer day1Part1File.Close()

	day1Part1Scanner := bufio.NewScanner(day1Part1File)
	TrebushetCalibrationValuesSum := aoc2023.GetCalibrationValuesSum(day1Part1Scanner)
	utils.WarnWithIntValue("Trebushet Calibration Value", TrebushetCalibrationValuesSum)

	day1Part2File := utils.OpenFile("./2023/input/day1part2.txt")
	defer day1Part2File.Close()

	day1Part2Scanner := bufio.NewScanner(day1Part2File)
	TrebushetSpelledCalibrationValuesSum := aoc2023.GetSpelledCalibrationValuesSum(day1Part2Scanner)
	utils.WarnWithIntValue("Trebushet Spelled Calibration Value", TrebushetSpelledCalibrationValuesSum)

	day2Part1File := utils.OpenFile("./2023/input/day2part1.txt")
	defer day2Part1File.Close()

	day2part1Scanner := bufio.NewScanner(day2Part1File)
	coloredCubeSum := aoc2023.GetSumOfGameIds(day2part1Scanner)
	utils.WarnWithIntValue("Valid game count", coloredCubeSum)

	day2Part2File := utils.OpenFile("./2023/input/day2part2.txt")
	defer day2Part2File.Close()

	day2Part2Scanner := bufio.NewScanner(day2Part2File)
	sumOfPowerGames := aoc2023.GetPowerOfGamesSum(day2Part2Scanner)
	utils.WarnWithIntValue("Sum of power of games", sumOfPowerGames)
}

func run2024() {
	day1Part1File := utils.OpenFile("./2024/input/day1part1.txt")
	defer day1Part1File.Close()

	day1Part1Scanner := bufio.NewScanner(day1Part1File)
	totalDistance := aoc2024.GetDistanceSum(day1Part1Scanner)
	utils.WarnWithIntValue("Total distance between list one and list two", totalDistance)

	day1Part2File := utils.OpenFile("./2024/input/day1part2.txt")
	defer day1Part2File.Close()

	day1Part2Scanner := bufio.NewScanner(day1Part2File)
	similarityScore := aoc2024.GetSimilarityScore(day1Part2Scanner)
	utils.WarnWithIntValue("Similarity score", similarityScore)

	day2Part1File := utils.OpenFile("./2024/input/day2part1.txt")
	defer day2Part1File.Close()
	day2part1Scanner := bufio.NewScanner(day2Part1File)
	safeReportCount := aoc2024.GetSafeReportSum(day2part1Scanner)
	utils.WarnWithIntValue("Safe Report Count", safeReportCount)

	day2Part2File := utils.OpenFile("./2024/input/day2part2.txt")
	defer day2Part2File.Close()
	day2Part2Scanner := bufio.NewScanner(day2Part2File)
	safeReportWithProblemDampener := aoc2024.GetSafeReportWithProblemDampener(day2Part2Scanner)
	utils.WarnWithIntValue("Safe Report with Problem Dampener", safeReportWithProblemDampener)

	day3Part1File := utils.OpenFile("./2024/input/day3part1.txt")
	defer day3Part1File.Close()
	day3Part1Scanner := bufio.NewScanner(day3Part1File)
	corruptedMemorySum := aoc2024.GetSumOfMultipiedCorruptedMemory(day3Part1Scanner)
	utils.WarnWithIntValue("Corrupted Memory Sum", corruptedMemorySum)

	day3Part2File := utils.OpenFile("./2024/input/day3part2.txt")
	defer day3Part2File.Close()
	day3Part2Scanner := bufio.NewScanner(day3Part2File)
	corruptedMemoryDisableableSum := aoc2024.GetSumofDisableableMultipieldCorruptedMemory(day3Part2Scanner)
	utils.WarnWithIntValue("Disableable Corrupted Memory Sum", corruptedMemoryDisableableSum)

	day4part1File := utils.OpenFile("./2024/input/day4part1.txt")
	defer day4part1File.Close()
	day4Part1Scanner := bufio.NewScanner(day4part1File)
	SumWordSearch := aoc2024.GetSumOfWordOccuranceInWordSearch(day4Part1Scanner)
	utils.WarnWithIntValue("Word search sum", SumWordSearch)

	day4Part2File := utils.OpenFile("./2024/input/day4part2.txt")
	defer day4Part2File.Close()
	day4Part2Scanner := bufio.NewScanner(day4Part2File)
	XMASCount := aoc2024.GetSumOfXCrossedWords(day4Part2Scanner)
	utils.WarnWithIntValue("X-MAS count", XMASCount)

	day5Part1File := utils.OpenFile("./2024/input/day5part1.txt")
	defer day5Part1File.Close()
	day5Part1Scanner := bufio.NewScanner(day5Part1File)
	MiddlePageNumberSum := aoc2024.GetSumOfMiddlePageNumbers(day5Part1Scanner)
	utils.WarnWithIntValue("Sum of Middle page number", MiddlePageNumberSum)

	day5Part2File := utils.OpenFile("./2024/input/day5part2.txt")
	defer day5Part2File.Close()
	day5Part2Scanner := bufio.NewScanner(day5Part2File)
	MiddlePageNumberIncorrectOrder := aoc2024.GetSumOfIncorrectlyOrderedUpdateMiddlePages(day5Part2Scanner)
	utils.WarnWithIntValue("Incorrectly ordered pages middle page number", MiddlePageNumberIncorrectOrder)

	day6Part1File := utils.OpenFile("./2024/input/day6part1.txt")
	defer day6Part1File.Close()
	day6Part1Scanner := bufio.NewScanner(day6Part1File)
	PlacesVisitedByGuard := aoc2024.GetGuardVisitedAreaCount(day6Part1Scanner)
	utils.WarnWithIntValue("Guard visited number of areas before leaving", PlacesVisitedByGuard)
}

func main() {
	run2024()
}

func fileToScanner(filepath string) *bufio.Scanner {
	file := utils.OpenFile(filepath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	return scanner
}
