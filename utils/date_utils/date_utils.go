package date_utils

import (
	"time"
)

const(
	ApiDateTimeFormat = time.RFC3339
	DbDateTimeFormat = "2006-01-02 15:04:05"
)

func GetCurrentDateTime() time.Time{
	return time.Now().UTC()
}

func GetApiCurrentDateTimeString() string {
	return GetCurrentDateTime().Format(ApiDateTimeFormat)
}

func GetDbCurrentDateTimeString() string {
	return GetCurrentDateTime().Format(DbDateTimeFormat)
}

func ConvertToApiDateFormat(t *time.Time) string {
	return t.Format(ApiDateTimeFormat)
}

func ConvertToDbDateFormat(t *time.Time) string {
	return t.Format(DbDateTimeFormat)
}

