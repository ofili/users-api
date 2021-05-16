package date

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDbLayout = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	// Date and time displays in standard time zone
	return GetNow().Format(apiDateLayout)
}

func GetNowDBFormat() string {
	return GetNow().Format(apiDbLayout)
}
