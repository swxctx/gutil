package xtime

import (
	"strings"
	"time"
)

// StringToUnixTs 日期格式转换为时间戳
func StringToUnixTs(t string) int64 {
	if len(t) == 0 {
		return 0
	}
	timeLayout := "yyyy-MM-dd"
	if strings.Contains(t, ":") {
		timeLayout += " HH:mm:ss"
	}
	theTime := ToTime(timeLayout, t)
	return theTime.Unix()
}

// ToTime 时间字符串转Time
func ToTime(layout, value string) *time.Time {
	dateTime, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &dateTime
}

// FormatDateTime 时间戳日期格式化
func FormatDateTime(timestamp int64, dateFormat string) string {
	return time.Unix(timestamp, 0).Format(dateFormat)
}

// GetThisMonthFirstZeros 获取本月初时间戳
func GetThisMonthFirstZeros() int64 {
	year, month, _ := time.Now().Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, time.Local).Unix()
}

// GetMonthZeroTs 获取每月时间
func GetMonthZeroTs() int64 {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, 1).Unix()
}

// GetTomorrowZeroTs 获取明天0点时间
func GetTomorrowZeroTs() int64 {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, 1).Unix()
}

// GetTodayZeroTs 获取今天0点时间
func GetTodayZeroTs() int64 {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).Unix()
}

// ReverseStringSlice 数组反转
func ReverseStringSlice(strs []string) []string {
	for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
		strs[i], strs[j] = strs[j], strs[i]
	}
	return strs
}

// GetZeroTimeStamp 获取一段时间后的零点
func GetZeroTimeStamp(years, months, days int) int64 {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).AddDate(years, months, days).Unix()
}
