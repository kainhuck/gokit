package time

import "time"

func Timestamp() int64 {
	return time.Now().Unix()
}
