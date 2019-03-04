package gutil

import (
	"fmt"
	"image"
	"math"
	"regexp"
	"strconv"
	"unicode/utf8"

	"github.com/usthooz/gutil/http"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
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

// ToGBK
func ToGBK(s string) (string, error) {
	enc := simplifiedchinese.GBK
	trans := enc.NewEncoder()
	r, _, err := transform.String(trans, s)
	return r, err
}

// FromGBK
func FromGBK(s string) (string, error) {
	enc := simplifiedchinese.GBK
	trans := enc.NewDecoder()
	r, _, err := transform.String(trans, s)
	return r, err
}

// GetImageSize 图片(在线url)宽和高
func GetImageSize(url string) (h, w int, err error) {
	resp, _, err := xhttp.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	icfg, _, err := image.DecodeConfig(resp.Body)
	if err != nil {
		return
	}
	return icfg.Height, icfg.Width, nil
}

// FilterEmoji 过滤 emoji 表情
func FilterEmoji(content string) string {
	var (
		newContent string
	)
	for _, value := range content {
		_, size := utf8.DecodeRuneInString(string(value))
		if size <= 3 {
			f := fmt.Sprintf("%#U", value)
			if f == "U+200D" {
				continue
			}
			newContent += string(value)
		}
	}
	return newContent
}
