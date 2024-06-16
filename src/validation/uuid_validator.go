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
