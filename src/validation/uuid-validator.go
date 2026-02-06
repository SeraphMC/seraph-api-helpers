package validation

import (
	"strings"

	"github.com/google/uuid"
)

// IsValidUuid checks if the given string is a valid UUID.
// The function returns false for nil UUIDs or if the string does not conform to the UUID format.
func IsValidUuid(uuidAsString string) bool {
	parsedId, err := uuid.Parse(uuidAsString)
	if err != nil {
		return false
	}

	if parsedId == uuid.Nil {
		return false
	}

	return true
}

// FormatString takes an input string, replaces all occurrences of "-" with "", and converts the string to lowercase, returning the transformed result.
func FormatString(str string) string {
	return strings.ToLower(strings.ReplaceAll(str, "-", ""))
}

// RemoveDuplicateWords removes duplicate words from the input string while preserving their order.
// It treats the "|" character as a special case, ensuring it is retained in the result irrespective of duplicates.
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

// RemoveDuplicateFromArray removes duplicate strings from a slice and returns a new slice with unique elements in the order of their first appearance.
// Deprecated: RemoveDuplicateFromArray Moved from validation to utils package. Use utils.RemoveStringDuplicates instead. This function will be removed without warning.
func RemoveDuplicateFromArray(strings []string) []string {
	seen := make(map[string]bool)
	uniqueStrings := []string{}

	for _, str := range strings {
		if !seen[str] {
			seen[str] = true
			uniqueStrings = append(uniqueStrings, str)
		}
	}

	return uniqueStrings
}
