package utils

import (
	"fmt"
	"strconv"
)

func Warn(text string) {
	fmt.Println(string(Yellow) + "[" + text + "]" + string(Reset))
}

func WarnWithValue(text string, value string) {
	fmt.Println(string(Yellow) + "[" + text + "]: " + string(Reset) + value)
}

func WarnWithIntValue(text string, value int) {
	fmt.Println(string(Yellow) + "[" + text + "]: " + string(Reset) + strconv.Itoa(value))
}