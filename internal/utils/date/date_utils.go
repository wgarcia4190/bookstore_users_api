package date

import (
	"time"
)

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDbLayout   = "2006-01-02 15:04:05"
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

// GetNowDB provides the actual date in UTC, following the next format:
// yyyy-MM-dd hh:mm:ss
func GetNowDB() string {
	return GetNow().Format(apiDbLayout)
}
