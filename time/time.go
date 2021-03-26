package time

import (
	"time"
)


func Timestamp() int64 {
	return time.Now().Unix()
}

func TimestampToTime(ts int64) (datetime time.Time, err error) {
	datetime = time.Unix(ts, 0)
	return
}