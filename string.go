package gutil

import (
	"fmt"
	"math/rand"
)

/**
    @date: 2022/3/29
**/

// SubString 截取字符串指定长度
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

// Returns an int >= min, < max
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// len 传入需要随机生产的字符串长度   min ,max 分别传入字符对应的ASCII值
func RandomString(len ,min,max int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(RandomInt(min, max))
	}
	return string(bytes)
}

