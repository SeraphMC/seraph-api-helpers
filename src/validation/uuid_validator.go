package validation

import (
	"github.com/google/uuid"
	"strings"
)

func IsValidUuid(u string) bool {
	if u == "00000000-0000-0000-0000-000000000000" {
		return false
	}
	_, err := uuid.Parse(u)
	return err == nil
}

func FormatString(str string) string {
	return strings.ToLower(strings.ReplaceAll(str, "-", ""))
}

func RemoveDuplicateWords(input string) string {
	words := strings.Fields(input)
	wordMap := make(map[string]bool)
	var result []string

	for _, word := range words {
		if word == "|" {
			result = append(result, word)
			continue
		}
		if !wordMap[word] {
			wordMap[word] = true
			result = append(result, word)
		}
	}

	return strings.Join(result, " ")
}
