package gutil

import "fmt"

/**
    @date: 2022/3/29
**/

func SubString(str string, start, end int) (string, error) {
	rs := []rune(str)
	rl := len(rs)

	if rl == 0 {
		return "", nil
	}
	if start < 0 || start > rl || end < 0 || end > rl || start > end {
		return "", fmt.Errorf("参数不合法")
	}
	return string(rs[start:end]), nil
}
