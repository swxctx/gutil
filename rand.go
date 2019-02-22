package gutil

import (
	"crypto/rand"
	"math/big"
	mrand "math/rand"
)

var (
	chars = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

// RandString
func RandString(l int) string {
	bs := []byte{}
	for i := 0; i < l; i++ {
		bs = append(bs, chars[mrand.Intn(len(chars))])
	}
	return string(bs)
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
