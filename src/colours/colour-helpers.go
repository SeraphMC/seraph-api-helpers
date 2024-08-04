package colours

import (
	"fmt"
	"strconv"
	"strings"
)

func HexToInt(hex string) (uint32, error) {
	hex = strings.TrimPrefix(hex, "#")

	value, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		return 0, err
	}

	if value < 0 || value > 0xFFFFFFFF {
		return 0, fmt.Errorf("value out of range for uint32: %d", value)
	}

	return uint32(value), nil
}

func IntToHex(value int64) string {
	return fmt.Sprintf("#%06X", value)
}
