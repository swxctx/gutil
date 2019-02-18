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
