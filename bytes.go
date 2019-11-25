package gutil

import (
	"encoding/json"
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

// InterfaceToBytes
func InterfaceToBytes(data interface{}) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// BytesToMapInterface
func BytesToMapInterface(data []byte) (map[string]interface{}, error) {
	m := map[string]interface{}{}
	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
