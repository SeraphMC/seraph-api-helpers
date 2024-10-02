package utils

import (
	"fmt"
	"time"
)

func FormatTimestamp(timestamp time.Time) string {
	totalDuration := time.Now().Sub(timestamp)
	totalHours := int(totalDuration.Hours())

	years := totalHours / (365 * 24)
	months := (totalHours % (365 * 24)) / (30 * 24)
	days := (totalHours % (30 * 24)) / 24
	hours := totalHours % 24

	var formattedTimeString string

	if years > 0 {
		formattedTimeString += fmt.Sprintf("%dY", years)
		if months > 0 || days > 0 || hours > 0 {
			formattedTimeString += ", "
		}
	}
	if months > 0 {
		formattedTimeString += fmt.Sprintf("%dM", months)
		if days > 0 || hours > 0 {
			formattedTimeString += ", "
		}
	}
	if days > 0 {
		formattedTimeString += fmt.Sprintf("%dD", days)
		if hours > 0 {
			formattedTimeString += ", "
		}
	}
	if hours > 0 {
		formattedTimeString += fmt.Sprintf("%dH", hours)
	}

	return formattedTimeString
}
