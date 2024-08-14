package colours

import (
	"fmt"
	"strconv"
	"strings"
)

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

func IntToHex(value int64) string {
	return fmt.Sprintf("#%06X", value)
}
