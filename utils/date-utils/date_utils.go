package date_utils

import "time"

var (
	apiDateFormat   = "2008-08-01T08:18:38Z"
	apiDbDateFormat = "2008-08-01 18:18:38"
)

func GetNow() time.Time {
	return time.Now().UTC()
}
func GetNowString() string {
	return GetNow().Format(apiDateFormat)
}

func GetNowDBFormat() string {
	return GetNow().Format(apiDbDateFormat)
}
