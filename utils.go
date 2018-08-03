package gutil

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"

	"github.com/shopspring/decimal"
)

/*
	input: [1,2,3,4]
	output: ?,?,?,? [1 2 3 4]
*/
// GetSQLPlaceholder sql参数拼接
func GetSQLPlaceholder(list []string, params []interface{}) (s string, paramList []interface{}) {
	n := len(list)
	for i := 0; i < len(list); i++ {
		s += "?"
		if i < n-1 {
			s += ","
		}
		paramList = append(paramList, list[i])
	}
	return
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

/*
	intput: 1,2
	output: 2
*/
// IntMax get max
func IntMax(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}

// GenRandCountForDiff 生成指定范围内的指定个数(不同的数字)
func GenRandCountForDiff(min, max int64, count int) []int64 {
	var (
		allCount map[int64]int64
		result   []int64
	)
	allCount = make(map[int64]int64)
	maxBigInt := big.NewInt(max)
	for {
		// rand
		i, _ := rand.Int(rand.Reader, maxBigInt)
		number := i.Int64()
		// 是否大于下标
		if i.Int64() >= min {
			// 是否已经存在
			_, ok := allCount[number]
			if !ok {
				result = append(result, number)
				// 添加到map
				allCount[number] = number
			}
		}
		if len(result) >= count {
			return result
		}
	}
}

// GenRandCountByArea 随机生成指定范围内的数
func GenPiecesCount(min, max int64) int32 {
	maxBigInt := big.NewInt(max)
	for {
		i, _ := rand.Int(rand.Reader, maxBigInt)
		if i.Int64() >= min {
			return int32(i.Int64())
		}
	}
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

// RemoveElementToint32 移除数组内的指定元素
func RemoveElementToint32(list []int32, value int32) []int32 {
	var result = make([]int32, 0)
	index := 0
	endIndex := len(list) - 1
	for i, s := range list {
		if s == value {
			result = append(result, list[index:i]...)
			index = i + 1
		} else if i == endIndex {
			result = append(result, list[index:endIndex+1]...)
		}
	}
	return result
}

// Int64ToString int64->string
func Int64ToString(n int64) string {
	return strconv.FormatInt(n, 10)
}

// FloatStringToInt32 小数转整数,返回格式：小数*100
func FloatStringToInt32(s string) int32 {
	d, err := decimal.NewFromString(s)
	if err != nil {
		return 0
	}
	return int32(d.Mul(decimal.New(100, 0)).IntPart())
}

/*
	input: 12.3455
	output: 12
*/
// Float32ToInt64 float32转为int64
func Float32ToInt64(count float32) (int64, error) {
	countStr := fmt.Sprintf("%0.0f", count)
	return strconv.ParseInt(depositStr, 10, 64)
}
