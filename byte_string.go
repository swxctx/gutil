package gutil

import (
	"unsafe"
)

// StringToBytes
func StringToBytes(s string) []byte {
	sp := *(*[2]uintptr)(unsafe.Pointer(&s))
	bp := [3]uintptr{sp[0], sp[1], sp[1]}
	return *(*[]byte)(unsafe.Pointer(&bp))
}

// BytesToString []byte to string
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
