package utils

import "strings"

func HasEightDecimals(value string) bool {
	parts := strings.Split(value, ".")
	if len(parts) == 2 {
		return len(parts[1]) == 8
	}
	return false
}
