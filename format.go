package gutil

import (
	"strings"
	"unsafe"
)

// FieldSnakeString input: XxYy output: xx_yy
func FieldSnakeString(s string) string {
	res := make([]byte, 0, len(s)*2)
	flag := false
	for _, d := range StringToBytes(s) {
		if d >= 'A' && d <= 'Z' {
			if flag {
				res = append(res, '_')
				flag = false
			}
		} else if d != '_' {
			flag = true
		}
		res = append(res, d)
	}
	return strings.ToLower(*(*string)(unsafe.Pointer(&res)))
}
