package aoc2021

import (
	"aoc/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type SonarSweep struct {
	Input string
}

func parseMesurements(file *os.File) []int {
	var sonarMesurements []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		mesurement, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		} else {
			sonarMesurements = append(sonarMesurements, mesurement)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sonarMesurements
}

func printMesurementStatus(methodName string, increaseCount int) {
	fmt.Println(string(utils.Yellow) + "[" + methodName + "]:" + string(utils.Reset))
	fmt.Printf("Measurement increased ")
	fmt.Printf(string(utils.Green) + strconv.Itoa(increaseCount) + string(utils.Reset))
	fmt.Printf(" times. \n\n")
}

func (sonar SonarSweep) CountDepthMesurementIncrease() int {
	file := utils.OpenFile(sonar.Input)
	defer file.Close()

	sonarMesurements := parseMesurements(file)

	mesurementIncrease := 0
	var previousMesurement int

	for index := range sonarMesurements {
		if index != 0 && sonarMesurements[index] > previousMesurement {
			mesurementIncrease++
		}
		previousMesurement = sonarMesurements[index]
	}

	printMesurementStatus("DEPTH MESUREMENT INCREASE", mesurementIncrease)
	return mesurementIncrease
}

func (sonar SonarSweep) ThreeMeasurementsWindow() {
	file := utils.OpenFile(sonar.Input)
	defer file.Close()

	sonarMesurement := parseMesurements(file)

	previousSum := 0
	var increaseTimes int

	for index := range sonarMesurement {
		secondIndex := index + 1
		thridIndex := index + 2

		if secondIndex < len(sonarMesurement) && thridIndex < len(sonarMesurement) {
			sum := sonarMesurement[index] + sonarMesurement[secondIndex] + sonarMesurement[thridIndex]

			if previousSum != 0 && sum > previousSum {
				increaseTimes++
			}

			previousSum = sum

		}
	}

	printMesurementStatus("THREE SONAR MESUREMENTS WINDOW", increaseTimes)
}
