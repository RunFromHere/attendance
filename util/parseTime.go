package util

import "time"

func ParseTime(t string) string {
	tt, _ := time.Parse("2006-01-02T15:04:05Z07:00", t)
	return tt.Format("2006-01-02 15:04:05")
}
