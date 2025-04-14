package model

import "time"

// ParseTime parses a time string in RFC3339 format
func ParseTime(timeStr string) time.Time {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return time.Time{}
	}
	return t
}
