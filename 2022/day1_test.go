package aoc2022

import "testing"

func TestStartingFunc(t *testing.T) {
	got := StartingFunc()
	if got != "Starting!" {
		t.Errorf("StartingFunc = %s; expected \"Starting!\"", got)
	}
}