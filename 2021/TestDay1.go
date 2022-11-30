package aoc2021

import "testing"

func TestCountDepthMesurementIncrease(t *testing.T) {
	sonar := SonarSweep{
		Input: "./2021/input/day1.txt",
	}
	got := sonar.CountDepthMesurementIncrease()

	if got != 1466 {
		t.Errorf("sonar.CountDepthMesurementIncrease() = %d; want 1466", got)
	}
}