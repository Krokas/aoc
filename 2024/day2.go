package aoc2024

import (
	"bufio"
	"strconv"
	"strings"
)

type Report struct {
	levels []int
	isSafe bool
}

func GetSafeReportSum(scanner *bufio.Scanner) int {
	var reports []Report
	for scanner.Scan() {
		line := scanner.Text()
		report := parseReport(line)
		report.isSafe = isReportSafe(report)
		reports = append(reports, report)
	}

	safeReportCount := 0
	for i := 0; i < len(reports); i++ {
		if reports[i].isSafe {
			safeReportCount++
		}
	}

	return safeReportCount
}

func parseReport(report string) Report {
	var levels []int
	chars := strings.Split(report, " ")
	for _, char := range chars {
		number, err := strconv.Atoi(strings.Trim(char, " "))

		if err == nil {
			levels = append(levels, number)
		}
	}
	
	return Report{levels: levels, isSafe: false}
}

func isReportSafe(report Report) bool {
	isSafe := true
	isIcreasing := true
	isDirectionSet := false
	if len(report.levels) == 0 {
		return false
	}
	for i := 0; i < len(report.levels); i++ {
		if i > 0 {
			difference := report.levels[i] - report.levels[i - 1]

			if isDirectionSet {
				if isIcreasing && difference > 0 {
					return false
				}

				if !isIcreasing && difference < 0 {
					return false
				}
			}

			if difference < 0 {
				isIcreasing = true
				isDirectionSet = true
				difference *= -1
			} else {
				isIcreasing = false
				isDirectionSet = true
			}


			isSafe = difference > 0 && difference < 4
		}

		if !isSafe {
			return false
		}
	}
	return isSafe
}