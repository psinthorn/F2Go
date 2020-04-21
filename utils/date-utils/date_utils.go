package date_utils

import "time"

var (
	apiDateFormat = "2020-04-01T08:18:38Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateFormat)
}
