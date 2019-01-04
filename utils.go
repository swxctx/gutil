package gutil

import (
	"crypto/rand"
	"math"
	"math/big"
	"regexp"
	"strconv"
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

// RemoveElementToint32 移除数组内的指定元素
func RemoveElementToint32(list []int32, value int32) []int32 {
	var (
		result = make([]int32, 0)
	)
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

// IsStringExists 判断字符是否在切片里面
func IsStringExists(needle string, haystack []string) bool {
	for _, b := range haystack {
		if b == needle {
			return true
		}
	}
	return false
}

// CheckLongitudeAndLatitude 校验经纬度
func CheckLongitudeAndLatitude(longitude, latitude float64) bool {
	// 校验经度范围+-(0-180),小数位不大于20位
	longitudeStr := strconv.FormatFloat(longitude, 'f', -1, 64)
	reg := regexp.MustCompile(`^(\-|\+)?(((\d|[1-9]\d|1[0-7]\d|0{1,3})\.\d{0,20})|(\d|[1-9]\d|1[0-7]\d|0{1,3})|180\.0{0,20}|180)$`)
	if !reg.MatchString(longitudeStr) {
		return false
	}
	// 校验纬度范围+-(0-90),小数位不大于20位
	latitudeStr := strconv.FormatFloat(latitude, 'f', -1, 64)
	reg = regexp.MustCompile(`^(\-|\+)?([0-8]?\d{1}\.\d{0,20}|90\.0{0,20}|[0-8]?\d{1}|90)$`)
	if !reg.MatchString(latitudeStr) {
		return false
	}
	return true
}

// EarthDistance 经纬度距离计算(返回m)
func EarthDistance(lat1, lng1, lat2, lng2 float64) float64 {
	radius := float64(6371000)
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * radius
}
