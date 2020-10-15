package datetime

import "time"

const (
	timeFormat = "01-02-2006 15:04:05"
)

// GetCurrentTime returns the current UTC time
func GetCurrentTime() time.Time {
	return time.Now().UTC()
}

// GetCurrentFormattedTime returns the current UTC time
// in a formatted and readable way
func GetCurrentFormattedTime() string {
	return time.Now().UTC().Format(timeFormat)
}
