package gutil

import (
	"fmt"
	"strconv"
)

// StringToInt32
func StringToInt32(s string) int32 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return int32(i)
}

// Int32ToString
func Int32ToString(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

// Int32ToStringSlice
func Int32ToStringSlice(a []int32) []string {
	r := make([]string, len(a))
	for k, v := range a {
		r[k] = Int32ToString(v)
	}
	return r
}

// StringToInt32Slice
func StringToInt32Slice(a []string) []int32 {
	r := make([]int32, len(a))
	for k, v := range a {
		r[k] = StringToInt32(v)
	}
	return r
}

/*
 input: [1,2,3,4]
 output: 1,2,3,4
*/
// GetInt64SliceToString int64数组转为string
func GetInt64SliceToString(list []int64) (s string) {
	n := len(list)
	for i := 0; i < len(list); i++ {
		s += fmt.Sprintf("%v", list[i])
		if i < n-1 {
			s += ","
		}
	}
	return
}

// StringSiceToInt64 string数组转[]int64
func StringSiceToInt64(str []string) []int64 {
	var (
		ids []int64
	)
	for _, i := range str {
		id, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}
	return ids
}

// Int64ToString int64->string
func Int64ToString(n int64) string {
	return strconv.FormatInt(n, 10)
}

/*
	input: 12.3455
	output: 12
*/
// Float32ToInt64 float32转为int64
func Float32ToInt64(count float32) (int64, error) {
	countStr := fmt.Sprintf("%0.0f", count)
	return strconv.ParseInt(countStr, 10, 64)
}
