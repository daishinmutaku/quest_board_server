package service

import "strings"

func isIncorrectID(id int) bool {
	if id <= 0 {
		return true
	}
	return false
}

func hasBlank(str string) bool {
	if strings.Index(str, " ") > -1 {
		return true
	}
	return false
}
