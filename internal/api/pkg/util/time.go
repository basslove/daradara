package util

import "time"

func TimeToStrRFC3339(t time.Time) string {
	return t.Format(time.RFC3339)
}

func TimeToStrRFC3339Nano(t time.Time) string {
	return t.Format(time.RFC3339Nano)
}

func TimeToStrYMD(t time.Time) string {
	return t.Format("2006-01-02")
}

func TimeToStrYMDHIS(t time.Time) string {
	return t.Format("2006-01-02T15:04:05")
}
