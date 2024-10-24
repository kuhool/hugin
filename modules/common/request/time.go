package request

import (
	"strings"
	"time"
)

const DATE_LAYOUT string = "2006-01-02"

type Date time.Time

// 匹配相应的日志
func (d *Date) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(DATE_LAYOUT, strings.Trim(string(data), "\""), time.Local)
	*d = Date(now)
	return
}

// 此类型只用于request，无须实现此方法
func (d Date) MarshalJSON() (data []byte, err error) {
	return
}

func (d Date) String() string {
	return d.Time().Format(DATE_LAYOUT)
}

func (d Date) Time() time.Time {
	return time.Time(d)
}
