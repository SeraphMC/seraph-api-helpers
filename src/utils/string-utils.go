package utils

import (
	"slices"
	"strings"
)

type RemoveStringDuplicatesOptions struct {
	Sort bool
}

// RemoveStringDuplicates removes duplicate strings from the given array and returns the result.
func RemoveStringDuplicates(strArray []string, options RemoveStringDuplicatesOptions) []string {
	for i, grant := range strArray {
		strArray[i] = strings.ToLower(strings.TrimSpace(grant))
	}

	// Sort and compact the array
	if options.Sort {
		slices.Sort(strArray)
		return slices.Compact(strArray)
	}

	// Compact requires sorting, so i'll need to manually remove duplicates
	strMap := make(map[string]bool)
	newStrArray := make([]string, 0, len(strArray))

	// Add strings to a map, if the string has been seen before, consider it a duplicate
	for _, str := range strArray {
		if !strMap[str] {
			strMap[str] = true
			newStrArray = append(newStrArray, str)
		}
	}
	return newStrArray
}
