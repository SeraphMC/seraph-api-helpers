package colours

import (
	"fmt"
	"strconv"
	"strings"
)

// HexToInt converts a hexadecimal colour code string to its uint32 representation.
// It removes a leading "#" if present and handles 6-digit or 8-digit hex strings.
// Returns the converted value and an error if the parsing fails.
func HexToInt(hex string) (uint32, error) {
	hex = strings.TrimPrefix(hex, "#")

	if len(hex) == 8 {
		hex = hex[:6]
	}

	value, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		return 0, err
	}

	return uint32(value), nil
}

// IntToHex converts an int64 value to a hexadecimal string formatted as #RRGGBB.
func IntToHex(value int64) string {
	return fmt.Sprintf("#%06X", value)
}
