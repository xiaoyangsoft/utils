//============================================================

// 作者: 杨大雷
// 日期: 2023/12/02 11:30 AM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package timer

import (
	"fmt"
	"time"
	"utils"
)

const (
	YMD     = "2006-01-02"
	YMD_HMS = "2006-01-02 15:04:05"
)

type DayTimeParam struct {
	TimeLayout string
	Day        string         // 某一天的日期，可指定解析TimeLayout
	Timestamp  int64          // 某一天的时间戳
	Location   *time.Location // 时区
}

// 当前时间戳
func NowTimestamp(params ...DayTimeParam) int64 {
	n := time.Now()
	location := n.Location()
	if len(params) > 0 {
		param := params[0]
		location = utils.Ternary(param.Location != nil, param.Location, n.Location())
	}
	return n.In(location).Unix()
}

// 一天的零点时间戳
func DayZeroTimestamp(params ...DayTimeParam) (int64, error) {
	n := time.Now()
	location := n.Location()
	if len(params) > 0 {
		param := params[0]
		location = utils.Ternary(param.Location != nil, param.Location, n.Location())
		if param.Day != "" {
			layout := utils.Ternary(param.TimeLayout != "", param.TimeLayout, YMD)
			t, err := time.ParseInLocation(layout, param.Day, location)
			if err != nil {
				return 0, err
			}
			return t.Unix(), nil
		}
		if param.Timestamp > 0 {
			n = time.Unix(param.Timestamp, 0)
			return time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, location).Unix(), nil
		}
	}
	return time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, location).Unix(), nil
}

// 一天的23时59分59秒时间戳
func DayEndTimestamp(params ...DayTimeParam) (int64, error) {
	n := time.Now()
	location := n.Location()
	if len(params) > 0 {
		param := params[0]
		location = utils.Ternary(param.Location != nil, param.Location, n.Location())
		if param.Day != "" {
			layout := utils.Ternary(param.TimeLayout != "", param.TimeLayout, YMD_HMS)
			param.Day = fmt.Sprintf("%s 23:59:59", param.Day)
			t, err := time.ParseInLocation(layout, param.Day, location)
			if err != nil {
				return 0, err
			}
			return t.Unix(), nil
		}
		if param.Timestamp > 0 {
			n = time.Unix(param.Timestamp, 0)
			return time.Date(n.Year(), n.Month(), n.Day(), 23, 59, 59, 0, location).Unix(), nil
		}
	}
	return time.Date(n.Year(), n.Month(), n.Day(), 23, 59, 59, 0, location).Unix(), nil
}

type WeekParam struct {
	Location *time.Location // 时区
	T        time.Time      // 指定时间所在的周一零点时间戳
	PreWeeks int            // 表示：当前时间或指定时间的前 几 周一零点时间戳 (如：1 为上一个周一零点)
}

// 周一的零点
func Monday(params ...WeekParam) string {
	n := time.Now()
	location := n.Location()
	preWeeks := 0
	if len(params) > 0 {
		param := params[0]
		location = utils.Ternary(param.Location != nil, param.Location, n.Location())
		preWeeks = param.PreWeeks
		n = utils.Ternary(!param.T.IsZero(), param.T, n)
	}
	n = n.In(location)
	wd := utils.Ternary(n.Weekday() == 0, 7, n.Weekday())
	skipDays := utils.Ternary(preWeeks == 0, -int(wd)+1, utils.Ternary(preWeeks == 1, -int(wd)-6, utils.Ternary(preWeeks > 1, -int(wd)-6-(preWeeks-1)*7, 0)))
	return n.AddDate(0, 0, skipDays).Format(YMD)
}

// 周日的23时59分59秒
func Sunday(params ...WeekParam) string {
	n := time.Now()
	location := n.Location()
	preWeeks := 0
	if len(params) > 0 {
		param := params[0]
		location = utils.Ternary(param.Location != nil, param.Location, n.Location())
		preWeeks = param.PreWeeks
		n = utils.Ternary(!param.T.IsZero(), param.T, n)
	}
	n.In(location)
	wd := utils.Ternary(n.Weekday() == 0, 7, n.Weekday())
	skipDays := utils.Ternary(preWeeks == 0, 7-int(wd), utils.Ternary(preWeeks == 1, -int(wd), -int(wd)-((preWeeks-1)*7)))
	return n.AddDate(0, 0, skipDays).Format(YMD)
}
