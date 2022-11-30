package utils

import (
	"fmt"
)

func Warn(text string) {
	fmt.Println(string(Yellow) + "[" + text + "]" + string(Reset))
}

func WarnWithValue(text string, value string) {
	fmt.Println(string(Yellow) + "[" + text + "]: " + string(Reset) + value)
}