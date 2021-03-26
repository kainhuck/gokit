package time

import (
	"fmt"
	"testing"
)

func TestTime(t *testing.T) {
	fmt.Println(Timestamp())
	fmt.Println(TimestampToTime(1616731096))
}
