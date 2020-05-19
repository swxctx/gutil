package xtime

import (
	"fmt"
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

// 获取本周一零点的时间戳
func GetMondayZeroTs() int64 {
	nowTime := time.Now()
	weekday := nowTime.Weekday()
	if nowTime.Weekday() == 0 {
		weekday = 7
	}
	return GetTodayZeroTs() - (int64(weekday)-1)*86400
}

// 获取指定时间的时间戳
func GetTsByTime(time time.Time) int32 {
	return int32(time.UTC().Unix())
}

// GetTsInt64ByTime 获取指定时间的时间戳
func GetTsInt64ByTime(time time.Time) int64 {
	return time.UTC().Unix()
}

// 获取指定时间戳的时间
func GetTimeByTs(ts int64) time.Time {
	time := time.Unix(ts, 0)
	return time
}

// 获取时间间隔函数 根据秒数返回 X天X小时X分钟X秒 的结果
func GetTimeRangeInfoDay(seconds int64) string {
	var (
		day, hour, minute, second int64
	)
	if seconds >= 86400 {
		day = day / 86400
		seconds = seconds % 86400
	}
	if seconds >= 3600 {
		hour = seconds / 3600
		seconds = seconds % 3600
	}
	if seconds >= 60 {
		minute = seconds / 60
		seconds = seconds % 60
	}
	second = seconds
	return fmt.Sprintf("%d小时%d分钟%d秒", hour, minute, second)
}

// GetRemainingTimeBySeconds X天X小时X分钟X秒
func GetRemainingTimeBySeconds(seconds int64) string {
	var (
		day, hour, minute, second int64
	)
	if seconds >= 86400 {
		day = seconds / 86400
		seconds = seconds % 86400
	}
	if seconds >= 3600 {
		hour = seconds / 3600
		seconds = seconds % 3600
	}
	if seconds >= 60 {
		minute = seconds / 60
		seconds = seconds % 60
	}
	second = seconds
	return fmt.Sprintf("%d天%d小时%d分钟%d秒", day, hour, minute, second)
}

// GetWeekByIntDay 获取每周返回为int值,1-7
func GetWeekByIntDay(weekDay string) int32 {
	switch weekDay {
	case "Monday":
		return 1
	case "Tuesday":
		return 2
	case "Wednesday":
		return 3
	case "Thursday":
		return 4
	case "Friday":
		return 5
	case "Saturday":
		return 6
	case "Sunday":
		return 7
	}
	return 0
}

// GetMonthByInt 获取每月返回为int值,1-12
func GetMonthByInt(weekDay string) int32 {
	switch weekDay {
	case "January":
		return 1
	case "February":
		return 2
	case "March":
		return 3
	case "April":
		return 4
	case "May":
		return 5
	case "June":
		return 6
	case "July":
		return 7
	case "August":
		return 8
	case "September":
		return 9
	case "October":
		return 10
	case "November":
		return 11
	case "December":
		return 12
	}
	return 0
}
