package main

import (
	"fmt"
	"time"
)

// time包 demo
func main() {
	now := time.Now()
	fmt.Println(now)
	yeas := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Println(yeas, month, day, hour, minute, second)

	// 时间戳从 1970。1.1 到现在的秒数
	timeStamp1 := now.Unix()
	timeStamp2 := now.UnixNano()
	fmt.Println(timeStamp1, timeStamp2)

	// 时间戳转化为具体的时间格式
	t := time.Unix(timeStamp1, 0)
	fmt.Println(t)
	
	//时间间隔
	startTime := time.Now()
	time.Sleep(5 * time.Second)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Println(duration)
	// 定时器
	for tmp := range time.Tick(time.Second) {
		fmt.Println(tmp)
	}
	//时间格式化
	today := now.Format("2006-01-02 3:04:05 PM")
	fmt.Println(today)

	// 1. 拿到时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
	}

	// 2. 根据时区取解析一个字符串
	timeStr := "2009-08-07 15:30:45"
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(timeObj)
}
