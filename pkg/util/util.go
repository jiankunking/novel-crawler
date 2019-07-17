package util

import (
	"strconv"
	"strings"
)

func IsEmpty(str string) bool {
	if str == "" || len(strings.TrimSpace(str)) == 0 {
		return true
	}
	return false
}

func TakeChineseNumberFromString(chTextString string) int {
	obj := takeChineseNumberFromString(chTextString, nil, true)
	dict, ok := obj.(map[string]interface{})
	if !ok {
		return -1
	}
	val, ok := dict["replacedText"]
	if !ok {
		return -1
	}
	str := val.(string)
	if !ok {
		return -1
	}
	result, err := strconv.Atoi(str)
	if err != nil {
		return -1
	}
	return result
}
