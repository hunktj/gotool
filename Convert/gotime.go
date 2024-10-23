package gotool

import (
	"time"
)

// GetNowTime 获取当前时间
func GetNowTime() string {
	nowTime := time.Now().Format("2006-01-02 15:04:05 ")
	return nowTime
}

// GetNowTimeUinx 获取当前时间戳
func GetNowTimeUinx() int64 {
	nowTimeUnix := time.Now().Unix()
	return nowTimeUnix
}

// StrToUnix 时间->时间戳 layout:="2006-01-02 15:04:05"
func StrToUnix(timeStr, layout string) (int64, error) {
	local, err := time.LoadLocation("Asia/Shanghai") //设置时区
	if err != nil {
		return 0, err
	}
	tt, err := time.ParseInLocation(layout, timeStr, local)
	if err != nil {
		return 0, err
	}
	timeUnix := tt.Unix()
	return timeUnix, nil
}

// UnixToStr 时间戳->时间 layout:="2006-01-02 15:04:05"
func UnixToStr(timeUnix int64, layout string) string {
	timeStr := time.Unix(timeUnix, 0).Format(layout)
	return timeStr
}

// GetTime 根据日期返回当月的第一天和最后一天的时间戳
// date= 2022-6/2022-06 返回第一天和最后一天
// date= 2000-06-08/2000-6-8 返回当天的时间戳
func GetTime(date string) (int, int) {
	if date == "" {
		date = time.Now().Format("2006-1")
	}
	stime := 0
	etime := 0
	if len(date) <= 7 {
		tt, _ := time.ParseInLocation("2006-1", date, time.Local)
		stime = (int(tt.Unix()) + 8*3600) / (3600 * 24)
		etime = (int(tt.AddDate(0, 1, -1).Unix()) + 8*3600) / (3600 * 24)
	} else {
		tt, _ := time.ParseInLocation("2006-1-2", date, time.Local)
		stime = (int(tt.Unix()) + 8*3600) / (3600 * 24)
		etime = stime
	}
	return stime, etime
}

// DayStimeAndEtime 返回一天的开始时间戳和最后时间戳如：2022-10-12 00:00:00 -- 2022-10-12 23:59:59
func DayStimeAndEtime(times string) (int64, int64) {
	loc, _ := time.LoadLocation("Local")
	var T time.Time
	if len(times) == 10 {
		tt, err := time.ParseInLocation("2006-01-02 ", times, loc)
		if err != nil {
			return 0, 0
		}
		T = tt
	} else {
		ti := times[0:11]
		tt, err := time.ParseInLocation("2006-01-02 ", ti, loc)
		if err != nil {
			return 0, 0
		}
		T = tt
	}

	sTime := T.Unix()
	eTime := sTime + 86399
	return sTime, eTime
}
