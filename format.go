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

// FieldsCamelString converts the accepted string to a camel string (xx_yy to XxYy)
func FieldsCamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	length := len(s) - 1
	for i := 0; i <= length; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && length > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}
