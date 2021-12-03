package main

import "aoc/2021"

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
}
