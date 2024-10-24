package main

import (
	"fmt"

	"github.com/vpertj/glog"
	"github.com/vpertj/gotool"
)

var log glog.Logger

func main() {
	//ils := gotools.ToString(23)
	log = glog.NewConsoleLog("debug")
	log.Debug("dddddddddddd")
	num := int64(34)
	sdd := gotool.ToString(num)
	fmt.Println(sdd)

	//r := gotools.Rounding(99449.543556, 3)
	tmp := "2006-01-02 15:04:05"
	//tm := gotool.UnixTimeToStr(1661843160, tmp)
	stm, _ := gotool.StrToUnix("2022-08-30 15:09:16", tmp)
	nowtime := gotool.GetNowTime()
	smp, emp := gotool.DayStimeAndEtime(nowtime)

	nowunixtime := gotool.GetNowTimeUinx()
	//startTime, endTime := gotools.GetTime("2022-06")
	fmt.Println(stm, sdd, smp, emp, nowtime, nowunixtime)
}
