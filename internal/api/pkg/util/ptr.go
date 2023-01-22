package util

import "time"

func StrPtr(s string) *string {
	return &s
}

func IntPtr(i int) *int {
	return &i
}

func Int64Ptr(i int64) *int64 {
	return &i
}

func Uint64Ptr(ui uint64) *uint64 {
	return &ui
}

func BoolPtr(b bool) *bool {
	return &b
}

func Float64Ptr(f float64) *float64 {
	return &f
}

func TimePtr(t time.Time) *time.Time {
	return &t
}

func DerefStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func DerefInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

func DerefBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

func DerefFloat64(f *float64) float64 {
	if f == nil {
		return 0
	}
	return *f
}

func DerefTime(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}
