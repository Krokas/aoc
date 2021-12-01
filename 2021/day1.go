package aoc2021

import (
	"aoc/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Day1() {
	file, err := os.Open("./2021/input/day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sonarMesurement []int

	for scanner.Scan() {
		mesurement, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		} else {
			sonarMesurement = append(sonarMesurement, mesurement)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

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

	fmt.Printf("Measurement increased ")
	fmt.Printf(string(aocColor.Green) + strconv.Itoa(increaseTimes) + string(aocColor.Reset))
	fmt.Printf(" times. \n")
}
