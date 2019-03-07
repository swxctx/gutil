package gutil

import (
	"time"
)

func StringPtr(s string) *string {
	return &s
}

func BoolPtr(i bool) *bool {
	return &i
}

func BytePtr(i byte) *byte {
	return &i
}

func Int32Ptr(i int32) *int32 {
	return &i
}

func Int64Ptr(i int64) *int64 {
	return &i
}

func Float64Ptr(i float64) *float64 {
	return &i
}

func TimePtr(t time.Time) *time.Time {
	return &t
}

func PtrString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func PtrBool(i *bool) bool {
	if i == nil {
		return false
	}
	return *i
}

func PtrByte(i *byte) byte {
	if i == nil {
		return byte(0)
	}
	return *i
}

func PtrInt32(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}

func PtrInt64(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

func PtrFloat64(i *float64) float64 {
	if i == nil {
		return float64(0)
	}
	return *i
}
