package colours

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// HexToInt converts a hexadecimal colour code string to its uint32 representation.
// It removes a leading "#" if present and handles 6-digit or 8-digit hex strings.
// Returns the converted value and an error if the parsing fails.
func HexToInt(unformattedHex string) (uint32, error) {
	hex := strings.TrimPrefix(unformattedHex, "#")

	// Allow shorthands in hex codes (CSS 3 support)
	if len(hex) == 3 || len(hex) == 4 {
		sb := strings.Builder{}
		for _, r := range hex {
			sb.WriteRune(r)
			sb.WriteRune(r)
		}
		hex = sb.String()
	}

	// Add alpha if not present
	if len(hex) == 6 {
		hex += "FF"
	}

	// Check to ensure we have a full hex string with alpha present
	if len(hex) != 8 {
		return 0, fmt.Errorf("hex string must be 6 or 8 letters, given %s (%d chars)", hex, len(hex))
	}

	value, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		return 0, errors.New("unable to parse hex string: " + err.Error())
	}

	return uint32(value), nil
}

// IntToHex converts an uint32 value to a hexadecimal string formatted as #RRGGBBAA.
func IntToHex(value uint32) string {
	// Check if the value is opaque
	isOpaque := (value & 0xFF) == 0xFF

	if isOpaque {
		// Remove the alpha channel and print the 6-digit RGB glory
		return fmt.Sprintf("#%06X", value>>8)
	}

	// Print the full 8-digit hex code
	return fmt.Sprintf("#%08X", value)
}
