package utils

import "time"

const (
	Layout      = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp        = "Jan _2 15:04:05"
	StampMilli   = "Jan _2 15:04:05.000"
	StampMicro   = "Jan _2 15:04:05.000000"
	StampNano    = "Jan _2 15:04:05.000000000"
	DateTime     = "2006-01-02 15:04:05"
	DateOnly     = "2006-01-02"
	TimeOnly     = "15:04:05"
	DateOnly_Ymd = "20060102"
	DateOnly_d   = "02"
	MonthOnly_Ym = "200601"
)

// TomorrowZeroTime 函数返回明天的零点时间
func TomorrowZeroTime() time.Time {
	now := time.Now()                                                                                    // 获取当前时间
	tomorrow := now.AddDate(0, 0, 1)                                                                     // 计算明天的日期
	return time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, 0, 0, 0, tomorrow.Location()) // 构造明天的零点时间并返回
}
