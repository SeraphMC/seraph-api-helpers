package validation

import "github.com/google/uuid"

func IsValidUuid(u string) bool {
	if u == "00000000-0000-0000-0000-000000000000" {
		return false
	}
	_, err := uuid.Parse(u)
	return err == nil
}
