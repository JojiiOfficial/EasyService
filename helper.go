package main

import "strings"

func isInStrArr(str string, arr []string) bool {
	str = strings.Trim(strings.ToLower(str), " ")
	for _, v := range arr {
		if strings.Trim(strings.ToLower(v), " ") == str {
			return true
		}
	}
	return false
}
