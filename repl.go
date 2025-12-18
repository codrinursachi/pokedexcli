package main

import "strings"

func cleanInput(text string) []string {
	result := []string{}
	for word := range strings.SplitSeq(text, " ") {
		trimmed := strings.TrimSpace(word)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}
