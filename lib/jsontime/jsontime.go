package jsontime

import (
	"fmt"
	"time"
)

const (
	TIMEFORMAT string = "2006-01-02 15:04:05"
)

type JsonTime time.Time

func (t *JsonTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+TIMEFORMAT+`"`, string(data), time.Local)
	*t = JsonTime(now)
	return
}

func (t JsonTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TIMEFORMAT)+2)
	//当时间为0值时返回空字符串
	if time.Time(t).IsZero() {
		b = append(b, '"', '"')
		return b, nil
	}

	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TIMEFORMAT)
	b = append(b, '"')
	return b, nil
}

func (t JsonTime) String() string {
	fmt.Println("jsontime string")
	return time.Time(t).Format(TIMEFORMAT)
}

func (t JsonTime) Now() JsonTime {
	return JsonTime(time.Now())
}
