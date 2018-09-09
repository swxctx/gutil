package gutil

import (
	"strings"
	"unsafe"
)

// FieldSnakeString input: XxYy output: xx_yy
func FieldSnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	for _, d := range StringToBytes(s) {
		if d >= 'A' && d <= 'Z' {
			if j {
				data = append(data, '_')
				j = false
			}
		} else if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(*(*string)(unsafe.Pointer(&data)))
}
