package gutil

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
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

// RandByte
func RandByte(l int) []byte {
	bs := []byte{}
	for i := 0; i < l; i++ {
		bs = append(bs, chars[mrand.Intn(len(chars))])
	}
	return bs
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

const (
	// UpperWordsAndNumber 数字与大写字母
	UpperWordsAndNumber = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// NewRandom creates a new padded Encoding defined by the given alphabet string.
func NewRandom(alphabet string) *Random {
	r := new(Random)
	diff := 64 - len(alphabet)
	if diff < 0 {
		r.substitute = []byte(alphabet[64:])
		r.substituteLen = len(r.substitute)
		alphabet = alphabet[:64]
	} else {
		r.substitute = []byte(alphabet)
		r.substituteLen = len(r.substitute)
		if diff > 0 {
			alphabet += string(bytes.Repeat([]byte{0x00}, diff))
		}
	}
	r.encoding = base64.NewEncoding(alphabet).WithPadding(base64.NoPadding)
	return r
}

// Random random string creater.
type Random struct {
	encoding      *base64.Encoding
	substitute    []byte
	substituteLen int
}

// RandomString returns a base64 encoded securely generated
// random string. It will panic if the system's secure random number generator
// fails to function correctly.
// The length n must be an integer multiple of 4, otherwise the last character will be padded with `=`.
func (r *Random) RandomString(n int) string {
	d := r.encoding.DecodedLen(n)
	buf := make([]byte, n)
	r.encoding.Encode(buf, RandomBytes(d))
	for k, v := range buf {
		if v == 0x00 {
			buf[k] = r.substitute[mrand.Intn(r.substituteLen)]
		}
	}
	return BytesToString(buf)
}

// RandomBytes returns securely generated random bytes. It will panic
// if the system's secure random number generator fails to function correctly.
func RandomBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}
