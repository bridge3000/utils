package utils

import (
	"strconv"
	"time"
)

const LOC = "Asia/Chongqing"
const YMD_LAYOUT = "2006-01-02"
const YMDHIS_LAYOUT = "2006-01-02 15:04:05"

type TimeUtil struct {
}

//获取毫秒时间戳
func (this *TimeUtil) GetMillsTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}

//日期字符串转linux时间戳
// @param timeLayout 日期字符串格式
// @param 待转的日期字符串
// @return linux时间戳
func (this *TimeUtil) DateFormat2Time(timeLayout string, toBeCharge string) time.Time {
	loc, _ := time.LoadLocation(LOC)                                //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	return theTime
}

//时间戳转日期字符串
func (this *TimeUtil) GetDateFromInt64UnixTime(timestamp int64, format string) string { //判断日期是否空
	if timestamp == 0 {
		return ""
	} else {
		tm := time.Unix(timestamp, 0)
		return tm.Format(format)
	}
}

//增加天数
func (this *TimeUtil) AddDay(nowDate time.Time, days int) time.Time {
	return time.Unix(nowDate.Unix()+24*3600*int64(days), 0)
}

func (this *TimeUtil) AddYear(nowDate time.Time, years int) time.Time {
	thisYear, _ := strconv.Atoi(nowDate.Format("2006"))
	newYear := thisYear + years
	newDateStr := strconv.Itoa(newYear) + nowDate.Format("-01-02")

	return this.DateFormat2Time("2006-01-02", newDateStr)
}
