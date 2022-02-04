package datesutils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDbLayout   = "2006-01-02 15:04:05Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

func GetDbFormat() string {
	return GetNow().Format(apiDbLayout)
}
