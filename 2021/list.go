package aoc2021

import (
	"aoc/utils"

	"github.com/charmbracelet/bubbles/list"
)

func GetList() []list.Item {
	return []list.Item{
		utils.Item{TitleValue: "DEPTH MESUREMENT INCREASE", Desc: ""},
		utils.Item{TitleValue: "THREE SONAR MESUREMENTS WINDOW", Desc: ""},
		utils.Item{TitleValue: "NAVIGATED: SUBMARINE POSITION AND DEPTH"},
		utils.Item{TitleValue: "STEERED: SUBMARINE POSITION AND DEPTH"},
		utils.Item{TitleValue: "SUBMARINE DIAGNOSTICS"},
		utils.Item{TitleValue: "STEERED: SUBMARINE POSITION AND DEPTH"},
	}
}

func Execute(cursor int) int {
	switch cursor {
	case 0:
		sonar := SonarSweep{
			Input: "./2021/input/day1.txt",
		}
		return sonar.CountDepthMesurementIncrease()
	default:
		return 0
	}
}
