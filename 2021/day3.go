package aoc2021

import (
	"aoc/utils"
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Diagnostics struct {
	Input                 string
	PowerConsumption      int
	GammaRate             int
	EpsilonRate           int
	OxygenGeneratorRating int
	Co2ScrubberRating     int
	LifeSupportRating     int
}

type bitArray []int

func (diagnostics *Diagnostics) GetPowerConsumption() {

	diagnostics.GammaRate = diagnostics.getDecRate(true)
	diagnostics.EpsilonRate = diagnostics.getDecRate(false)
	diagnostics.PowerConsumption = diagnostics.GammaRate * diagnostics.EpsilonRate

	diagnostics.PrintSubmarineDiagnosticStatus("SUBMARINE DIAGNOSTICS")
}

func (diagnostics *Diagnostics) GetLifeSupportRating() {
	diagnostics.getRating(true)
}

func (diagnostics *Diagnostics) getDecRate(getGamma bool) int {

	bitMatrix := diagnostics.parseDiagnostics()

	var binaryString string
	lineCount := len(bitMatrix)
	lineLength := len(bitMatrix[0])
	bitIndex := 0
	lineIndex := 0
	for bitIndex < lineLength {
		bitOn := 0
		bitOff := 0
		for lineIndex < lineCount {
			if bitMatrix[lineIndex][bitIndex] == 0 {
				bitOff++
			} else {
				bitOn++
			}

			lineIndex++
		}

		if getGamma {
			if bitOn > bitOff {
				binaryString = binaryString + "1"
			} else {
				binaryString = binaryString + "0"
			}
		} else {
			if bitOn > bitOff {
				binaryString = binaryString + "0"
			} else {
				binaryString = binaryString + "1"
			}
		}

		bitIndex++
		lineIndex = 0
	}

	rateDec, err := strconv.ParseInt(binaryString, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return int(rateDec)
}

func (diagnostics *Diagnostics) getRating(isOxygenGenerator bool) {
	bitMatrix := diagnostics.parseDiagnostics()

	bitIndexToCompare := 0
	indicesToRemove := getIndicesToRemove(bitMatrix, bitIndexToCompare)

	printDeletableLines(bitMatrix, bitIndexToCompare, indicesToRemove)
	bitMatrix2 := filterBitMatrixByIndices(bitMatrix, indicesToRemove)

	bitIndexToCompare2 := 1
	indicesToRemove2 := getIndicesToRemove(bitMatrix, bitIndexToCompare2)

	printDeletableLines(bitMatrix2, bitIndexToCompare2, indicesToRemove2)
}

func printDeletableLines(bitMatrix []bitArray, bitIndexToCompare int, indicesToRemove []int) {
	for _, indexToRemove := range indicesToRemove {
		fmt.Printf(string(utils.Red) + strconv.Itoa(indexToRemove) + string(utils.Reset) + " ")
	}
	fmt.Printf("\n\n")

	for lineIndex, line := range bitMatrix {
		for bitIndex, bit := range line {
			if bitIndex == bitIndexToCompare {
				if utils.IntContains(indicesToRemove, lineIndex) {
					fmt.Printf(string(utils.Yellow) + strconv.Itoa(bit) + string(utils.Reset))
				} else {
					fmt.Printf(string(utils.Green) + strconv.Itoa(bit) + string(utils.Reset))
				}
			} else {
				if utils.IntContains(indicesToRemove, lineIndex) {
					fmt.Printf(string(utils.Red) + strconv.Itoa(bit) + string(utils.Reset))
				} else {
					fmt.Printf(strconv.Itoa(bit))
				}
			}
		}
		fmt.Printf(" " + string(utils.Yellow) + strconv.Itoa(lineIndex) + string(utils.Reset))
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func filterBitMatrixByIndices(bitMatrix []bitArray, indices []int) []bitArray {
	newMatrix := make([]bitArray, 0)
	for lineIndex, line := range bitMatrix {
		if !utils.IntContains(indices, lineIndex) {
			newMatrix = append(newMatrix, line)
		}
	}

	return newMatrix
}

func getIndicesToRemove(bitMatrix []bitArray, bitIndexToCompare int) []int {
	oneCount := 0
	zeroCount := 0
	zeroToRemove := false
	var indicesToRemove []int
	countLines := 0
	for _, lineToConsider := range bitMatrix {
		if lineToConsider[bitIndexToCompare] == 0 {
			zeroCount++
		} else {
			oneCount++
		}
		countLines++
	}
	fmt.Printf("One count: " + string(utils.Green) + strconv.Itoa(oneCount) + string(utils.Reset) +
		" Zero count: " + string(utils.Green) + strconv.Itoa(zeroCount) + string(utils.Reset) + " lines counted: " + strconv.Itoa(countLines) + "\n\n")
	if oneCount >= zeroCount {
		zeroToRemove = true
	}

	for indexToRemove, lineConsidered := range bitMatrix {
		if zeroToRemove && lineConsidered[bitIndexToCompare] == 0 {
			indicesToRemove = append(indicesToRemove, indexToRemove)
		}
	}

	return indicesToRemove
}

func (diagnostics *Diagnostics) parseDiagnostics() []bitArray {
	file := utils.OpenFile(diagnostics.Input)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// put bits into matrix
	var bitMatrix []bitArray
	for scanner.Scan() {
		line := scanner.Text()
		lineBits := strings.Split(line, "")

		var bitLine bitArray

		for _, bitString := range lineBits {
			if bit, err := strconv.Atoi(bitString); err != nil {
				log.Fatal(err)
			} else {
				bitLine = append(bitLine, bit)
			}
		}
		bitMatrix = append(bitMatrix, bitLine)

	}

	return bitMatrix
}

func (diagnostics *Diagnostics) PrintSubmarineDiagnosticStatus(statusTitle string) {
	fmt.Println(string(utils.Yellow) + "[" + statusTitle + "]:" + string(utils.Reset))
	fmt.Println("Gamma rate: " + string(utils.Green) + strconv.Itoa(diagnostics.GammaRate) + string(utils.Reset))
	fmt.Println("Epsilon rate: " + string(utils.Green) + strconv.Itoa(diagnostics.EpsilonRate) + string(utils.Reset))
	fmt.Println("Power consumption: " + string(utils.Green) + strconv.Itoa(diagnostics.PowerConsumption) + string(utils.Reset))
	fmt.Printf("\n")
}
