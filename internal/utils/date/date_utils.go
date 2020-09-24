package date

import (
	"time"
)

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

// GetNow returns the actual date in UTC
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString provides the actual date in UTC, following the next format:
// yyyy-MM-ddThh:mm:ssZ
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
