package aoc2021

import (
	"aoc/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Submarine struct {
	Input      string
	Horizontal int
	Depth      int
	Aim        int
}

type Instruction struct {
	Type    string
	Advance int
}

func (submarine *Submarine) Forward(advance int) {
	submarine.Horizontal = advance + submarine.Horizontal
}

func (submarine *Submarine) Down(advance int) {
	submarine.Depth = submarine.Depth + advance
}

func (submarine *Submarine) Up(advance int) {
	submarine.Depth = submarine.Depth - advance
}

func (submarine *Submarine) FaceDown(advance int) {
	submarine.Aim = submarine.Aim + advance
}

func (submarine *Submarine) FaceUp(advance int) {
	submarine.Aim = submarine.Aim - advance
}

func (submarine *Submarine) MoveForward(advance int) {
	submarine.Horizontal = submarine.Horizontal + advance
	submarine.Depth = submarine.Depth + (submarine.Aim * advance)
}

func parseInstructions(file *os.File) []Instruction {
	scanner := bufio.NewScanner(file)
	var instructions []Instruction

	for scanner.Scan() {
		instructionArray := strings.Split(scanner.Text(), " ")
		advance, err := strconv.Atoi(instructionArray[1])
		if err != nil {
			log.Fatal(err)
		}
		instruction := Instruction{Type: instructionArray[0], Advance: advance}

		instructions = append(instructions, instruction)
	}

	return instructions
}

func (submarine *Submarine) PrintSubmarineStatus(statusTitle string) {
	fmt.Println(string(utils.Yellow) + "[" + statusTitle + "]:" + string(utils.Reset))
	fmt.Println("Position: " + string(utils.Green) + strconv.Itoa(submarine.Horizontal) + string(utils.Reset))
	fmt.Println("Depth: " + string(utils.Green) + strconv.Itoa(submarine.Depth) + string(utils.Reset))
	fmt.Println("Multiplied: " + string(utils.Green) + strconv.Itoa(submarine.Horizontal*submarine.Depth) + string(utils.Reset))
	fmt.Printf("\n")
}

func (submarine *Submarine) Navigate() {
	file := utils.OpenFile(submarine.Input)
	defer file.Close()

	instructions := parseInstructions(file)

	for index := range instructions {

		switch instructions[index].Type {
		case "forward":
			submarine.Forward(instructions[index].Advance)
			break
		case "up":
			submarine.Up(instructions[index].Advance)
			break
		case "down":
			submarine.Down(instructions[index].Advance)
		}
	}

	submarine.PrintSubmarineStatus("NAVIGATED: SUBMARINE POSITION AND DEPTH")
}

func (submarine *Submarine) BackToStartingPosition() {
	submarine.Horizontal = 0
	submarine.Depth = 0
	submarine.Aim = 0
}

func (submarine *Submarine) Steer() {
	file := utils.OpenFile(submarine.Input)
	defer file.Close()

	instructions := parseInstructions(file)

	for index := range instructions {
		switch instructions[index].Type {
		case "forward":
			submarine.MoveForward(instructions[index].Advance)
			break
		case "up":
			submarine.FaceUp(instructions[index].Advance)
			break
		case "down":
			submarine.FaceDown(instructions[index].Advance)
			break
		}
	}

	submarine.PrintSubmarineStatus("STEERED: SUBMARINE POSITION AND DEPTH")
}
