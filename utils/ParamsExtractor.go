package utils

import "strconv"

func GetStringOrDefault(s string, defaultValue string) string {
	if s == "" {
		return defaultValue
	}
	return s
}

func GetIntOrDefault(s string, defaultValue int) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		return defaultValue
	}
	return r
}
